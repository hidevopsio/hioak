package gitlab_test

import (
	"github.com/hidevopsio/hioak/starter/scm/gitlab"
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
	"github.com/magiconair/properties/assert"
	gogitlab "github.com/xanzy/go-gitlab"
	"os"
	"testing"
)

func TestListProjectMembers(t *testing.T) {
	fs := new(fake.ProjectsService)
	cli := &fake.Client{
		ProjectsService: fs,
	}
	s := gitlab.NewProjectMember(func(url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	gra := &gogitlab.GroupMember{
		ID:          100,
		Name:        "chulei",
		AccessLevel: gogitlab.DeveloperPermissions,
	}
	var gro []*gogitlab.GroupMember
	gro = append(gro, gra)
	resp := new(gogitlab.Response)
	gid := 4
	fs.On("ListGroupMembers", gid, nil, nil).Return(gro, resp, nil)
	projectMembers := &gogitlab.ProjectMember{
		ID:          100,
		Name:        "chulei",
		AccessLevel: gogitlab.DeveloperPermissions,
	}
	fs.On("GetProjectMember", nil, nil).Return(projectMembers, resp, nil)
	_, err := s.GetProjectMember(os.Getenv("Token"), "", 1, 1, gid)
	assert.Equal(t, nil, err)
	_, err = s.GetProjectMember(os.Getenv("Token"), "", 1, 100, gid)
	assert.Equal(t, nil, err)
}
