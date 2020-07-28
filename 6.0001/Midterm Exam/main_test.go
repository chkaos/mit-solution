package main

import (
	"reflect"
	"testing"
)

func TestCount7(t *testing.T) {
	// pdf上复制过来的字段带有特殊字符但是看上去一样一直报错,坑爹啊!!!!!!!!! ->semiAnnualRaise
	var tests = []struct {
		num   int
		count int
	}{
		{717, 2},
		{6, 0},
		{7346772394047, 4},
	}

	for _, test := range tests {
		actual := Count7(test.num)
		if actual != test.count {
			t.Errorf("Get %d; expected %d", actual, test.count)
		}
	}
}

func TestLessThan4(t *testing.T) {

	var array = []string{"apple", "cat", "dog", "banana"}
	var res = []string{"cat", "dog"}

	actual := LessThan4(array)
	isTrue := reflect.DeepEqual(actual, res)
	if !isTrue {
		t.Errorf("Get %s; expected %s", actual, res)
	}
}
