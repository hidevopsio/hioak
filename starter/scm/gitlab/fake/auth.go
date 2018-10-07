package fake

import (
	"github.com/stretchr/testify/mock"
	"github.com/xanzy/go-gitlab"
)

type SessionService struct {
	mock.Mock
}

func (c *SessionService) GetSession(opt *gitlab.GetSessionOptions, options ...gitlab.OptionFunc) (*gitlab.Session, *gitlab.Response, error) {
	args := c.Called(nil, nil)
	return args[0].(*gitlab.Session), args[1].(*gitlab.Response), args.Error(2)
}
