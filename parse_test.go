package securitytxt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)

	f, _ := FromString(`
		Contact: email@example.com
		Signature: https://example.txt.sig
	`)

	assert.Len(f.Errors, 1)
	assert.Equal("unknown key \"Signature\"", f.Errors[0].Error())
}
