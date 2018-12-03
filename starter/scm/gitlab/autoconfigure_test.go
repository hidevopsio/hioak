package gitlab

import (
	"github.com/magiconair/properties/assert"
	"github.com/xanzy/go-gitlab"
	"net/http"
	"testing"
)

func TestAutoNewClient(t *testing.T)  {
	c := newConfiguration()
	c.NewClient()
	gitlab.NewClient(&http.Client{}, "")
	project := c.Project(nil)
	assert.Equal(t, &Project{}, project)
	group := c.Group(nil)
	assert.Equal(t, &Group{}, group)
	groupMember := c.GroupMember(nil)
	assert.Equal(t, &GroupMember{}, groupMember)
	session := c.Session(nil)
	assert.Equal(t, &Session{}, session)
	user := c.User(nil)
	assert.Equal(t, &User{}, user)
	repository := c.Repository(nil)
	assert.Equal(t, &Repository{}, repository)
	repositoryFile := c.RepositoryFile(nil)
	assert.Equal(t, &RepositoryFile{}, repositoryFile)
	projectMember := c.ProjectMember(nil)
	assert.Equal(t, &ProjectMember{}, projectMember)
}
