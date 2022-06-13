package v2

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"strings"
)

type Client interface {
	Auth(ctx context.Context) (*AuthResponse, error)
	DeliveryPoints(ctx context.Context, input *DeliveryPointsRequest) (*DeliveryPointsResponse, error)
	Regions(ctx context.Context, input *RegionsRequest) (*RegionsResponse, error)
	Cities(ctx context.Context, input *CitiesRequest) (*CitiesResponse, error)
	CalculatorTrafiffList(ctx context.Context, input *CalculatorTrafiffListRequest) (*CalculatorTrafiffListResponse, error)
	OrderRegister(ctx context.Context, input *OrderRegisterRequest) (*OrderRegisterResponse, error)
	OrderStatus(ctx context.Context, uuid string) (*OrderStatusResponse, error)
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

func (c *clientImpl) buildUri(p string, values interface{}) string {
	v, _ := query.Values(values)
	return strings.TrimRight(fmt.Sprintf(
		"%s/%s?%s",
		strings.TrimRight(c.opts.Endpoint, "/"),
		strings.TrimLeft(p, "/"),
		v.Encode(),
	), "?")
}
