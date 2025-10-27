package securitytxt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpireDate(t *testing.T) {
	assert := assert.New(t)

	f, _ := FromString(`Expires: 2026-01-31T22:00:00.000Z`)

	assert.NotNil(f.Expiration())
	assert.Equal(f.Expiration().Year(), 2026)

	f, _ = FromString(`Contact: example@example.com`)

	assert.Nil(f.Expiration())
	assert.False(f.Expired())
}
