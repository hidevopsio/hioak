package docker

import (
	"testing"
	"github.com/magiconair/properties/assert"
	)


func TestPullImages(t *testing.T) {
	image := &Image{
		Username: "",
		Password: "",
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
	image := &Image{
		Username: "",
		Password: "",
		FromImage:"docker.io/library/nginx",
		Tag: "latest",
	}

	err := image.PushImage()
	assert.Equal(t, nil, err)
}

func TestImage_GetImage(t *testing.T) {

	image := &Image{
		FromImage:"docker.io/library/nginx",
	}
	_, err := image.GetImage()
	assert.Equal(t, nil, err)
}