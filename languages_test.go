package securitytxt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguages(t *testing.T) {
	assert := assert.New(t)

	f, _ := FromString(`
		Contact: mailto:test@example.com
		Expires: Test
		Preferred-Languages: en, es, fr
	`)

	assert.Equal(f.Language().Code, "en")
	assert.NotEqual(f.Language().Code, "es")

	assert.True(f.ContainsLanguage("en"))
	assert.False(f.ContainsLanguage("de"))

	assert.Len(f.PreferredLanguages, 3)
	assert.NotEqual(len(f.PreferredLanguages), 2)

	f, _ = FromString(`
		Contact: mailto:test@example.com
		Expires: Test
		Preferred-Languages:
	`)
	assert.Equal(f.Language().Code, "en")
	assert.Equal(len(f.PreferredLanguages), 1)
	assert.True(f.ContainsLanguage("en"))
	assert.True(f.ContainsLanguage("en"))
}
