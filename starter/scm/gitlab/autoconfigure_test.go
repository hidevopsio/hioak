package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"net/http"
	"testing"
)

func TestAutoNewClient(t *testing.T)  {
	c := newConfiguration()
	c.NewClient()
	gitlab.NewClient(&http.Client{}, "")
}
