package client

import (
	"github.com/xanzy/go-gitlab"
)

type Interface interface {
	SetBaseURL(urlStr string) error
	GetSession(opt *gitlab.GetSessionOptions, options ...gitlab.OptionFunc) (*gitlab.Session, *gitlab.Response, error)
}
