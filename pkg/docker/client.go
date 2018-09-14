package docker

import (
	"github.com/docker/docker/client"
	"net/http"
	"github.com/docker/docker/api/types"
	"io/ioutil"
	"bytes"
	"encoding/json"
)

type Client struct {
	// scheme sets the scheme for the client
	scheme string
	// host holds the server address to connect to
	host string
	// proto holds the client protocol i.e. unix.
	proto string
	// addr holds the client address.
	addr string
	// basePath holds the path to prepend to the requests.
	basePath string
	// client used to send and receive http requests.
	client *http.Client
	// version of the server to talk to.
	version string
	// custom http headers configured by users.
	customHTTPHeaders map[string]string
	// manualOverride is set to true when the version was set by users.
	manualOverride bool
}

func NewClient() (*http.Client, error) {
	cli, err := client.NewEnvClient()
	c := newMockClient(errorMock(http.StatusUnauthorized, "Unauthorized error"))
	// get the fake ClientSet for testing
	cl := &Client{
		client: c,
	}
	return cli, err
}


func newMockClient(doer func(*http.Request) (*http.Response, error)) *http.Client {
	return &http.Client{
		Transport: transportFunc(doer),
	}
}

type transportFunc func(*http.Request) (*http.Response, error)


func errorMock(statusCode int, message string) func(req *http.Request) (*http.Response, error) {
	return func(req *http.Request) (*http.Response, error) {
		header := http.Header{}
		header.Set("Content-Type", "application/json")

		body, err := json.Marshal(&types.ErrorResponse{
			Message: message,
		})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: statusCode,
			Body:       ioutil.NopCloser(bytes.NewReader(body)),
			Header:     header,
		}, nil
	}
}