package params_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/woodstock-tokyo/woodstock-deepl-go/params"
)

func TestUsageSetAuthnKey(t *testing.T) {
	k := "test-authn-key"
	p := params.UsageParams{}
	p.SetAuthnKey(k)

	assert.Equal(t, k, p.AuthKey)
}

func TestUsageHeaders(t *testing.T) {
	p := params.UsageParams{}
	p.SetAuthnKey("test-authn-key")

	assert.Equal(t, "DeepL-Auth-Key test-authn-key", p.Headers().Get("Authorization"))
}

func TestUsageBody(t *testing.T) {
	cases := []struct {
		name   string
		params params.UsageParams
		expect string
	}{
		{
			name:   "normal",
			params: params.UsageParams{AuthKey: "test-authn-key"},
			expect: "",
		},
		{
			name:   "normal: with white space",
			params: params.UsageParams{AuthKey: "test key"},
			expect: "",
		},
		{
			name:   "normal: empty",
			params: params.UsageParams{},
			expect: "",
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
