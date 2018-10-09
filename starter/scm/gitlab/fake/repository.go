package fake

import (
	"github.com/stretchr/testify/mock"
	"github.com/xanzy/go-gitlab"
)

type RepositoriesService struct {
	mock.Mock
}

func (c *RepositoriesService) ListTree(pid interface{}, opt *gitlab.ListTreeOptions, options ...gitlab.OptionFunc) ([]*gitlab.TreeNode, *gitlab.Response, error) {
	args := c.Called(nil, nil, nil)
	return args[0].([]*gitlab.TreeNode), args[1].(*gitlab.Response), args.Error(2)
}
