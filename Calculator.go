package main

import "math"

func calculate(principalBorrowed float64, mortgageRate float64, years float64) float64 {
	var n = countPayments(years)
	var r = assumePayments(mortgageRate)

	return principalBorrowed * (r / 100 * math.Pow(1+r/100, n)) / (math.Pow(1+r/100, n) - 1)
}

func countPayments(Y float64) float64 {
	return math.Ceil(Y * 12) // Calculate payments from years
}

func assumePayments(apr float64) float64 {
	return apr / 12 // assuming monthly payments
}
