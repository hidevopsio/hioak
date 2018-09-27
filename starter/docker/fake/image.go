package fake

import (
	"context"
	"github.com/docker/docker/api/types"
	"io"
	"net/http"
)

type serverResponse struct {
	body       io.ReadCloser
	header     http.Header
	statusCode int
}

func (c *Client) ImagePull(ctx context.Context, ref string, options types.ImagePullOptions) (io.ReadCloser, error) {
	args := c.Called(nil, nil, nil)
	serverResp := serverResponse{statusCode: -1}
	return serverResp.body, args.Error(1)
}

func (c *Client) ImageTag(ctx context.Context, imageID, ref string) error {
	args := c.Called(nil, nil, nil)
	return args.Error(0)
}

func (c *Client) ImagePush(ctx context.Context, ref string, options types.ImagePushOptions) (io.ReadCloser, error) {
	args := c.Called(nil, nil, nil)
	serverResp := serverResponse{statusCode: -1}
	return serverResp.body, args.Error(1)
}

func (c *Client) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error) {
	args := c.Called(nil, nil)
	return args[0].([]types.ImageSummary), args.Error(1)
}
