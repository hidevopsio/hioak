package gitlab_test

import (
	"github.com/hidevopsio/hioak/starter/scm/gitlab"
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
	"github.com/magiconair/properties/assert"
	gogitlab "github.com/xanzy/go-gitlab"
	"os"
	"testing"
)

func TestGetProject(t *testing.T) {
	fs := new(fake.ProjectsService)
	cli := &fake.Client{
		ProjectsService: fs,
	}
	s := gitlab.NewProject(func(url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	gra := &gogitlab.Project{
		ID:   100,
		Name: "chulei",
		Namespace: &gogitlab.ProjectNamespace{
			ID: 30,
		},
	}
	pid := "4"
	resp := new(gogitlab.Response)
	fs.On("GetProject", pid, nil).Return(gra, resp, nil)
	_, _, err := s.GetProject("", pid, os.Getenv("Token"))
	assert.Equal(t, nil, err)
}

func TestGetGroupId(t *testing.T) {
	fs := new(fake.ProjectsService)
	cli := &fake.Client{
		ProjectsService: fs,
	}
	s := gitlab.NewProject(func(url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	gra := &gogitlab.Project{
		ID:   100,
		Name: "chulei",
		Namespace: &gogitlab.ProjectNamespace{
			ID: 30,
		},
	}
	pid := 4
	resp := new(gogitlab.Response)
	fs.On("GetProject", pid, nil).Return(gra, resp, nil)
	_, err := s.GetGroupId("", os.Getenv("Token"), pid)
	assert.Equal(t, nil, err)
}

func TestListProjects(t *testing.T) {
	fs := new(fake.ProjectsService)
	cli := &fake.Client{
		ProjectsService: fs,
	}
	s := gitlab.NewProject(func(url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	gra := &gogitlab.Project{
		ID:   100,
		Name: "chulei",
		Namespace: &gogitlab.ProjectNamespace{
			ID: 30,
		},
	}
	var pro []*gogitlab.Project
	pro = append(pro, gra)
	resp := new(gogitlab.Response)
	fs.On("ListProjects", nil, nil).Return(pro, resp, nil)
	_, err := s.ListProjects("", os.Getenv("Token"), "1", 1)
	_, err = s.ListProjects("", os.Getenv("Token"), "", 1)
	assert.Equal(t, nil, err)

}

func TestSearch(t *testing.T) {
	fs := new(fake.ProjectsService)
	cli := &fake.Client{
		ProjectsService: fs,
	}
	s := gitlab.NewProject(func(url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	gra := &gogitlab.Project{
		ID:   100,
		Name: "chulei",
		Namespace: &gogitlab.ProjectNamespace{
			ID: 30,
		},
	}
	var pro []*gogitlab.Project
	pro = append(pro, gra)
	resp := new(gogitlab.Response)
	fs.On("ListProjects", nil, nil).Return(pro, resp, nil)
	_, err := s.Search("", os.Getenv("Token"), "")
	assert.Equal(t, nil, err)
}
