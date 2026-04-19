package securitytxt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpireDate(t *testing.T) {
	assert := assert.New(t)

	// Missing value
	f, _ := FromString(`Contact: example@example.com`)

	assert.Nil(f.Expiration())
	assert.Equal(f.DaysUntilExpiration(), 0)
	assert.False(f.Expired())

	// Empty value
	f, _ = FromString(`Expires:`)

	assert.Nil(f.Expiration())
	assert.False(f.Expired())

	// Valid value
	f, _ = FromString(`Expires: 2026-01-31T22:00:00.000Z`)

	assert.NotNil(f.Expiration())
	assert.Equal(f.Expiration().Year(), 2026)

	f, _ = FromString(`Contact: example@example.com`)

	assert.Nil(f.Expiration())
	assert.False(f.Expired())
}
