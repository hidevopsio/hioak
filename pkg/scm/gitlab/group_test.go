package gitlab

import (
	"testing"
	"github.com/magiconair/properties/assert"
		"os"
	"github.com/hidevopsio/hioak/pkg/scm/gitlab/fake"
	"github.com/xanzy/go-gitlab"
)


func TestListGroups(t *testing.T) {
	baseUrl := os.Getenv("SCM_URL")
	s := fake.NewClient("")
	s.On("SetBaseURL", nil).Return(nil)
	gra := &gitlab.Group{
		ID: 100,
		Name: "chulei",
	}
	var gro []*gitlab.Group
	gro = append(gro, gra)
	gr := new(gitlab.Response)
	s.On("ListGroups", nil, nil).Return(gro, gr, nil)
	page := 1
	group := NewGroup(s)
	_, err := group.ListGroups(os.Getenv("SCM_TOKEN"), baseUrl, page)
	assert.Equal(t, nil, err)
}

func TestGetGroup(t *testing.T) {
	baseUrl := os.Getenv("SCM_URL")
	s := fake.NewClient("")
	s.On("SetBaseURL", nil).Return(nil)
	gra := &gitlab.Group{
		ID: 100,
		Name: "chulei",
	}
	gr := new(gitlab.Response)
	s.On("GetGroup", 1, nil).Return(gra, gr, nil)
	group := NewGroup(s)
	_, err := group.GetGroup(os.Getenv("SCM_TOKEN"), baseUrl, 1)
	assert.Equal(t, nil, err)
}

func TestListGroupProjects1(t *testing.T)   {
	baseUrl := os.Getenv("SCM_URL")
	s := fake.NewClient("")
	s.On("SetBaseURL", nil).Return(nil)
	gra := &gitlab.Project{
		ID: 100,
		Name: "chulei",
	}
	var projects []*gitlab.Project
	projects = append(projects, gra)
	gr := new(gitlab.Response)
	s.On("ListGroupProjects", 1, nil).Return(projects, gr, nil)
	group := NewGroup(s)
	_, err := group.ListGroupProjects(os.Getenv("SCM_TOKEN"), baseUrl, 1, 1)
	assert.Equal(t, nil, err)
}

