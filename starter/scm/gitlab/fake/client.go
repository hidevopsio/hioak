package fake

import (
	"github.com/stretchr/testify/mock"
	"github.com/xanzy/go-gitlab"
)

type Client struct {
	mock.Mock
}

func NewClient(token string) *Client {
	return &Client{}
}

func (c *Client) SetBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	args := c.Called(nil)
	return args.Error(0)
}

func (c *Client) GetSession(opt *gitlab.GetSessionOptions, options ...gitlab.OptionFunc) (*gitlab.Session, *gitlab.Response, error) {
	args := c.Called(nil, nil)
	return args[0].(*gitlab.Session), args[1].(*gitlab.Response), args.Error(2)
}

func (c *Client) ListGroups(opt *gitlab.ListGroupsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Group, *gitlab.Response, error) {
	args := c.Called(nil, nil)
	return args[0].([]*gitlab.Group), args[1].(*gitlab.Response), args.Error(2)
}

func (c *Client) GetGroup(gid interface{}, options ...gitlab.OptionFunc) (*gitlab.Group, *gitlab.Response, error) {
	args := c.Called(gid, nil)
	return args[0].(*gitlab.Group), args[1].(*gitlab.Response), args.Error(2)
}

func (c *Client) ListGroupProjects(gid interface{}, opt *gitlab.ListGroupProjectsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Project, *gitlab.Response, error) {
	args := c.Called(gid, nil)
	return args[0].([]*gitlab.Project), args[1].(*gitlab.Response), args.Error(2)
}

func (c *Client) ListGroupMembers(gid interface{}, opt *gitlab.ListGroupMembersOptions, options ...gitlab.OptionFunc) ([]*gitlab.GroupMember, *gitlab.Response, error) {
	args := c.Called(gid, nil, nil)
	return args[0].([]*gitlab.GroupMember), args[1].(*gitlab.Response), args.Error(2)
}

func (c *Client) GetProject(pid interface{}, options ...gitlab.OptionFunc) (*gitlab.Project, *gitlab.Response, error) {
	args := c.Called(pid, nil)
	return args[0].(*gitlab.Project), args[1].(*gitlab.Response), args.Error(2)
}

func (c *Client) ListProjects(opt *gitlab.ListProjectsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Project, *gitlab.Response, error) {
	args := c.Called(nil, nil)
	return args[0].([]*gitlab.Project), args[1].(*gitlab.Response), args.Error(2)
}

func (c *Client) ListTree(pid interface{}, opt *gitlab.ListTreeOptions, options ...gitlab.OptionFunc) ([]*gitlab.TreeNode, *gitlab.Response, error) {
	args := c.Called(nil, nil, nil)
	return args[0].([]*gitlab.TreeNode), args[1].(*gitlab.Response), args.Error(2)
}

func (c *Client) GetProjectMember(pid interface{}, user int, options ...gitlab.OptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error) {
	args := c.Called(nil, nil)
	return args[0].(*gitlab.ProjectMember), args[1].(*gitlab.Response), args.Error(2)
}
func (c *Client) GetFile(pid interface{}, opt *gitlab.GetFileOptions, options ...gitlab.OptionFunc) (*gitlab.File, *gitlab.Response, error) {
	args := c.Called(nil, nil, nil)
	return args[0].(*gitlab.File), args[1].(*gitlab.Response), args.Error(2)
}

func (c *Client) CurrentUser(options ...gitlab.OptionFunc) (*gitlab.User, *gitlab.Response, error) {
	args := c.Called(nil)
	return args[0].(*gitlab.User), args[1].(*gitlab.Response), args.Error(2)
}
