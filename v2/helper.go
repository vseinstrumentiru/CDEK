package v2

import (
	"context"
)

func HelperCitiesAll(ctx context.Context, c Client, input *CitiesRequest, first int) (*CitiesResponse, error) {
	resp := &CitiesResponse{}

	if input == nil {
		input = &CitiesRequest{Size: 500}
	}

	if input.Size > 500 {
		input.Size = 500
	}

	for {
		chunk, err := c.Cities(ctx, input)
		if err != nil {
			return nil, err
		}

		if len(*chunk) == 0 {
			break
		}

		*resp = append(*resp, *chunk...)

		if len(*resp) >= first {
			break
		}

		input.Page += 1
	}

	return resp, nil
}
