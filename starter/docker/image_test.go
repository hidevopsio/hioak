package docker

import (
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/hidevopsio/hioak/starter/docker/fake"
	"github.com/magiconair/properties/assert"
	"io"
	"testing"
)

func TestPullImages(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	image := &Image{
		Username:  "",
		Password:  "",
		FromImage: "docker.io/library/nginx",
		Tag:       "latest",
		client:    c,
	}
	var i io.ReadCloser
	c.On("ImagePull", nil, nil, nil).Return(i, errors.New("1"))
	err = image.PullImage()
	assert.Equal(t, errors.New("1"), err)
}

func TestImage_TagImage(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	image := &Image{
		FromImage: "docker.io/library/nginx",
		Tag:       "1.0",
		client:    c,
	}
	c.On("ImageTag", nil, nil, nil).Return(nil)
	err = image.TagImage("sha256:c82521676580c4850bb8f0d72e47390a50d60c8ffe44d623ce57be521bca9869")
	assert.Equal(t, nil, err)
}

func TestImage_PushImage(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	image := &Image{
		Username:  "",
		Password:  "",
		FromImage: "docker.io/library/nginx",
		Tag:       "latest",
		client:    c,
	}
	var i io.ReadCloser
	c.On("ImagePush", nil, nil, nil).Return(i, errors.New("1"))
	err = image.PushImage()
	assert.Equal(t, errors.New("1"), err)
}

func TestImage_GetImage(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	image := &Image{
		FromImage: "docker.io/library/nginx",
		client:    c,
	}
	var s []types.ImageSummary
	c.On("ImageList", nil, nil).Return(s, nil)
	_, err = image.GetImage()
	assert.Equal(t, nil, err)
}
