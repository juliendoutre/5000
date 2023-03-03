package player

import (
	"github.com/juliendoutre/5000/pkg/engine/decision"
	"github.com/juliendoutre/5000/pkg/engine/hand"
	"github.com/juliendoutre/5000/pkg/engine/logger"
)

// CautiousAI will keep all counters and will only roll at least Threshold dices again.
type CautiousAI struct {
	Threshold uint
	Log       logger.Logger
}

func (c *CautiousAI) Decide(hand *hand.Hand) decision.Decision {
	d := decision.New()

	if hand.AllMarked() && c.Threshold <= 5 {
		d.All = true
		return *d
	}

	counters := hand.Count()
	for index, dice := range hand.Dices {
		if dice != 5 && dice != 1 && dice != counters.ThreeOfAKind {
			d.Insert(uint(index))
		}
	}

	if d.Size() < c.Threshold {
		return *decision.New()
	}

	return *d
}

func (c *CautiousAI) Logger() logger.Logger {
	return c.Log
}

var _ Player = &CautiousAI{}
