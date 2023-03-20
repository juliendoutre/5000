package main

import (
	"fmt"
	"log"

	"github.com/juliendoutre/5000/pkg/distribution"
	"github.com/juliendoutre/5000/pkg/engine/logger"
	"github.com/juliendoutre/5000/pkg/engine/player"
	"github.com/juliendoutre/5000/pkg/engine/turn"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	repetitions = 10000
)

func main() {
	DiceThresholdAIs := [...]*player.DiceThresholdAI{
		{Threshold: 1, Log: &logger.NoOp{}},
		{Threshold: 2, Log: &logger.NoOp{}},
		{Threshold: 3, Log: &logger.NoOp{}},
		{Threshold: 4, Log: &logger.NoOp{}},
		{Threshold: 5, Log: &logger.NoOp{}},
		{Threshold: 6, Log: &logger.NoOp{}},
	}

	for _, ai := range DiceThresholdAIs {
		fmt.Printf("Running Dice Threshold AI with threshold %d...\n", ai.Threshold)

		scores := []float64{}
		total := 0.0

		for i := 0; i < repetitions; i++ {
			turn, err := turn.Play(ai)
			if err != nil {
				log.Fatal(err)
			}

			scores = append(scores, float64(turn.Score))
			total += float64(turn.Score)
		}

		histPlot(scores, fmt.Sprintf("dice_threshold_ai_%d", ai.Threshold))
		fmt.Printf("Average score: %f\n", total/repetitions)
	}

	for i, distrib := range distribution.Scores {
		xys := []plotter.XY{}
		for score, probability := range distrib {
			xys = append(xys, plotter.XY{X: float64(score), Y: probability})
		}

		histogramPlot(xys, fmt.Sprintf("score_probability_distribution_for_%d_dices", i))
	}
}

func histPlot(values plotter.Values, name string) error {
	p := plot.New()

	p.Title.Text = name

	hist, err := plotter.NewHist(values, 1000)
	if err != nil {
		return err
	}

	p.Add(hist)

	if err := p.Save(3*vg.Inch, 3*vg.Inch, fmt.Sprintf("img/%s.png", name)); err != nil {
		return err
	}

	return nil
}

func histogramPlot(xys plotter.XYs, name string) error {
	p := plot.New()

	p.Title.Text = name

	hist, err := plotter.NewHistogram(xys, 1000)
	if err != nil {
		return err
	}

	p.Add(hist)

	if err := p.Save(3*vg.Inch, 3*vg.Inch, fmt.Sprintf("img/%s.png", name)); err != nil {
		return err
	}

	return nil
}
