package fake

import (
	"github.com/stretchr/testify/mock"
	"github.com/xanzy/go-gitlab"
)

type GroupsService struct {
	mock.Mock
}

func (c *GroupsService) ListGroups(opt *gitlab.ListGroupsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Group, *gitlab.Response, error) {
	args := c.Called(nil, nil)
	return args[0].([]*gitlab.Group), args[1].(*gitlab.Response), args.Error(2)
}

func (c *GroupsService) GetGroup(gid interface{}, options ...gitlab.OptionFunc) (*gitlab.Group, *gitlab.Response, error) {
	args := c.Called(gid, nil)
	return args[0].(*gitlab.Group), args[1].(*gitlab.Response), args.Error(2)
}

func (c *GroupsService) ListGroupProjects(gid interface{}, opt *gitlab.ListGroupProjectsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Project, *gitlab.Response, error) {
	args := c.Called(gid, nil)
	return args[0].([]*gitlab.Project), args[1].(*gitlab.Response), args.Error(2)
}

func (c *GroupsService) ListGroupMembers(gid interface{}, opt *gitlab.ListGroupMembersOptions, options ...gitlab.OptionFunc) ([]*gitlab.GroupMember, *gitlab.Response, error) {
	args := c.Called(gid, nil, nil)
	return args[0].([]*gitlab.GroupMember), args[1].(*gitlab.Response), args.Error(2)
}
