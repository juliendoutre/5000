package player

import (
	"github.com/juliendoutre/5000/pkg/engine/decision"
	"github.com/juliendoutre/5000/pkg/engine/hand"
	"github.com/juliendoutre/5000/pkg/engine/logger"
)

// ProbabilityThresholdAI will choose its next move based on the higher probability for the score distribution.
type ProbabilityThresholdAI struct {
	Threshold float64
	Log       logger.Logger
}

func (p *ProbabilityThresholdAI) Decide(hand *hand.Hand) decision.Decision {
	d := decision.New()

	return *d
}

func (p *ProbabilityThresholdAI) Logger() logger.Logger {
	return p.Log
}

var _ Player = &ProbabilityThresholdAI{}
