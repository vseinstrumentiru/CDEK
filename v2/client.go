package v2

import (
	"context"
	"fmt"
	"strings"
)

type Client interface {
	Auth(ctx context.Context) (*AuthResponse, error)
	DeliveryPoints(ctx context.Context) (*DeliveryPointsResponse, error)
}

type Options struct {
	Endpoint    string
	Credentials *Credentials
}

func NewClient(opts *Options) Client {
	return &clientImpl{opts: opts}
}

type clientImpl struct {
	opts        *Options
	accessToken string
	expireIn    int
}

func (c *clientImpl) buildUri(p string) string {
	return fmt.Sprintf(
		"%s/%s",
		strings.TrimRight(c.opts.Endpoint, "/"),
		strings.TrimLeft(p, "/"),
	)
}
