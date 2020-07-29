package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsTriangular(t *testing.T) {
	fmt.Println("test for TestIsTriangular")
	var tests = []struct {
		num int
		res bool
	}{
		{1, true},
		{3, true},
		{10, true},
		{11, false},
	}

	for _, test := range tests {
		actual := IsTriangular(test.num)
		if actual != test.res {
			t.Errorf("Get %t; expected %t", actual, test.res)
		}
	}
}

func TestGetPrimeList(t *testing.T) {
	fmt.Println("test for TestGetPrimeList")
	var tests = []struct {
		num int
		res []int
	}{
		{6, []int{2, 3, 5}},
		{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}

	for _, test := range tests {
		actual := GetPrimeList(test.num)
		equal := reflect.DeepEqual(actual, test.res)
		if !equal {
			fmt.Println(actual, test.res)
			t.Errorf("Fail")
		}
	}
}

func TestCipher(t *testing.T) {
	fmt.Println("test for TestCipher")
	var tests = []struct {
		from, to, target, res string
		dict                  map[string]string
	}{
		{"abcd", "dcba", "dab", "adc", map[string]string{"a": "d", "b": "c", "d": "a", "c": "b"}},
	}

	for _, test := range tests {
		actual, text := Cipher(test.from, test.to, test.target)
		dictEqual := reflect.DeepEqual(actual, test.dict)
		textEqual := text == test.res
		if !(dictEqual && textEqual) {
			t.Errorf("Fail")
		}
	}
}
