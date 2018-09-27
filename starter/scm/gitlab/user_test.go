package gitlab

import (
	"github.com/hidevopsio/hioak/starter/scm/gitlab/fake"
	"github.com/stretchr/testify/assert"
	"github.com/xanzy/go-gitlab"
	"os"
	"testing"
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
