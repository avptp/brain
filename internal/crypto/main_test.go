package crypto

import (
	"math/rand"
	"testing"

	testify "github.com/stretchr/testify/assert"
)

func TestRandomBytes(t *testing.T) {
	assert := testify.New(t)
	length := randomInt(8, 16)

	t.Run("length", func(t *testing.T) {
		t.Parallel()

		bytes, err := RandomBytes(length)

		assert.Nil(err)
		assert.Equal(length, len(bytes))
	})

	t.Run("non-determinism", func(t *testing.T) {
		t.Parallel()

		a, err := RandomBytes(length)
		assert.Nil(err)

		b, err := RandomBytes(length)
		assert.Nil(err)

		assert.NotEqual(a, b)
	})
}

func randomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
