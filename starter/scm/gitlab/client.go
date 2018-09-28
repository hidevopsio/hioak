package gitlab

import (
	"github.com/xanzy/go-gitlab"
)

type ClientInterface interface {
	SetBaseURL(baseUrl string) error
	Session() SessionInterface
	Group() GroupInterface
	GroupMember() GroupMemberInterface
	Project() ProjectInterface
	ProjectMember() ProjectMemberInterface
	RepositoryFile() RepositoryFileInterface
	Repository() RepositoryInterface
	User() UserInterface
}

type SessionInterface interface {
	GetSession(opt *gitlab.GetSessionOptions, options ...gitlab.OptionFunc) (*gitlab.Session, *gitlab.Response, error)
}

type GroupInterface interface {
	ListGroups(opt *gitlab.ListGroupsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Group, *gitlab.Response, error)
	GetGroup(gid interface{}, options ...gitlab.OptionFunc) (*gitlab.Group, *gitlab.Response, error)
	ListGroupProjects(gid interface{}, opt *gitlab.ListGroupProjectsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Project, *gitlab.Response, error)
}

type GroupMemberInterface interface {
	ListGroupMembers(gid interface{}, opt *gitlab.ListGroupMembersOptions, options ...gitlab.OptionFunc) ([]*gitlab.GroupMember, *gitlab.Response, error)
}

type ProjectInterface interface {
	GetProject(pid interface{}, options ...gitlab.OptionFunc) (*gitlab.Project, *gitlab.Response, error)
	ListProjects(opt *gitlab.ListProjectsOptions, options ...gitlab.OptionFunc) ([]*gitlab.Project, *gitlab.Response, error)
}

type ProjectMemberInterface interface {
	GetProjectMember(pid interface{}, user int, options ...gitlab.OptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error)
}

type RepositoryFileInterface interface {
	GetFile(pid interface{}, opt *gitlab.GetFileOptions, options ...gitlab.OptionFunc) (*gitlab.File, *gitlab.Response, error)
}

type RepositoryInterface interface {
	ListTree(pid interface{}, opt *gitlab.ListTreeOptions, options ...gitlab.OptionFunc) ([]*gitlab.TreeNode, *gitlab.Response, error)
}

type UserInterface interface {
	CurrentUser(options ...gitlab.OptionFunc) (*gitlab.User, *gitlab.Response, error)
}

type Client struct {
	*gitlab.Client
}

func (c *Client) Session() SessionInterface {
	return c.Client.Session
}

func (c *Client) Group() GroupInterface {
	return c.Client.Groups
}

func (c *Client) GroupMember() GroupMemberInterface {
	return c.Client.Groups
}

func (c *Client) ProjectMember() ProjectMemberInterface {
	return c.Client.Projects
}

func (c *Client) Project() ProjectInterface {
	return c.Client.Projects
}

func (c *Client) Repository() RepositoryInterface {
	return c.Client.Repositories
}

func (c *Client) RepositoryFile() RepositoryFileInterface {
	return c.Client.RepositoryFiles
}

func (c *Client) User() UserInterface {
	return c.Client.Users
}