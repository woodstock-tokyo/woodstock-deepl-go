package params

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/woodstock-tokyo/woodstock-deepl-go/types"
)

// LanguagesParams is parameters struct for Languages API.
type LanguagesParams struct {
	AuthKey  string
	LangType types.LangType
}

func (p *LanguagesParams) SetAuthnKey(k string) {
	p.AuthKey = k
}

func (p *LanguagesParams) Headers() http.Header {
	header := make(http.Header)
	if p.AuthKey != "" {
		header.Set("Authorization", "DeepL-Auth-Key "+p.AuthKey)
	}
	return header
}

func (p *LanguagesParams) Body() (*strings.Reader, error) {
	uv := url.Values{}

	uv.Add("type", string(p.LangType))
	return strings.NewReader(uv.Encode()), nil
}
