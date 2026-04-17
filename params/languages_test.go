package params_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/woodstock-tokyo/woodstock-deepl-go/params"
	"github.com/woodstock-tokyo/woodstock-deepl-go/types"
)

func TestLanguagesSetAuthnKey(t *testing.T) {
	k := "test-authn-key"
	p := params.LanguagesParams{}

	p.SetAuthnKey(k)

	assert.Equal(t, k, p.AuthKey)
}

func TestLanguagesHeaders(t *testing.T) {
	p := params.LanguagesParams{}
	p.SetAuthnKey("test-authn-key")

	assert.Equal(t, "DeepL-Auth-Key test-authn-key", p.Headers().Get("Authorization"))
}

func TestLanguagesBody(t *testing.T) {
	cases := []struct {
		name   string
		params params.LanguagesParams
		expect string
	}{
		{
			name:   "normal: langType=source",
			params: params.LanguagesParams{AuthKey: "test-authn-key", LangType: types.LangTypeSource},
			expect: "type=source",
		},
		{
			name:   "normal: langType=target",
			params: params.LanguagesParams{AuthKey: "test-authn-key", LangType: types.LangTypeTarget},
			expect: "type=target",
		},
		{
			name:   "normal: with white space",
			params: params.LanguagesParams{AuthKey: "test key", LangType: types.LangTypeSource},
			expect: "type=source",
		},
		{
			name:   "normal: empty",
			params: params.LanguagesParams{},
			expect: "type=",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()

			assert.NoError(tt, err)

			s := []byte{}
			for {
				b, err := r.ReadByte()
				if err == io.EOF {
					break
				}
				s = append(s, b)
			}

			assert.Equal(tt, c.expect, string(s))
		})
	}
}
