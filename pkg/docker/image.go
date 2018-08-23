package docker

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"os"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/docker/docker/client"
)

type Image struct {
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

func (i *Image) PullImage() error {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Info("client new envclient error:", err)
		return err
	}

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
	out, err := cli.ImagePull(ctx, ref, types.ImagePullOptions{RegistryAuth: authStr})
	if err != nil {
		log.Info("ImagePull error:", err)
		return err
	}

	defer out.Close()
	io.Copy(os.Stdout, out)
	return nil
}

func (i *Image) TagImage(imageID string) error {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	ref := GetTag(i.Tag, i.FromImage)
	cli.ImageTag(ctx, imageID, ref)
	return err
}

func (i *Image) PushImage() error {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Info("client new envclient error:", err)
		return err
	}

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
	out, err := cli.ImagePush(ctx, ref, types.ImagePushOptions{RegistryAuth: authStr})
	if err != nil {
		log.Info("ImagePull error:", err)
		return err
	}

	defer out.Close()
	io.Copy(os.Stdout, out)
	return nil
}

func (i *Image) GetImage() (types.ImageSummary, error) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	s := types.ImageSummary{}
	if err != nil {
		return s, err
	}
	ref := GetTag(i.Tag, i.FromImage)
	opt := types.ImageListOptions{}
	summaries, err := cli.ImageList(ctx, opt)
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
