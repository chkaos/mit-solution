// plus The semi­annual salary raise (semi_annual_raise)

// 这里为了省时间不再做用户输入相关代码

package main

import (
	"fmt"
)

// CountMonthB
func CountMonthB(annualSalary, portionSaved, totalCost, semiAnnualRaise float64) (month int) {

	var currentSavings float64
	var portionDownPayment float64 = 0.25
	var investRate = 0.04 // current_savings*r/12

	fmt.Printf("Enter your annual salary: %.0f\n", annualSalary)
	fmt.Printf("Enter the percent of your salary to save, as a decimal: %.2f\n", portionSaved)
	fmt.Printf("Enter the cost of your dream home: %.0f\n", totalCost)
	fmt.Printf("Enter the percent of your semi annual raise, as a decimal: %.2f\n", semiAnnualRaise)

	DownPayment := totalCost * portionDownPayment
	monthPayment := annualSalary / 12

	for DownPayment > currentSavings {
		currentSavings += currentSavings*investRate/12 + portionSaved*monthPayment
		month++
		if month%6 == 0 {
			monthPayment *= (1 + semiAnnualRaise)
		}
	}

	fmt.Printf("Number of months: %d\n", month)
	return month
}

func main() {

}
