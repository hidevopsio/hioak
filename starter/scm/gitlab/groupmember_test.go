package gitlab

import (
"testing"
"os"
"github.com/magiconair/properties/assert"
	"github.com/xanzy/go-gitlab"
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
)

func TestListGroupMembers(t *testing.T) {
	baseUrl := os.Getenv("SCM_URL")
	s := fake.NewClient("")
	s.On("SetBaseURL", nil).Return(nil)
	gra := &gitlab.GroupMember{
		ID: 100,
		Name: "chulei",
		AccessLevel: gitlab.DeveloperPermissions,
	}
	var gro []*gitlab.GroupMember
	gro = append(gro, gra)
	gr := new(gitlab.Response)
	gid := 4
	s.On("ListGroupMembers", gid, nil, nil).Return(gro, gr, nil)

	groupMember := NewGroupMember(s)
	_, err := groupMember.ListGroupMembers(os.Getenv(""), baseUrl, gid, 100)
	assert.Equal(t, nil, err)
}

func TestGetGroupMember(t *testing.T) {
	baseUrl := os.Getenv("SCM_URL")
	s := fake.NewClient("")
	s.On("SetBaseURL", nil).Return(nil)
	gra := &gitlab.GroupMember{
		ID: 100,
		Name: "chulei",
	}
	var gro []*gitlab.GroupMember
	gro = append(gro, gra)
	gr := new(gitlab.Response)
	gid := 4
	s.On("ListGroupMembers", gid, nil, nil).Return(gro, gr, nil)

	groupMember := NewGroupMember(s)
	_, err := groupMember.GetGroupMember(os.Getenv(""), baseUrl, gid, 100)
	assert.Equal(t, nil, err)
}

