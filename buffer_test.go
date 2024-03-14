package sharedbuffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Buffer(t *testing.T) {
	var done bool
	sb := New(10, 2, func(*Buffer) {
		done = true
	})
	assert.Len(t, sb.Bytes(), 10)
	assert.EqualValues(t, sb.RC(), 2)
	assert.False(t, done)
	sb.Done()
	assert.False(t, done)
	sb.Done()
	assert.Len(t, sb.Bytes(), 10)
	assert.EqualValues(t, sb.RC(), 0)
	assert.True(t, done)
}
