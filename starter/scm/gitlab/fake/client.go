package fake

import (
	"github.com/stretchr/testify/mock"
	"hidevops.io/hioak/starter/scm/gitlab"
)

type Client struct {
	mock.Mock
	GroupsService          *GroupsService
	ProjectsService        *ProjectsService
	RepositoriesService    *RepositoriesService
	RepositoryFilesService *RepositoryFilesService
	SessionService         *SessionService
	UsersService           *UsersService
}

func NewClient(url, token string) (client gitlab.ClientInterface) {
	cli := new(Client)
	return cli
}

func (c *Client) SetBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	args := c.Called(nil)
	return args.Error(0)
}

func (c *Client) Session() gitlab.SessionInterface {
	return c.SessionService
}

func (c *Client) Group() gitlab.GroupInterface {
	return c.GroupsService
}

func (c *Client) GroupMember() gitlab.GroupMemberInterface {
	return c.GroupsService
}

func (c *Client) ProjectMember() gitlab.ProjectMemberInterface {
	return c.ProjectsService
}

func (c *Client) Project() gitlab.ProjectInterface {
	return c.ProjectsService
}

func (c *Client) Repository() gitlab.RepositoryInterface {
	return c.RepositoriesService
}

func (c *Client) RepositoryFile() gitlab.RepositoryFileInterface {
	return c.RepositoryFilesService
}

func (c *Client) User() gitlab.UserInterface {
	return c.UsersService
}
