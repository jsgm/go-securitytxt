package securitytxt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)

	_, err := FromString("")
	assert.Error(err)
}
