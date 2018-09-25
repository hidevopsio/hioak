package fake

import (
	"github.com/stretchr/testify/mock"
	"github.com/xanzy/go-gitlab"
)

type Client struct {
	mock.Mock

}
type tokenType int
const (
	privateToken tokenType = iota
	oAuthToken
)

func NewClient(token string) *Client {
	return &Client{

	}
}

func (c *Client) SetBaseURL(urlStr string) error {
	// Make sure the given URL end with a slash
	args := c.Called(nil)
	return args.Error(0)
}

func (c *Client) GetSession(opt *gitlab.GetSessionOptions, options ...gitlab.OptionFunc) (*gitlab.Session, *gitlab.Response, error)   {
	args := c.Called(nil, nil)
	return args[0].(*gitlab.Session), args[1].(*gitlab.Response), args.Error(2)
}