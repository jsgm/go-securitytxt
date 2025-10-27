package securitytxt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHiring(t *testing.T) {
	assert := assert.New(t)

	sectxt, _ := FromString(`
		Contact: mailto:test@example.com
		Hiring: https://example.com
		Hiring: https://example.net
		Hiring: https://example.com
	`)

	assert.Equal(len(sectxt.Hiring()), 3)
	assert.Equal(sectxt.Hiring().First().String(), "https://example.com")
}
