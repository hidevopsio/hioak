package gitlab_test

import (
	"github.com/magiconair/properties/assert"
	gogitlab "github.com/xanzy/go-gitlab"
	"hidevops.io/hioak/starter/scm/gitlab"
	"hidevops.io/hioak/starter/scm/gitlab/fake"
	"os"
	"testing"
)

func TestListGroupMembers(t *testing.T) {
	fs := new(fake.GroupsService)
	cli := &fake.Client{
		GroupsService: fs,
	}
	s := gitlab.NewGroupMember(func(url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	gra := &gogitlab.GroupMember{
		ID:          100,
		Name:        "chulei",
		AccessLevel: gogitlab.DeveloperPermissions,
	}
	var gro []*gogitlab.GroupMember
	gro = append(gro, gra)
	gr := new(gogitlab.Response)
	gid := 4
	fs.On("ListGroupMembers", gid, nil, nil).Return(gro, gr, nil)

	_, err := s.ListGroupMembers(os.Getenv(""), "", gid, 100)
	assert.Equal(t, nil, err)
}

func TestGetGroupMember(t *testing.T) {
	fs := new(fake.GroupsService)
	cli := &fake.Client{
		GroupsService: fs,
	}
	s := gitlab.NewGroupMember(func(url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	gra := &gogitlab.GroupMember{
		ID:   100,
		Name: "chulei",
	}
	var gro []*gogitlab.GroupMember
	gro = append(gro, gra)
	gr := new(gogitlab.Response)
	gid := 4
	fs.On("ListGroupMembers", gid, nil, nil).Return(gro, gr, nil)
	_, err := s.GetGroupMember(os.Getenv(""), "", gid, 100)
	assert.Equal(t, nil, err)
}
