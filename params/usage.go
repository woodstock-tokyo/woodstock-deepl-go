package params

import (
	"net/http"
	"net/url"
	"strings"
)

// UsageParams is parameters struct for Usage API.
type UsageParams struct {
	AuthKey string
}

func (p *UsageParams) SetAuthnKey(k string) {
	p.AuthKey = k
}

func (p *UsageParams) Headers() http.Header {
	header := make(http.Header)
	if p.AuthKey != "" {
		header.Set("Authorization", "DeepL-Auth-Key "+p.AuthKey)
	}
	return header
}

func (p *UsageParams) Body() (*strings.Reader, error) {
	uv := url.Values{}

	return strings.NewReader(uv.Encode()), nil
}
