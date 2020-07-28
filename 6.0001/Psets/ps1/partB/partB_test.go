package main

import (
	"testing"
)

func TestCountMonthB(t *testing.T) {
	// pdf上复制过来的字段带有特殊字符但是看上去一样一直报错,坑爹啊!!!!!!!!! ->semiAnnualRaise
	var tests = []struct {
		annualSalary    float64
		portionSaved    float64
		totalCost       float64
		semiAnnualRaise float64
		month           int
	}{
		{120000, .05, 500000, .03, 142},
		{80000, .1, 800000, .03, 159},
		{75000, .05, 1500000, .05, 261},
	}

	for _, test := range tests {
		actual := CountMonthB(test.annualSalary, test.portionSaved, test.totalCost, test.semiAnnualRaise)
		if actual != test.month {
			t.Errorf("Get %d; expected %d", actual, test.month)
		}
	}
}
