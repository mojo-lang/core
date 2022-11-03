package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlCodec(t *testing.T) {
	url := Url{
		Scheme: "https",
		Authority: &Url_Authority{
			Host: "mojo-lang.org",
		},
		Path:  "root",
		Query: NewUrlQuery(),
	}

	url.GetQuery().Add("test", "value")
	assert.Equal(t, "https://mojo-lang.org/root?test=value", url.Format())
}
