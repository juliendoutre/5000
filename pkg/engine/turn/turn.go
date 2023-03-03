package turn

import (
	"github.com/juliendoutre/5000/pkg/engine/player"
	"github.com/juliendoutre/5000/pkg/engine/round"
)

type Turn struct {
	Rounds []*round.Round
	Score  uint
}

func Play(player player.Player) (*Turn, error) {
	turn := &Turn{}

	player.Logger().Log("Starting a new turn!")
	defer func() {
		player.Logger().Log("The turn ends, your final score is %d!", turn.Score)
	}()

	n := uint(5)

	for n > 0 {
		round, err := round.Play(n, player)
		if err != nil {
			return nil, err
		}

		turn.Rounds = append(turn.Rounds, round)

		// The player drawed a garbage.
		if round.PlayerDecision == nil {
			turn.Score = 0
			return turn, nil
		}

		turn.Score += round.FinalHand.Score()

		// The player chose to roll all 5 dices again.
		if round.PlayerDecision.All && round.InitialHand.AllMarked() {
			n = 5
			continue
		}

		// The player chose to stop.
		if round.PlayerDecision.IsEmpty() {
			return turn, nil
		}

		// The player will draw the next round with less dices.
		n = round.PlayerDecision.Size()
	}

	return turn, nil
}
