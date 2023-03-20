package distribution

import "math"

type Discrete map[uint]float64

func (d Discrete) Expectation() float64 {
	expectation := 0.0

	for realization, probability := range d {
		expectation += float64(realization) * probability
	}

	return expectation
}

func (d Discrete) Variance() float64 {
	variance := 0.0
	expectation := d.Expectation()

	for realization, probability := range d {
		variance += probability * math.Pow(expectation-float64(realization), 2)
	}

	return variance
}

func (d Discrete) StandardDeviation() float64 {
	return math.Sqrt(d.Variance())
}
