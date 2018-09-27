package fake

import "github.com/stretchr/testify/mock"

type Client struct {
	mock.Mock
}

func NewClient() (*Client, error) {
	return &Client{}, nil
}
