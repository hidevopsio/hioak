package gitlab

import (
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
	"github.com/magiconair/properties/assert"
	"github.com/xanzy/go-gitlab"
	"os"
	"testing"
)

func TestListProjectMembers(t *testing.T) {
	baseUrl := os.Getenv("SCM_URL")
	s := fake.NewClient("")
	s.On("SetBaseURL", nil).Return(nil)
	gra := &gitlab.GroupMember{
		ID:          100,
		Name:        "chulei",
		AccessLevel: gitlab.DeveloperPermissions,
	}
	var gro []*gitlab.GroupMember
	gro = append(gro, gra)
	resp := new(gitlab.Response)
	gid := 4
	s.On("ListGroupMembers", gid, nil, nil).Return(gro, resp, nil)
	projectMembers := &gitlab.ProjectMember{
		ID:          100,
		Name:        "chulei",
		AccessLevel: gitlab.DeveloperPermissions,
	}
	s.On("GetProjectMember", nil, nil).Return(projectMembers, resp, nil)
	projectMember := NewProjectMember(s)
	_, err := projectMember.GetProjectMember(os.Getenv("Token"), baseUrl, 1, 1, gid)
	assert.Equal(t, nil, err)
	_, err = projectMember.GetProjectMember(os.Getenv("Token"), baseUrl, 1, 100, gid)
	assert.Equal(t, nil, err)
}
