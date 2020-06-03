// Calculate the monthly payments of a fixed term mortgage over given Nth terms
// at a given interest rate.

package main

import (
	"fmt"
)

func main() {
	var (
		mortgageRate      float64
		principalBorrowed float64
		years             float64
		typeOfCredit      int
	)

	fmt.Println("Choose a type of loan. If it is a mortgage enter 0, if a regular loan enter 1: ")
	_, err := fmt.Scanf("%d", &typeOfCredit)
	validate(err)

	fmt.Println("Enter the principal owed ($): ")
	_, err = fmt.Scanf("%f", &principalBorrowed)
	validate(err)

	fmt.Println("Enter the mortgage rate or APR (%): ")
	_, err = fmt.Scanf("%f", &mortgageRate)
	validate(err)

	fmt.Println("Enter the length of the mortgage (years): ")
	_, err = fmt.Scanf("%f", &years)
	validate(err)

	A := calculate(principalBorrowed, mortgageRate, years)

	printResult(principalBorrowed, mortgageRate, years, A, typeOfCredit)
}

func printResult(principalBorrowed float64,
	mortgageRate float64, years float64, A float64, typeOfCredit int) {

	if typeOfCredit == 1 {
		fmt.Printf("Monthly regular loan payment for $%.2f at an APR of ", principalBorrowed)
	} else {
		fmt.Printf("Monthly mortgage payment for $%.2f at an APR of ", principalBorrowed)
	}

	fmt.Printf("%.2f%% over %.0f years:\n", mortgageRate, years)
	fmt.Printf("$%.2f \n", A)
}
