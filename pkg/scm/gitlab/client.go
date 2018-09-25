package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"net/http"
)

type ClientInterface interface {
	SetBaseURL(baseUrl string) error
	GetSession(opt *gitlab.GetSessionOptions, options ...gitlab.OptionFunc) (*gitlab.Session, *gitlab.Response, error)
}


func NewClient(token string) *gitlab.Client {
	// get the real ClientSet
	clientSet := gitlab.NewClient(&http.Client{}, token)
	return clientSet
}
