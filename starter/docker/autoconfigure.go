package docker

import (
	"github.com/hidevopsio/hiboot/pkg/app"
	"github.com/prometheus/common/log"
)

type configuration struct {
	app.Configuration
}

func init() {
	app.AutoConfiguration(newConfiguration)
}

func newConfiguration() *configuration {
	return &configuration{}
}

func (c *configuration) DockerImage() (image *Image) {
	clientSet, err := NewClient()
	if err != nil {
		log.Errorf("new image err :%v", err)
		return
	}
	image = NewImage(clientSet)
	return
}
