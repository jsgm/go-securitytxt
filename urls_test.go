package securitytxt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHiring(t *testing.T) {
	assert := assert.New(t)

	f, _ := FromString(`
		Contact: mailto:test@example.com
		Hiring: https://example.com
		Hiring: https://example.net
		Hiring: https://example.com
	`)

	assert.Equal(len(f.Hiring), 3)
	assert.Equal(f.Hiring.First().String(), "https://example.com")
}
