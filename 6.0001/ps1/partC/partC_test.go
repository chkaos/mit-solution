package main

import (
	"testing"
)

// Note: There are multiple right ways to implement bisection search/number of steps so your
// results may not perfectly match those of the test case.
func TestCountMonth(t *testing.T) {

	var tests = []struct {
		startingSalary float64
		savingRate     float64
		step           int
	}{
		{150000, 0.4411, 12},
		{300000, .2206, 9},
		{10000, .2206, -1},
	}

	for _, test := range tests {
		sRate, sStep := CountMonth(test.startingSalary)
		if sRate != test.savingRate {
			t.Errorf("Get %f; expected %f", sRate, test.savingRate)
		}
		if sStep != test.step {
			t.Errorf("Get %d; expected %d", sStep, test.step)
		}

	}
}
