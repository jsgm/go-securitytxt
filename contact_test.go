package securitytxt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContact(t *testing.T) {
	assert := assert.New(t)

	f, _ := FromString(`
		Contact: mailto:test@example.com
		Contact: https://example.com
		Contact: tel:+103054078449
		Contact: 

		Hiring: https://example.com/hiring
	`)

	assert.Equal(len(f.Contact), 4)

	c, e := f.PreferredContact()
	assert.True(c.IsEmail())
	assert.Nil(e)

	assert.True(f.Contact[0].IsEmail())
	assert.False(f.Contact[0].IsUnknown())
	assert.True(f.Contact[1].IsURL())
	assert.True(f.Contact[2].IsPhone())
	assert.True(f.Contact[3].IsUnknown())
}
