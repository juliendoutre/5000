package distribution

import (
	"math"

	"github.com/juliendoutre/5000/pkg/engine/hand"
)

var (
	Scores = map[uint]Discrete{}
)

func init() {
	for i := uint(1); i < 6; i++ {
		Scores[i] = generateScoreDistribution(i)
	}
}

func generateScoreDistribution(n uint) map[uint]float64 {
	distrib := map[uint]float64{}

	counts := map[uint]uint{}
	countScores(n, counts, []uint{})

	for score, count := range counts {
		distrib[score] = float64(count) / math.Pow(6, float64(n))
	}

	return distrib
}

func countScores(n uint, counts map[uint]uint, dices []uint) {
	if n == 0 {
		h, _ := hand.From(dices...)
		score := h.Score()
		counts[score] += 1
	} else {
		for i := uint(1); i < 7; i++ {
			countScores(n-1, counts, append(dices, i))
		}
	}
}
