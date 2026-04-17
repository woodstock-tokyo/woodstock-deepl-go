package deepl

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/woodstock-tokyo/woodstock-deepl-go/params"
)

func TestRequesterPostSetsAuthorizationHeader(t *testing.T) {
	var gotAuthorization string
	var gotContentType string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuthorization = r.Header.Get("Authorization")
		gotContentType = r.Header.Get("Content-Type")
		_, _ = io.WriteString(w, `{"translations":[]}`)
	}))
	defer server.Close()

	p := &params.TranslateTextParams{
		Text:       []string{"hello"},
		TargetLang: "JA",
	}
	p.SetAuthnKey("test-authn-key")

	requester := NewRequester(server.URL, p)

	errRes, err := requester.Exec(&struct{}{})

	assert.NoError(t, err)
	assert.Nil(t, errRes)
	assert.Equal(t, "DeepL-Auth-Key test-authn-key", gotAuthorization)
	assert.Equal(t, contentType, gotContentType)
}
