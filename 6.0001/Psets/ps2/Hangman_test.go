package main

import (
	"fmt"
	"testing"
)

func TestIsWordGuessed(t *testing.T) {
	fmt.Println("test for TestIsWordGuessed")

	var tests = []struct {
		selectedWord  string
		letterGuessed []string
		result        bool
	}{
		{"apple", []string{"a", "c", "p", "l", "e"}, true},
		{"apple", []string{"e", "i", "k", "p", "r", "s"}, false},
	}

	for _, test := range tests {
		actual := IsWordGuessed(test.selectedWord, test.letterGuessed)
		if actual != test.result {
			t.Errorf("Get %t; expected %t", actual, test.result)
		}
	}
}

func TestGetWordGuessed(t *testing.T) {
	fmt.Println("test for TestGetWordGuessed")
	var tests = []struct {
		selectedWord  string
		letterGuessed []string
		result        string
	}{
		{"apple", []string{"a", "p", "p", "l", "e"}, "apple"},
		{"apple", []string{"a", "c", "c", "l", "e"}, "a__le"},
	}

	for _, test := range tests {
		actual := GetWordGuessed(test.selectedWord, test.letterGuessed)
		if actual != test.result {
			t.Errorf("Get %s; expected %s", actual, test.result)
		}
	}
}

func TestGetAvaliableLetters(t *testing.T) {
	fmt.Println("test for TestGetAvaliableLetters")
	var tests = []struct {
		letterGuessed []string
		result        string
	}{
		{[]string{"a", "p", "p", "l", "e"}, "bcdfghijkmnoqrstuvwxyz"},
		{[]string{"a", "c", "p", "l", "e"}, "bdfghijkmnoqrstuvwxyz"},
	}

	for _, test := range tests {
		actual := GetAvaliableLetters(test.letterGuessed)
		if actual != test.result {
			t.Errorf("Get %s; expected %s", actual, test.result)
		}
	}
}

func TestShowPossibleMatches(t *testing.T) {
	fmt.Println("test for TestShowPossibleMatches")
	var tests = []struct {
		wordGuessed string
		wordList    []string
		result      string
	}{
		{"ap___", []string{"apple", "atilt", "apk"}, "apple "},
		{"a____", []string{"apple", "atilt", "apk"}, "apple atilt "},
	}

	for _, test := range tests {
		actual := showPossibleMatches(test.wordGuessed, test.wordList)
		if actual != test.result {
			t.Errorf("Get %s; expected %s", actual, test.result)
		}
	}
}
