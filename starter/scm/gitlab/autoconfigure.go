package gitlab

import (
	"github.com/hidevopsio/hiboot/pkg/app"
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

func (c *configuration) GitlabGroup(token string) *Group {
	clientSet := NewClient(token)
	return  NewGroup(clientSet.Groups)
}

func (c *configuration) GitlabGroupMember(token string) *GroupMember {
	clientSet := NewClient(token)
	return  NewGroupMember(clientSet.Groups)
}


func (c *configuration) GitlabProject(token string) *Project {
	clientSet := NewClient(token)
	return  NewProject(clientSet.Projects)
}


func (c *configuration) GitlabProjectMember(token string) *ProjectMember {
	clientSet := NewClient(token)
	return  NewProjectMember(clientSet.Projects)
}


func (c *configuration) GitlabRepository(token string) *Repository {
	clientSet := NewClient(token)
	return  NewRepository(clientSet.Repositories)
}



func (c *configuration) GitlabUser(token string) *User {
	clientSet := NewClient(token)
	return  NewUser(clientSet.Users)
}


func (c *configuration) GitlabSession(token string) *Session {
	clientSet := NewClient(token)
	return  NewSession(clientSet.Session)
}


func (c *configuration) GitlabRepositoryFile(token string) *RepositoryFile {
	clientSet := NewClient(token)
	return  NewRepositoryFile(clientSet.RepositoryFiles)
}