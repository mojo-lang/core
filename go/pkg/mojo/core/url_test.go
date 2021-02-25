package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlCodec(t *testing.T) {
	url := Url{
		Scheme: "https",
		Authority: &Url_Authority{
			Host: "mojo-lang.org",
		},
		Path:  "root",
		Query: NewQuery(),
	}

	url.GetQuery().Add("test", "value")
	assert.Equal(t, "https://mojo-lang.org/root?test=value", url.Encode())
}
