package player

import (
	"github.com/juliendoutre/5000/pkg/engine/decision"
	"github.com/juliendoutre/5000/pkg/engine/hand"
	"github.com/juliendoutre/5000/pkg/engine/logger"
)

// DiceThresholdAI will keep all counters and will only roll at least Threshold dices again.
type DiceThresholdAI struct {
	Threshold uint
	Log       logger.Logger
}

func (d *DiceThresholdAI) Decide(hand *hand.Hand) decision.Decision {
	dec := decision.New()

	if hand.AllMarked() && d.Threshold <= 5 {
		dec.All = true
		return *dec
	}

	counters := hand.Count()
	for index, dice := range hand.Dices {
		if dice != 5 && dice != 1 && dice != counters.ThreeOfAKind {
			dec.Insert(uint(index))
		}
	}

	if dec.Size() < d.Threshold {
		return *decision.New()
	}

	return *dec
}

func (d *DiceThresholdAI) Logger() logger.Logger {
	return d.Log
}

var _ Player = &DiceThresholdAI{}
