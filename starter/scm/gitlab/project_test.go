package gitlab

import (
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
	"github.com/magiconair/properties/assert"
	"github.com/xanzy/go-gitlab"
	"os"
	"testing"
)

func TestGetProject(t *testing.T) {
	baseUrl := os.Getenv("SCM_URL")
	s := fake.NewClient("")
	s.On("SetBaseURL", nil).Return(nil)
	gra := &gitlab.Project{
		ID:   100,
		Name: "chulei",
		Namespace: &gitlab.ProjectNamespace{
			ID: 30,
		},
	}
	pid := "4"
	resp := new(gitlab.Response)
	s.On("GetProject", pid, nil).Return(gra, resp, nil)
	project := NewProject(s)
	_, _, err := project.GetProject(baseUrl, pid, os.Getenv("Token"))
	assert.Equal(t, nil, err)
}

func TestGetGroupId(t *testing.T) {
	baseUrl := os.Getenv("SCM_URL")
	s := fake.NewClient("")
	s.On("SetBaseURL", nil).Return(nil)
	gra := &gitlab.Project{
		ID:   100,
		Name: "chulei",
		Namespace: &gitlab.ProjectNamespace{
			ID: 30,
		},
	}
	pid := 4
	resp := new(gitlab.Response)
	s.On("GetProject", pid, nil).Return(gra, resp, nil)
	project := NewProject(s)
	_, err := project.GetGroupId(baseUrl, os.Getenv("Token"), pid)
	assert.Equal(t, nil, err)
}

func TestListProjects(t *testing.T) {
	baseUrl := os.Getenv("SCM_URL")
	s := fake.NewClient("")
	s.On("SetBaseURL", nil).Return(nil)
	gra := &gitlab.Project{
		ID:   100,
		Name: "chulei",
		Namespace: &gitlab.ProjectNamespace{
			ID: 30,
		},
	}
	var pro []*gitlab.Project
	pro = append(pro, gra)
	resp := new(gitlab.Response)
	s.On("ListProjects", nil, nil).Return(pro, resp, nil)
	project := NewProject(s)
	_, err := project.ListProjects(baseUrl, os.Getenv("Token"), "1", 1)
	_, err = project.ListProjects(baseUrl, os.Getenv("Token"), "", 1)
	assert.Equal(t, nil, err)

}

func TestSearch(t *testing.T) {
	baseUrl := os.Getenv("SCM_URL")
	s := fake.NewClient("")
	s.On("SetBaseURL", nil).Return(nil)
	gra := &gitlab.Project{
		ID:   100,
		Name: "chulei",
		Namespace: &gitlab.ProjectNamespace{
			ID: 30,
		},
	}
	var pro []*gitlab.Project
	pro = append(pro, gra)
	resp := new(gitlab.Response)
	s.On("ListProjects", nil, nil).Return(pro, resp, nil)
	project := NewProject(s)
	_, err := project.Search(baseUrl, os.Getenv("Token"), "")
	assert.Equal(t, nil, err)
}
