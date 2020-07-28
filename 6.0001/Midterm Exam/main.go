package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

// Exercise 3:
// Write a recursive Python function, given a non-negative integer N,
// to calculate the no. of occurrences of digit 7 in N.
// Hint: Mod (%) by 10 gives you the rightmost digit (126 % 10 is 6),
// while doing integer division by 10 removes the rightmost digit (126 / 10 is 12).
// This function has to be recursive; you may not use loops!
// This function takes in one integer and returns one integer.\
func Count7(num int) (count int) {
	if num <= 6 {
		return 0
	}

	for num%10 == 7 {
		count++
		num = num / 10
	}

	return count + Count7(num/10)
}

// Exercise 4:
// Write a Python function that returns the sublist of strings in aList
// that contain fewer than 4 characters.
// For example, if aList = ["apple", "cat", "dog", "banana"],
// your function should return: ["cat", "dog"]
// This function takes in a list of strings and returns a list of strings.
// Your function should not modify aList.
func LessThan4(arr []string) (res []string) {
	for _, item := range arr {
		if len(item) < 4 {
			res = append(res, item)
		}
	}
	return
}

// Exercise 5:
// Write a Python function that returns a list of keys in aDict with the value
// target. The list of keys you return should be sorted in increasing order.
// The keys and values in aDict are both integers. (If aDict does not contain
// the value target, you should return an empty list.)
// This function takes in a dictionary and an integer and returns a list.
func keysWithValue(dict map[int]int, target int) (res []int) {
	for k, v := range dict {
		if target == v {
			res = append(res, k)
		}
	}

	sort.Ints(res)
	return
}

// Exercise 6:

// Write a function to flatten a list. The list contains other lists, strings,
// or ints. For example, [[1,'a',['cat'],2],[[[3]],'dog'],4,5] is flattened
// into [1,'a','cat',2,3,'dog',4,5] (order matters).
// 这里要保证嵌套的切片最终是同一类型,

func Flatten(arr interface{}) ([]int, error) {
	return doFlatten([]int{}, arr)
}

func doFlatten(acc []int, arr interface{}) ([]int, error) {
	var err error

	switch v := arr.(type) {
	case []int:
		acc = append(acc, v...)
	case int:
		acc = append(acc, v)
	case []interface{}:
		for i := range v {
			acc, err = doFlatten(acc, v[i])
			if err != nil {
				return nil, errors.New("not int or []int given")
			}
		}
	default:
		return nil, errors.New("not int given")
	}

	return acc, nil
}

// Exercise 7:

// L, a list of numbers (n0, n1, n2, ... nk)
// Returns a function, which when applied to a value x, returns the value
// n0 * x^k + n1 * x^(k-1) + ... nk * x^0
// If L = [1, 2, 3, 4]
// Invoking function with list L and value of x = 10 -
// general_poly(L)(10)
// result should be 1234
// i.e. 1 * 10^3 + 2 * 10^2 + 3 * 10^1 + 4 * 10^0

func generalPoly(list []float64) func(float64) float64 {

	if len(list) == 0 {
		fmt.Println("[]float64 given has no element")
		return nil
	}

	return func(n float64) (res float64) {
		var length = len(list)
		for _, i := range list {
			res += i * math.Pow(n, float64(length))
			length--
		}
		return
	}

}
func main() {
	res, err := Flatten([]interface{}{[]interface{}{1, 2, []int{3}}, 4, []int{4, 6, 8}})

	fmt.Println(res, err)

	res2 := keysWithValue(map[int]int{5: 5, 1: 5, 3: 4, 0: 5}, 5)

	fmt.Println(res2)

	res3 := generalPoly([]float64{1, 2, 3, 4})(10)

	fmt.Println(res3)
}
