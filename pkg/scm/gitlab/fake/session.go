package fake

import (
	"github.com/xanzy/go-gitlab"
	"github.com/stretchr/testify/mock"
)

type Session struct {
	mock.Mock
}

func (s *Session) GetSession(opt *gitlab.GetSessionOptions, options ... gitlab.OptionFunc) (*gitlab.Session, *gitlab.Response, error) {
	args := s.Called(nil, nil)
	return args[0].(*gitlab.Session),  args[1].(*gitlab.Response), args.Error(2)
}
