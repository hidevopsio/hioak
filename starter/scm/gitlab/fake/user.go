package fake

import (
	"github.com/stretchr/testify/mock"
	"github.com/xanzy/go-gitlab"
)

type UsersService struct {
	mock.Mock
}

func (c *UsersService) CurrentUser(options ...gitlab.OptionFunc) (*gitlab.User, *gitlab.Response, error) {
	args := c.Called(nil)
	return args[0].(*gitlab.User), args[1].(*gitlab.Response), args.Error(2)
}
