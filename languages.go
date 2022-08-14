package deepl

import (
	"context"

	"github.com/woodstock-tokyo/woodstock-deepl-go/params"
	"github.com/woodstock-tokyo/woodstock-deepl-go/types"
)

// Languages calls the languages API with type "target" of the Deepl API.
func (c *Client) TargetLanguages(ctx context.Context) (*types.TargetLanguagesResponse, *types.ErrorResponse, error) {
	params := &params.LanguagesParams{LangType: types.LangTypeTarget}
	res := types.TargetLanguagesResponse{}
	errRes, err := languages(ctx, c, params, &res)
	return &res, errRes, err
}

// Languages calls the languages API with type "source" of the Deepl API.
func (c *Client) SourceLanguages(ctx context.Context) (*types.SourceLanguagesResponse, *types.ErrorResponse, error) {
	params := &params.LanguagesParams{LangType: types.LangTypeSource}
	res := types.SourceLanguagesResponse{}
	errRes, err := languages(ctx, c, params, &res)
	return &res, errRes, err
}

func languages(ctx context.Context, c *Client, params *params.LanguagesParams, res any) (*types.ErrorResponse, error) {
	endpoint := c.EndpointBase + types.EndpointLanguages
	params.SetAuthnKey(c.AuthenticationKey)
	requester := NewRequester(endpoint, params)

	errRes, err := requester.Exec(res)
	if err != nil {
		return nil, err
	}
	if errRes != nil {
		return errRes, nil
	}

	return nil, nil
}
