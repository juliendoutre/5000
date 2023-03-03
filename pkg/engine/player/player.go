package player

import (
	"github.com/juliendoutre/5000/pkg/engine/decision"
	"github.com/juliendoutre/5000/pkg/engine/hand"
	"github.com/juliendoutre/5000/pkg/engine/logger"
)

type Player interface {
	// A Player must decides if it wants to continue playing or not based on a new hand.
	Decide(hand *hand.Hand) decision.Decision

	// A Player needs to expose a logger so that the game engine can inform it about the game state.
	Logger() logger.Logger
}
