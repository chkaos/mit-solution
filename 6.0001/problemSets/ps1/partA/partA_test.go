package main

import (
	"testing"
)

func TestCountMonth(t *testing.T) {

	var tests = []struct {
		annualSalary float64
		portionSaved float64
		totalCost    float64
		month        int
	}{
		{120000, .10, 1000000, 183},
		{80000, .15, 500000, 105},
	}

	for _, test := range tests {
		actual := CountMonth(test.annualSalary, test.portionSaved, test.totalCost)
		if actual != test.month {
			t.Errorf("Get %d; expected %d", actual, test.month)
		}
	}
}
