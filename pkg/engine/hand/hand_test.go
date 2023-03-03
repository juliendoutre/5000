package hand

import (
	"testing"

	"github.com/juliendoutre/5000/pkg/engine/decision"
	"github.com/stretchr/testify/assert"
	"pgregory.net/rapid"
)

func TestRoll(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		n := roll()
		assert.GreaterOrEqual(t, n, uint(1))
		assert.LessOrEqual(t, n, uint(6))
	})
}

func TestDraw(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		n := uint(rapid.IntRange(1, 5).Draw(t, ""))

		hand, err := Draw(n)
		assert.NoError(t, err)
		assert.Equal(t, hand.N, n)
	})
}

func TestCount(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		n := uint(rapid.IntRange(1, 5).Draw(t, ""))

		hand, err := Draw(n)
		assert.NoError(t, err)
		counters := hand.Count()

		totalCount := uint(0)
		for _, count := range counters.Counts {
			totalCount += count
		}
		assert.Equal(t, totalCount, n)

		if counters.ThreeOfAKind != 0 {
			assert.GreaterOrEqual(t, hand.N, uint(3))
		}
	})
}

func TestScore(t *testing.T) {
	testCases := map[string]struct {
		dices []uint
		score uint
	}{
		"garbage": {
			dices: []uint{2, 3, 4, 2, 6},
			score: 0,
		},
		"three-of-a-kind": {
			dices: []uint{2, 2, 3, 3, 3},
			score: 300,
		},
		"three-of-a-kind with additional counters": {
			dices: []uint{1, 1, 1, 5, 5},
			score: 1100,
		},
		"three-of-a-kind with additional counters bis": {
			dices: []uint{1, 1, 5, 5, 5},
			score: 700,
		},
		"three-of-a-kind with additional counters ter": {
			dices: []uint{1, 3, 3, 3, 2},
			score: 400,
		},
	}

	for testCaseName, testCaseData := range testCases {
		t.Run(testCaseName, func(t *testing.T) {
			hand, err := From(testCaseData.dices...)
			assert.NoError(t, err)
			assert.Equal(t, hand.Score(), testCaseData.score)
		})
	}
}

func TestExclude(t *testing.T) {
	testCases := map[string]struct {
		dices             []uint
		decision          []uint
		expectedFinalHand []uint
	}{
		"empty decision": {
			dices:             []uint{1, 1, 2, 3, 4},
			decision:          []uint{},
			expectedFinalHand: []uint{1, 1, 2, 3, 4},
		},
		"pick one": {
			dices:             []uint{1, 1, 2, 3, 4},
			decision:          []uint{0},
			expectedFinalHand: []uint{1, 2, 3, 4},
		},
		"pick multiple": {
			dices:             []uint{1, 1, 2, 3, 4},
			decision:          []uint{0, 3, 4},
			expectedFinalHand: []uint{1, 2},
		},
		"pick all": {
			dices:             []uint{1, 1, 2, 3, 4},
			decision:          []uint{0, 1, 2, 3, 4},
			expectedFinalHand: []uint{},
		},
	}

	for testCaseName, testCaseData := range testCases {
		t.Run(testCaseName, func(t *testing.T) {
			hand, err := From(testCaseData.dices...)
			assert.NoError(t, err)
			decision, err := decision.From(testCaseData.decision...)
			assert.NoError(t, err)
			expectedFinalHand, err := From(testCaseData.expectedFinalHand...)
			assert.NoError(t, err)
			actualFinalHand, err := hand.Exclude(decision)
			assert.NoError(t, err)

			assert.Equal(t, actualFinalHand.Dices, expectedFinalHand.Dices)
		})
	}
}
