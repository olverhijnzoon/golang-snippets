package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Net Present Value (NPV)")

	/*
		Net present value (NPV) is a financial metric used to evaluate the profitability of an investment or project. It is calculated by subtracting the initial cost of the investment from the sum of all the future cash flows that are expected to be generated by the investment, discounted at a given rate. The discount rate is used to account for the time value of money, which states that a dollar received today is worth more than a dollar received in the future. The NPV calculation is often used by businesses to determine whether an investment is worth pursuing, as a positive NPV indicates that the investment is expected to generate a return that is higher than the cost of capital.

		The formula used here is the following:

		NPV = \sum_{t=1}^{n}\frac{CF_t}{(1+r)^t} - C_0

	*/

	// Initial cost of the investment
	initialCost := 100000000.0

	// Discount rate, expressed as a decimal
	discountRate := 0.05

	// Sequence of future cash flows
	cashFlows := []float64{0, 1, -66600000, 42000000, 690000000}

	// Calculate the net present value
	var npv float64
	for i, cashFlow := range cashFlows {
		discountFactor := 1.0 / math.Pow(1+discountRate, float64(i+1))
		npv += cashFlow * discountFactor
	}
	npv -= initialCost

	// Print the result
	fmt.Printf("NPV = %.2f\n", npv)
}
