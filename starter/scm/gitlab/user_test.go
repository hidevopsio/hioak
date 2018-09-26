package gitlab

import (
	"testing"
		"github.com/stretchr/testify/assert"
	"os"
	"github.com/xanzy/go-gitlab"
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
)

func TestUser_GetUser(t *testing.T) {
	baseUrl := os.Getenv("SCM_URL")
	s := fake.NewClient("")
	s.On("SetBaseURL", nil).Return(nil)
	user := &gitlab.User{
		Name: "chulei",
	}
	resp := new(gitlab.Response)
	s.On("CurrentUser", nil).Return(user, resp, nil)
	u := NewUser(s)
	_, err := u.GetUser(baseUrl, os.Getenv("Token"))
	assert.Equal(t, nil, err)
}
