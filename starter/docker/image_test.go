package docker_test

import (
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/hidevopsio/hioak/starter/docker/fake"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
	"github.com/hidevopsio/hioak/starter/docker"
	"github.com/hidevopsio/hiboot/pkg/app/cli"
	"github.com/prometheus/common/log"
)

func TestPullImages(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	image := &docker.Image{
		Username:  "",
		Password:  "",
		FromImage: "docker.io/library/nginx",
		Tag:       "latest",
		Client:    c,
	}
	var i io.ReadCloser
	c.On("ImagePull", nil, nil, nil).Return(i, errors.New("1"))
	err = image.PullImage()
	assert.Equal(t, errors.New("1"), err)
}

func TestImage_TagImage(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	image := &docker.Image{
		FromImage: "docker.io/library/nginx",
		Tag:       "1.0",
		Client:    c,
	}
	c.On("ImageTag", nil, nil, nil).Return(nil)
	err = image.TagImage("sha256:c82521676580c4850bb8f0d72e47390a50d60c8ffe44d623ce57be521bca9869")
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
		Client:    c,
	}
	var i io.ReadCloser
	c.On("ImagePush", nil, nil, nil).Return(i, errors.New("1"))
	err = image.PushImage()
	assert.Equal(t, errors.New("1"), err)
}

func TestImage_GetImage(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	image := &docker.Image{
		FromImage: "docker.io/library/nginx",
		Client:    c,
	}
	var s []types.ImageSummary
	c.On("ImageList", nil, nil).Return(s, nil)
	_, err = image.GetImage()
	assert.Equal(t, nil, err)
}

func TestImage_BuildImage(t *testing.T) {
	c, err := fake.NewClient()
	assert.Equal(t, nil, err)
	file,err:=os.Create("Dockerfile")
	assert.Equal(t, nil, err)
	file.Write([]byte("FROM k8s.gcr.io/pause:3.1"))
	file.Close()
	defer os.RemoveAll("./Dockerfile")
	image := &docker.Image{
		DockerFile: "./Dockerfile",
		Tags:       []string{"pause:latest"},
		Client:     c,
	}

	t.Run("should err is nil", func(t *testing.T) {
		c.On("ImageBuild", nil, nil, nil).Return(types.ImageBuildResponse{}, nil)
		_,err = image.BuildImage()
		assert.Equal(t, nil, err)
	})


	t.Run("shoud file not found" , func(t *testing.T) {
		image.DockerFile = "notfile"
		c.On("ImageBuild", nil, nil, nil).Return(types.ImageBuildResponse{}, nil)
		_,err = image.BuildImage()
		assert.NotEqual(t, nil, err)
	})
}

// TestCommand is the root command
type TestCommand struct {
	// embedded cli.BaseCommand
	cli.BaseCommand

	dockerImage *docker.Image
}

func newTestCommand(dockerImage *docker.Image) *TestCommand  {
	return &TestCommand{dockerImage: dockerImage}
}

func (c *TestCommand) OnCreate(args []string)  {
	log.Debugf("OnNewImage")
}

func TestApp(t *testing.T) {
	testApp := cli.NewTestApplication(t, newTestCommand)
	t.Run("should run test command", func(t *testing.T) {
		_, err := testApp.RunTest("create")
		assert.Equal(t, nil, err)
	})
}