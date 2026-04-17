package types

import (
	"net/http"
	"strings"
)

type RequestParams interface {
	SetAuthnKey(k string)
	Headers() http.Header
	Body() (*strings.Reader, error)
}
