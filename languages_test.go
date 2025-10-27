package securitytxt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguages(t *testing.T) {
	assert := assert.New(t)

	sectxt, _ := FromString(`
		Contact: mailto:test@example.com
		Expires: Test
		Preferred-Languages: en, es, fr
	`)

	assert.Equal(sectxt.Language().Code, "en")
	assert.NotEqual(sectxt.Language().Code, "es")

	assert.True(sectxt.ContainsLanguage("en"))
	assert.False(sectxt.ContainsLanguage("de"))

	assert.Equal(len(sectxt.Languages()), 3)
	assert.NotEqual(len(sectxt.Languages()), 2)

	sectxt, _ = FromString(`
		Contact: mailto:test@example.com
		Expires: Test
		Preferred-Languages:
	`)
	assert.Equal(sectxt.Language().Code, "en")
	assert.Equal(len(sectxt.Languages()), 1)
	assert.True(sectxt.ContainsLanguage("en"))
	assert.True(sectxt.ContainsLanguage("en"))
}
