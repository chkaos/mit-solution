package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetWordScore(t *testing.T) {
	fmt.Println("test for TestIsWordGuessed")

	var tests = []struct {
		word  WordScore
		score int
	}{
		{WordScore{"", 7}, 0},
		{WordScore{"it", 7}, 2},
		{WordScore{"WaYbILl", 7}, 735},
		{WordScore{"FORK", 4}, 308},
	}

	for _, test := range tests {
		actual := GetWordScore(test.word)
		if actual != test.score {
			t.Errorf("Get %d; expected %d", actual, test.score)
		}
	}
}

func TestGetFrequencyDict(t *testing.T) {
	fmt.Println("test for TestGetFrequencyDict")

	var tests = []struct {
		word   string
		result map[string]int
	}{
		{"hello", map[string]int{"h": 1, "e": 1, "l": 2, "o": 1}},
	}

	for _, test := range tests {
		dict := GetFrequencyDict(test.word)
		actual := reflect.DeepEqual(dict, test.result)
		if !actual {
			t.Errorf("Fail")
		}
	}
}

func TestDisplayHand(t *testing.T) {
	fmt.Println("test for TestDisplayHand")

	var tests = []struct {
		dict   map[string]int
		result string
	}{
		{map[string]int{"h": 1, "e": 1, "l": 2, "o": 1}, "h e l l o"},
		{map[string]int{"b": 3, "e": 1, "l": 2, "a": 1}, "b b b e l l a"},
		{map[string]int{"j": 2, "o": 1, "l": 1, "w": 1, "n": 2}, "j j o l w n n"},
	}

	for _, test := range tests {
		actual := DisplayHand(test.dict)

		if actual != test.result {
			t.Errorf("Get %s; expected %s", actual, test.result)
		}
	}
}

func TestUpdateHand(t *testing.T) {
	fmt.Println("test for TestUpdateHand")

	var tests = []struct {
		dict   map[string]int
		word   string
		result map[string]int
	}{
		{map[string]int{"j": 2, "o": 1, "l": 1, "w": 1, "n": 2}, "jolly", map[string]int{"j": 1, "w": 1, "n": 2}},
	}

	for _, test := range tests {
		dict := UpdateHand(test.dict, test.word)
		equal := reflect.DeepEqual(dict, test.result)
		if !equal {
			fmt.Println(dict, test.result)
			t.Errorf("Fail")
		}
	}
}

func TestIsValidWord(t *testing.T) {
	fmt.Println("test for TestIsValidWord")

	var tests = []struct {
		dict   map[string]int
		result bool
	}{
		{map[string]int{"h": 1, "e": 1, "l": 2, "o": 1}, true},
		{map[string]int{"j": 2, "o": 1, "l": 1, "w": 1, "n": 2}, false},
	}

	for _, test := range tests {
		actual := isValidWord(test.dict)

		if actual != test.result {
			t.Errorf("Get %t; expected %t", actual, test.result)
		}
	}
}
