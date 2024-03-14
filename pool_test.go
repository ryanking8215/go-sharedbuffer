package sharedbuffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pool(t *testing.T) {
	pool := NewPool(1000)
	sb := pool.Get()
	sb.Add(3)
	assert.Len(t, sb.Bytes(), 1000)
	assert.EqualValues(t, sb.RC(), 3)
	sb.Done()
	sb.Done()
	sb.Done()
}
