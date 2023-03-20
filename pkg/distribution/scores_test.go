package distribution

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScores(t *testing.T) {
	for _, distrib := range Scores {
		sum := 0.0
		for _, probability := range distrib {
			sum += probability
		}

		assert.InDelta(t, sum, 1.0, math.Pow10(-6))
	}
}
