package calc

import "math"

// CentralMoment returns the rth sample central moment.
// See https://en.wikipedia.org/wiki/Moment_(mathematics)#Sample_moments.
func CentralMoment(vals []float64, r int) float64 {
	mean := Mean(vals)
	rf := float64(r)
	sum := 0.0
	for _, val := range vals {
		dev := val - mean
		sum += math.Pow(dev, rf)
	}
	return sum / float64(len(vals)-1)
}
