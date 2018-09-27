package docker

import (
	"encoding/base64"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/hidevopsio/hiboot/pkg/log"
	"golang.org/x/net/context"
	"io"
	"os"
)

type Image struct {
	client        ClientInterface
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
}

type ImageInterface interface {
	PullImage() error
}

func NewImage(c ClientInterface) *Image {
	return &Image{
		client: c,
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
	out, err := i.client.ImagePull(ctx, ref, types.ImagePullOptions{RegistryAuth: authStr})
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
	err := i.client.ImageTag(ctx, imageID, ref)
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
	out, err := i.client.ImagePush(ctx, ref, types.ImagePushOptions{RegistryAuth: authStr})
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
	summaries, err := i.client.ImageList(ctx, opt)
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
