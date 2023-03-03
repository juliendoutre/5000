package round

import (
	"github.com/juliendoutre/5000/pkg/engine/decision"
	"github.com/juliendoutre/5000/pkg/engine/hand"
	"github.com/juliendoutre/5000/pkg/engine/player"
)

type Round struct {
	InitialHand    *hand.Hand         // The hand the player drawed.
	PlayerDecision *decision.Decision // The dice indices the player chose to roll again if it could.
	FinalHand      *hand.Hand         // The round final hand.
}

func Play(n uint, player player.Player) (*Round, error) {
	initialHand, err := hand.Draw(n)
	if err != nil {
		return nil, err
	}

	round := &Round{InitialHand: initialHand}
	score := initialHand.Score()
	player.Logger().Log("You drawed %s (%d points)", initialHand, score)

	if score == 0 {
		player.Logger().Log("This is garbage, too bad!")
		round.FinalHand = initialHand
		return round, nil
	}

	playerDecision := player.Decide(initialHand)

	finalHand, err := initialHand.Exclude(&playerDecision)
	if err != nil {
		return nil, err
	}

	dicesToRollAgain, err := initialHand.Pick(&playerDecision)
	if err != nil {
		return nil, err
	}

	player.Logger().Log("You chose to roll the following dices again: %s", dicesToRollAgain)

	finalScore := finalHand.Score()
	// Invalid decision returned by the player (it should keep one counter). We ignore it.
	if finalScore == 0 {
		player.Logger().Log("Incorrect decision, you must keep at least one counter! It will be ignored.")
		round.PlayerDecision = decision.New()
		round.FinalHand = initialHand
		finalScore = score
	} else {
		round.PlayerDecision = &playerDecision
		round.FinalHand = finalHand
	}

	player.Logger().Log("Round's final hand is %s (%d)", round.FinalHand, finalScore)

	return round, nil
}
