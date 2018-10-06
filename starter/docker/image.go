package docker

import (
	"archive/tar"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/hidevopsio/hiboot/pkg/log"
	"golang.org/x/net/context"
	"io"
	"os"
	"time"
)

type Image struct {
	Client        ClientInterface
	FromImage     string            `json:"from_image"`
	Tag           string            `json:"tag"`
	Username      string            `json:"username"`
	Password      string            `json:"password"`
	Size          string            `json:"size"`
	ParenID       string            `json:"paren_id"`
	Labels        map[string]string `json:"labels"`
	Endpoint      string            `json:"endpoint"`
	IdentityToken string            `json:"identitytoken,omitempty"`
	RegistryToken string            `json:"registrytoken,omitempty"`
	ServerAddress string            `json:"server_address"`
	BuildFiles    []string          `json:"build_file"`
	Tags          []string          `json:"tags"`
}

type ImageInterface interface {
	PullImage() error
}

func NewImage(c ClientInterface) *Image {
	return &Image{
		Client: c,
	}
}

func (i *Image) PullImage() error {
	log.Info("image pull :")
	ctx := context.Background()
	authConfig := types.AuthConfig{
		Username: i.Username,
		Password: i.Password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	ref := GetTag(i.Tag, i.FromImage)
	out, err := i.Client.ImagePull(ctx, ref, types.ImagePullOptions{RegistryAuth: authStr})
	if err != nil {
		log.Info("ImagePull error:", err)
		return err
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
	return nil
}

func (i *Image) TagImage(imageID string) error {
	log.Info("imgaes tag")
	ctx := context.Background()
	ref := GetTag(i.Tag, i.FromImage)
	err := i.Client.ImageTag(ctx, imageID, ref)
	return err
}

func (i *Image) PushImage() error {
	log.Info("image push ")
	ctx := context.Background()
	authConfig := types.AuthConfig{
		Username: i.Username,
		Password: i.Password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	ref := GetTag(i.Tag, i.FromImage)
	out, err := i.Client.ImagePush(ctx, ref, types.ImagePushOptions{RegistryAuth: authStr})
	if err != nil {
		log.Info("ImagePush error:", err)
		return err
	}

	defer out.Close()
	io.Copy(os.Stdout, out)
	return nil
}

func (i *Image) GetImage() (types.ImageSummary, error) {
	log.Info("get image ")
	ctx := context.Background()
	s := types.ImageSummary{}
	ref := GetTag(i.Tag, i.FromImage)
	opt := types.ImageListOptions{}
	summaries, err := i.Client.ImageList(ctx, opt)
	for _, summary := range summaries {
		log.Infof("Summary RepoTags: %v ", summary.RepoTags)
		for _, tag := range summary.RepoTags {
			if tag == ref {
				log.Infof("Summary RepoTags: %v ", summary.RepoTags)
				return summary, nil
			}
		}
	}
	return s, err
}

func GetTag(tag, name string) string {
	ref := ""
	if tag == "" {
		ref = name + ":latest"
	} else {
		ref = name + ":" + tag
	}
	return ref
}

func (i *Image) BuildImage() (*types.ImageBuildResponse, error) {
	var files []*os.File
	for _, fileName := range i.BuildFiles {
		f, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	defer func() {
		for _, f := range files {
			f.Close()
		}
	}()
	tarName := fmt.Sprintf("%s-build.tar", time.Now())
	if err := Compress(files, tarName); err != nil {
		return nil, err
	}
	defer os.RemoveAll(tarName)

	dockerBuildContext, err := os.Open(tarName)
	if err != nil {
		return nil, err
	}
	defer dockerBuildContext.Close()

	options := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       i.Tags,
		Remove:     true}

	ImageBuildResponse, err := i.Client.ImageBuild(context.Background(), dockerBuildContext, options)
	if err != nil {
		log.Debugf("Error %v", err)
		return nil, err
	}
	return &ImageBuildResponse, nil
}

func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	gw := gzip.NewWriter(d)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	for _, file := range files {
		err := compress(file, "", tw)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, tw *tar.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(info, "")
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
