package player

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/juliendoutre/5000/pkg/engine/decision"
	"github.com/juliendoutre/5000/pkg/engine/hand"
	"github.com/juliendoutre/5000/pkg/engine/logger"
)

// Console is designed for human players to play the game in a terminal.
type Console struct{}

func (c *Console) Decide(hand *hand.Hand) decision.Decision {
	if hand.AllMarked() {
		fmt.Print("All dices marked. Do you want to roll all fice dices again? (yes/No): ")
		var answer string
		fmt.Scanln(&answer)
		if strings.ToLower(answer) == "yes" {
			decision := decision.New()
			decision.All = true
			return *decision
		}
	}

	fmt.Print("Choose dices to roll again (you must keep at least one counter): ")
	var dices string
	fmt.Scanln(&dices)

	d := decision.New()
	for _, token := range strings.Split(dices, ",") {
		index, err := strconv.ParseUint(token, 10, 0)
		if err != nil {
			fmt.Printf("invalid input: %s. Please enter valid indices.\n", err)
			return *decision.New()
		}

		d.Insert(uint(index))
	}

	return *d
}

func (c *Console) Logger() logger.Logger {
	return &logger.Stdout{}
}

var _ Player = &Console{}
