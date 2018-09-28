package fake

import (
	"github.com/xanzy/go-gitlab"
	"github.com/stretchr/testify/mock"
)

type UsersService struct {
	mock.Mock
}
func (c *UsersService) CurrentUser(options ...gitlab.OptionFunc) (*gitlab.User, *gitlab.Response, error) {
	args := c.Called(nil)
	return args[0].(*gitlab.User), args[1].(*gitlab.Response), args.Error(2)
}

