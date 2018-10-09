package fake

import (
	"github.com/stretchr/testify/mock"
	"github.com/xanzy/go-gitlab"
)

type RepositoryFilesService struct {
	mock.Mock
}

func (c *RepositoryFilesService) GetFile(pid interface{}, opt *gitlab.GetFileOptions, options ...gitlab.OptionFunc) (*gitlab.File, *gitlab.Response, error) {
	args := c.Called(nil, nil, nil)
	return args[0].(*gitlab.File), args[1].(*gitlab.Response), args.Error(2)
}
