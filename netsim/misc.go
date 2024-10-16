package netsim

import (
	"math"
)

type Requirements struct {
	DistThreshold float64
}

// Calculates the Euclidean distance between two points
func Dist(loc1 [3]float64, loc2 [3]float64) float64 {
	sum := math.Pow(loc1[0]-loc2[0], 2) + math.Pow(loc1[1]-loc2[1], 2) + math.Pow(loc1[2]-loc2[2], 2)
	return math.Sqrt(sum)
}
