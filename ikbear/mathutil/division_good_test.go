package mathutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Division(t *testing.T) {
	_, err := Division(6, 2)
	assert.Nil(t, err)

	_, err2 := Division(6, 0)
	assert.NotNil(t, err2)
}
