package hand

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/juliendoutre/5000/pkg/engine/decision"
)

type Hand struct {
	Dices [5]uint // A Hand is composed of the results of between 1 and 5 dices.
	N     uint    // Number of dices in the Hand. This way we can use an array instead of a slice to store dices results.
}

type Counters struct {
	Counts       [6]uint // A dice takes its value between 1 and 6 (included). For performance reason, we index them in an 6-sized array with a -1 offset.
	ThreeOfAKind uint    // There's at most one three-of-kind in one Hand. 0 indicates no three-of-a-kind.
}

func Draw(n uint) (*Hand, error) {
	if n > 5 {
		return nil, fmt.Errorf("%d is strictly greater than 5", n)
	}

	dices := make([]uint, n)
	for i := uint(0); i < n; i++ {
		dices[i] = roll()
	}

	return From(dices...)
}

func roll() uint {
	return uint(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(6)) + 1
}

func From(dices ...uint) (*Hand, error) {
	if len(dices) > 5 {
		return nil, fmt.Errorf("%d is strictly greater than 5", len(dices))
	}

	hand := New()
	hand.N = uint(len(dices))

	copy(hand.Dices[:], dices)

	return hand, nil
}

func New() *Hand {
	return &Hand{}
}

func (h Hand) String() string {
	return fmt.Sprintf("%+v", h.Dices[:h.N])
}

func (h Hand) Score() uint {
	counters := h.Count()

	score := counters.Counts[0]*100 + counters.Counts[4]*50

	if counters.ThreeOfAKind == 1 {
		score += 700
	} else if counters.ThreeOfAKind == 5 {
		score += 350
	} else {
		score += counters.ThreeOfAKind * 100
	}

	return score
}

func (h Hand) Count() Counters {
	counters := Counters{}

	for i := uint(0); i < h.N; i++ {
		counters.Counts[h.Dices[i]-1] += 1
	}

	for value, count := range counters.Counts {
		if count >= 3 {
			counters.ThreeOfAKind = uint(value) + 1
		}
	}

	return counters
}

func (h *Hand) Exclude(decision *decision.Decision) (*Hand, error) {
	if decision.IsEmpty() {
		return h, nil
	}

	finalHand := []uint{}

	for index := uint(0); index < h.N; index++ {
		if !decision.DoesContain(uint(index)) {
			finalHand = append(finalHand, h.Dices[index])
		}
	}

	return From(finalHand...)
}

func (h Hand) Pick(decision *decision.Decision) (*Hand, error) {
	finalHand := []uint{}

	for index := uint(0); index < h.N; index++ {
		if decision.DoesContain(uint(index)) {
			finalHand = append(finalHand, h.Dices[index])
		}
	}

	return From(finalHand...)
}

func (h Hand) AllMarked() bool {
	counters := h.Count()

	count := counters.Counts[0] + counters.Counts[4]

	if counters.ThreeOfAKind != 0 {
		count += 3
	}

	return count >= h.N
}
