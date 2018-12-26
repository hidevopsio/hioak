package docker_test

import (
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/prometheus/common/log"
	"github.com/stretchr/testify/assert"
	"hidevops.io/hiboot/pkg/app/cli"
	"hidevops.io/hioak/starter/docker"
	"hidevops.io/hioak/starter/docker/fake"
	"io"
	"os"
	"testing"
)

func TestPullImages(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	image := &docker.Image{
		Username:  "",
		Password:  "",
		FromImage: "docker.io/library/nginx",
		Tag:       "latest",
	}
	client := docker.ImageClient{
		Client: c,
	}
	var i io.ReadCloser
	c.On("ImagePull", nil, nil, nil).Return(i, errors.New("1"))
	err = client.PullImage(image)
	assert.Equal(t, errors.New("1"), err)
}

func TestImage_TagImage(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	image := &docker.Image{
		FromImage: "docker.io/library/nginx",
		Tag:       "1.0",
	}
	client := docker.ImageClient{
		Client: c,
	}
	c.On("ImageTag", nil, nil, nil).Return(nil)
	err = client.TagImage(image, "sha256:c82521676580c4850bb8f0d72e47390a50d60c8ffe44d623ce57be521bca9869")
	assert.Equal(t, nil, err)
}

func TestImage_PushImage(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	image := &docker.Image{
		Username:  "",
		Password:  "",
		FromImage: "docker.io/library/nginx",
		Tag:       "latest",
	}
	client := docker.ImageClient{
		Client: c,
	}
	var i io.ReadCloser
	c.On("ImagePush", nil, nil, nil).Return(i, errors.New("1"))
	err = client.PushImage(image)
	assert.Equal(t, errors.New("1"), err)
}

func TestImage_GetImage(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	image := &docker.Image{
		FromImage: "docker.io/library/nginx",
	}
	client := docker.ImageClient{
		Client: c,
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
	image := &docker.Image{
		BuildFiles: []string{"./Dockerfile"},
		Tags:       []string{"pause:latest"},
	}
	client := docker.ImageClient{
		Client: c,
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

// TestCommand is the root command
type TestCommand struct {
	// embedded cli.BaseCommand
	cli.RootCommand

	imageClient *docker.ImageClient
}

func newTestCommand(imageClient *docker.ImageClient) *TestCommand {
	return &TestCommand{imageClient: imageClient}
}

func (c *TestCommand) OnCreate(args []string) bool {
	log.Debugf("OnNewImage")
	return true
}

func TestApp(t *testing.T) {
	testApp := cli.NewTestApplication(t, newTestCommand)
	t.Run("should run create command", func(t *testing.T) {
		_, err := testApp.Run("create")
		assert.Equal(t, nil, err)
	})
}


/*
func TestImageGetImage(t *testing.T) {
	c, err := docker.NewClient()
	assert.Equal(t, nil, err)
	image := &docker.Image{
		FromImage: "hiadmin1",
	}
	client := docker.ImageClient{
		Client: c,
	}
	s, err := client.GetImage(image)
	log.Info(s)
}*/