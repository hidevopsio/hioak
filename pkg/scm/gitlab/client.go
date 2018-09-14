package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"net/http"
	"github.com/hidevopsio/hiboot/pkg/utils/gotest"
	"github.com/hidevopsio/hioak/pkg/scm/gitlab/fake"
)

type ClientInterface interface {
	SetBaseURL(baseUrl string) error
}


func NewClient(client *http.Client, token string) ClientInterface {


	// get the fake ClientSet for testing
	if gotest.IsRunning() {
		return fake.NewClient(client, token)
	}

	// get the real ClientSet
	clientSet := gitlab.NewClient(client, token)
	return clientSet
}
