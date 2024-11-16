package domain

import (
	"math"
)

func CalculateRating(positiveReviews, totalReviews int) float64 {
	if totalReviews == 0 {
		return 0
	}
	reviewScore := float64(positiveReviews) / float64(totalReviews)

	logValue := math.Log10(float64(totalReviews + 1))
	decayFactor := math.Pow(2, -logValue)

	rating := reviewScore - (reviewScore-0.5)*decayFactor
	return rating
}
