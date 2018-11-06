package gitlab_test

import (
	"github.com/stretchr/testify/assert"
	gogitlab "github.com/xanzy/go-gitlab"
	"hidevops.io/hioak/starter/scm/gitlab"
	"hidevops.io/hioak/starter/scm/gitlab/fake"
	"os"
	"testing"
)

func TestUser_GetUser(t *testing.T) {
	fs := new(fake.UsersService)
	cli := &fake.Client{
		UsersService: fs,
	}
	s := gitlab.NewUser(func(url, token string) (client gitlab.ClientInterface) {
		return cli
	})
	user := &gogitlab.User{
		Name: "chulei",
	}
	resp := new(gogitlab.Response)
	fs.On("CurrentUser", nil).Return(user, resp, nil)
	_, err := s.GetUser("", os.Getenv("Token"))
	assert.Equal(t, nil, err)
}
