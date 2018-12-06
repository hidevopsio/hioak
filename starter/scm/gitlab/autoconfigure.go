package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"hidevops.io/hiboot/pkg/app"
	"hidevops.io/hiboot/pkg/at"
	"net/http"
	"strings"
)

type configuration struct {
	at.AutoConfiguration
}

func init() {
	app.Register(newConfiguration)
}

func newConfiguration() *configuration {
	return &configuration{}
}

type NewClient func(url, token string) ClientInterface

func (c *configuration) NewClient() NewClient {
	return func(url, token string) (client ClientInterface) {
		length := strings.Count(token, "") - 1
		cli := new(Client)
		if length <= 20 {
			cli.Client = gitlab.NewClient(&http.Client{}, token)
		}
		cli.Client = gitlab.NewOAuthClient(&http.Client{}, token)
		cli.Client.SetBaseURL(url)
		return cli
	}
}

func (c *configuration) Group(newCli NewClient) *Group {
	return NewGroup(newCli)
}

func (c *configuration) GroupMember(newCli NewClient) *GroupMember {
	return NewGroupMember(newCli)
}

func (c *configuration) Project(newCli NewClient) *Project {
	return NewProject(newCli)
}

func (c *configuration) ProjectMember(newCli NewClient) *ProjectMember {
	return NewProjectMember(newCli)
}

func (c *configuration) Repository(newCli NewClient) *Repository {
	return NewRepository(newCli)
}

func (c *configuration) User(newCli NewClient) *User {
	return NewUser(newCli)
}

func (c *configuration) Session(newCli NewClient) *Session {
	return NewSession(newCli)
}

func (c *configuration) RepositoryFile(newCli NewClient) *RepositoryFile {
	return NewRepositoryFile(newCli)
}
