package decision

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pgregory.net/rapid"
)

func TestDecision(t *testing.T) {
	t.Run("empty decision", func(t *testing.T) {
		decision := Decision{}
		assert.Equal(t, decision.size, uint(0))
		assert.True(t, decision.IsEmpty())
		assert.Equal(t, decision.toSlice(), []uint{})
	})

	t.Run("any decision", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			input := rapid.SliceOfN(rapid.UintRange(0, 4), 1, 100).Draw(t, "")
			decision, err := From(input...)
			assert.NoError(t, err)
			assert.GreaterOrEqual(t, decision.size, uint(1))
			assert.LessOrEqual(t, decision.size, uint(5))
		})
	})
}
