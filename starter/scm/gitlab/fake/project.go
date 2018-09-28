package fake

import (
	"github.com/xanzy/go-gitlab"
	"github.com/stretchr/testify/mock"
)


type ProjectsService struct {
	mock.Mock
}

func (c *ProjectsService) GetProject(pid interface{}, options ...gitlab.OptionFunc) (*gitlab.Project, *gitlab.Response, error) {
	args := c.Called(pid, nil)
	return args[0].(*gitlab.Project), args[1].(*gitlab.Response), args.Error(2)
}

func (c *ProjectsService) ListProjects(opt *gitlab.ListProjectsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Project, *gitlab.Response, error) {
	args := c.Called(nil, nil)
	return args[0].([]*gitlab.Project), args[1].(*gitlab.Response), args.Error(2)
}

func (c *ProjectsService) GetProjectMember(pid interface{}, user int, options ...gitlab.OptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error) {
	args := c.Called(nil, nil)
	return args[0].(*gitlab.ProjectMember), args[1].(*gitlab.Response), args.Error(2)
}