package docker

import (
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/hidevopsio/hioak/starter/docker/fake"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
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
	}
	client := ImageClient{
		Client:c,
	}
	var i io.ReadCloser
	c.On("ImagePull", nil, nil, nil).Return(i, errors.New("1"))
	err = client.PullImage(image)
	assert.Equal(t, errors.New("1"), err)
}

func TestImage_TagImage(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	image := &Image{
		FromImage: "docker.io/library/nginx",
		Tag:       "1.0",
	}
	client := ImageClient{
		Client:c,
	}
	c.On("ImageTag", nil, nil, nil).Return(nil)
	err = client.TagImage(image,"sha256:c82521676580c4850bb8f0d72e47390a50d60c8ffe44d623ce57be521bca9869")
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
	}
	client := ImageClient{
		Client:c,
	}
	var i io.ReadCloser
	c.On("ImagePush", nil, nil, nil).Return(i, errors.New("1"))
	err = client.PushImage(image)
	assert.Equal(t, errors.New("1"), err)
}

func TestImage_GetImage(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	image := &Image{
		FromImage: "docker.io/library/nginx",
	}
	client := ImageClient{
		Client:c,
	}
	var s []types.ImageSummary
	c.On("ImageList", nil, nil).Return(s, nil)
	_, err = client.GetImage(image)
	assert.Equal(t, nil, err)
}

func TestImage_BuildImage(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	file, err := os.Create("Dockerfile")
	assert.Equal(t, nil, err)
	file.Write([]byte("FROM k8s.gcr.io/pause:3.1"))
	file.Close()
	defer os.RemoveAll("./Dockerfile")
	image := &Image{
		BuildFiles: []string{"./Dockerfile"},
		Tags:       []string{"pause:latest"},
	}
	client := ImageClient{
		Client:c,
	}
	t.Run("should err is nil", func(t *testing.T) {
		c.On("ImageBuild", nil, nil, nil).Return(types.ImageBuildResponse{}, nil)
		_, err = client.BuildImage(image)
		assert.Equal(t, nil, err)
	})

	t.Run("shoud file not found", func(t *testing.T) {
		image.BuildFiles = []string{"notfile"}
		c.On("ImageBuild", nil, nil, nil).Return(types.ImageBuildResponse{}, nil)
		_, err = client.BuildImage(image)
		assert.NotEqual(t, nil, err)
	})
}
