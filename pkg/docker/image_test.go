package docker

import (
	"testing"
		"github.com/magiconair/properties/assert"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hioak/pkg"
)


func TestPullImages(t *testing.T) {
	cli := orch.GetClientInstance()
	token := cli.Config().BearerToken
	image := &Image{
		Username: "unused",
		Password: token,
		FromImage:"docker-registry-default.app.vpclub.io/ecmp-dev/major-web",
		Tag: "v1",
	}

	err := image.PullImage()
	assert.Equal(t, nil, err)
}

func TestPullImages1(t *testing.T) {
	cli := orch.GetClientInstance()
	token := cli.Config().BearerToken
	image := &Image{
		Username: "unused",
		Password: token,
		FromImage:"docker.io/library/nginx",
		Tag: "latest",
	}

	err := image.PullImage()
	assert.Equal(t, nil, err)
}

func TestImage_TagImage(t *testing.T) {
	image := &Image{
		FromImage:"docker.io/library/nginx",
		Tag: "1.0",
	}

	err := image.TagImage("sha256:c82521676580c4850bb8f0d72e47390a50d60c8ffe44d623ce57be521bca9869")
	assert.Equal(t, nil, err)
}

func TestImage_PushImage(t *testing.T) {
	cli := orch.GetClientInstance()
	token := cli.Config().BearerToken
	image := &Image{
		Username: "unused",
		Password: token,
		FromImage:"docker-registry-default.app.vpclub.io/ecmp-dev/major-web",
		Tag: "v1",
	}

	err := image.PushImage()
	assert.Equal(t, nil, err)
}

func TestImage_GetImage(t *testing.T) {

	image := &Image{
		FromImage:"nginx",
	}
	s, err := image.GetImage()
	assert.Equal(t, nil, err)
	log.Info(s.ID)
	assert.Equal(t, s.ID, "sha256:c82521676580c4850bb8f0d72e47390a50d60c8ffe44d623ce57be521bca9869")
}