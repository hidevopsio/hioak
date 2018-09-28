package gitlab_test

import (
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
	gogitlab "github.com/xanzy/go-gitlab"
	"github.com/hidevopsio/hioak/starter/scm/gitlab"
)

func TestListGroups(t *testing.T) {
	fs := new(fake.GroupsService)
	cli := &fake.Client{
		GroupsService: fs,
	}
	s := gitlab.NewGroup(func (url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	cli.On("Session", nil, nil).Return(fs)
	gra := &gogitlab.Group{
		ID:   100,
		Name: "chulei",
	}
	var gro []*gogitlab.Group
	gro = append(gro, gra)
	gr := new(gogitlab.Response)
	fs.On("ListGroups", nil, nil).Return(gro, gr, nil)
	page := 1
	_, err := s.ListGroups(os.Getenv("SCM_TOKEN"), "", page)
	assert.Equal(t, nil, err)
}

func TestGetGroup(t *testing.T) {
	fs := new(fake.GroupsService)
	cli := &fake.Client{
		GroupsService: fs,
	}
	s := gitlab.NewGroup(func (url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	gra := &gogitlab.Group{
		ID:   100,
		Name: "chulei",
	}
	gr := new(gogitlab.Response)
	fs.On("GetGroup", 1, nil).Return(gra, gr, nil)

	_, err := s.GetGroup(os.Getenv("SCM_TOKEN"), "", 1)
	assert.Equal(t, nil, err)
}

func TestListGroupProjects1(t *testing.T) {
	fs := new(fake.GroupsService)
	cli := &fake.Client{
		GroupsService: fs,
	}
	s := gitlab.NewGroup(func (url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	fs.On("SetBaseURL", nil).Return(nil)
	gra := &gogitlab.Project{
		ID:   100,
		Name: "chulei",
	}
	var projects []*gogitlab.Project
	projects = append(projects, gra)
	gr := new(gogitlab.Response)
	fs.On("ListGroupProjects", 1, nil).Return(projects, gr, nil)
	_, err := s.ListGroupProjects(os.Getenv("SCM_TOKEN"), "", 1, 1)
	assert.Equal(t, nil, err)
}
