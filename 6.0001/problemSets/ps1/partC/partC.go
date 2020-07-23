// plus The semi­annual salary raise (semi_annual_raise)

// 这里为了省时间不再做用户输入相关代码

package main

import (
	"fmt"
	"math"
)

var portionDownPayment float64 = 0.25
var investRate = 0.04 // current_savings*r/12
var totalCost float64 = 1000000
var semiAnnualRaise = .07
var downPayment = totalCost * portionDownPayment

// CountMonth
func CountMonth(startingSalary float64) (portionSaved float64, step int) {
	var rateMax float64 = 10000
	var rateMin float64 = 0
	var currentSavings float64

	fmt.Printf("Enter your starting salary: %.0f\n", startingSalary)

	if startingSalary < 62523 {
		fmt.Printf("it is not possible to pay down payment in three years")
		return 0, -1
	}

	for math.Abs(downPayment-currentSavings) > 1000 {
		step++
		portionSaved = ((rateMax+rateMin)/2 + 0.5)
		monthSalaryPortion := startingSalary * portionSaved / 10000

		currentSavings = 0

		for i := 1; i < 37; i++ {
			currentSavings += monthSalaryPortion + (currentSavings * investRate / 12)
			if i%6 == 0 {
				monthSalaryPortion *= 1 + semiAnnualRaise
			}
		}

		if currentSavings < downPayment {
			rateMin = portionSaved
		} else {
			rateMax = portionSaved
		}
	}

	fmt.Printf("best saving rate: %f\n", portionSaved/10000)
	fmt.Printf("Steps in bisecton search: %d\n", step)

	return portionSaved / 10000, step
}

func main() {

}
