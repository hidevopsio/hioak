package docker

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"github.com/hidevopsio/hiboot/pkg/log"
)

type Image struct {
	FromImage string            `json:"from_image"`
	Tag       string            `json:"tag"`
	Username  string            `json:"username"`
	Password  string            `json:"password"`
	Size      string            `json:"size"`
	ParenID   string            `json:"paren_id"`
	Labels    map[string]string `json:"labels"`
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
	ref := ""
	if i.Tag == "" {
		ref = i.FromImage + ":latest"
	} else {
		ref = i.FromImage + ":" + i.Tag
	}
	out, err := cli.ImagePull(ctx, ref, types.ImagePullOptions{RegistryAuth: authStr})
	log.Info(err)
	if err != nil {
		return err
	}

	defer out.Close()
	io.Copy(os.Stdout, out)
	return nil
}

