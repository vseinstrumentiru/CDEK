package v2

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Credentials struct {
	ClientID     string
	ClientSecret string
}

func (c *Credentials) UrlValues() url.Values {
	data := url.Values{}

	data.Set("grant_type", "client_credentials")
	data.Set("client_id", c.ClientID)
	data.Set("client_secret", c.ClientSecret)

	return data
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	Jti         string `json:"jti"`
}

func (c *clientImpl) getAccessToken(ctx context.Context) (string, error) {
	d, err := time.ParseDuration(fmt.Sprintf("%ds", c.expireIn))
	if err != nil {
		return "", err
	}
	isExpired := time.Now().After(time.Now().Add(d))

	if len(c.accessToken) == 0 || isExpired {
		resp, err := c.Auth(ctx)
		if err != nil {
			return "", err
		}
		c.accessToken = resp.AccessToken
		c.expireIn = resp.ExpiresIn
	}

	return c.accessToken, nil
}

func (c *clientImpl) Auth(ctx context.Context) (*AuthResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost,
		c.buildUri("/v2/oauth/token"),
		strings.NewReader(c.opts.Credentials.UrlValues().Encode()),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return jsonReq[AuthResponse](req)
}
