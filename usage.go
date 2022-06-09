package deepl

import (
	"context"

	"github.com/woodstock-tokyo/woodstock-deepl-go/params"
	"github.com/woodstock-tokyo/woodstock-deepl-go/types"
)

// Usage calls the translate text API of the Deepl API.
func (c *Client) Usage(ctx context.Context, params *params.UsageParams) (*types.UsageResponse, *types.ErrorResponse, error) {
	var res types.UsageResponse

	endpoint := c.EndpointBase + types.EndpointUsage
	params.SetAuthnKey(c.AuthenticationKey)
	requester := NewRequester(endpoint, params)

	errRes, err := requester.Exec(&res)
	if err != nil {
		return nil, nil, err
	}
	if errRes != nil {
		return nil, errRes, nil
	}

	return &res, nil, nil
}
