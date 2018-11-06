package docker

import (
	"github.com/prometheus/common/log"
	"hidevops.io/hiboot/pkg/app"
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

func (c *configuration) ImageClient() (imageClient *ImageClient) {
	clientSet, err := NewClient()
	if err != nil {
		log.Errorf("new image err :%v", err)
		return
	}
	imageClient = NewImage(clientSet)
	return
}
