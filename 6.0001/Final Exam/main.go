package main

import (
	"math"
)

//  Problem 3
//
//  Write a function is_triangular that meets the specification below. A triangular
//  number is a number obtained by the continued summation of integers starting
//  from 1. For example, 1, 1+2, 1+2+3, 1+2+3+4, etc., corresponding to
//  1, 3, 6, 10, etc., are triangular numbers.
func IsTriangular(num int) bool {
	if num == 1 {
		return true
	}
	var base int = 1
	for num > 0 {
		num -= base
		if num == 0 {
			return true
		}
		base++
	}

	return false
}

// Problem 4
//
// Write a Python function that creates and returns a list of prime numbers
// between 2 and N, inclusive, sorted in increasing order. A prime number is a
// number that is divisible only by 1 and itself. This function takes in an
// integer and returns a list of integers.
// Eratosthenes 素数筛 https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func GetPrimeList(n int) (list []int) {
	if n <= 1 {
		return
	}

	n++

	primeList := make([]bool, n)
	for i := range primeList {
		primeList[i] = true
	}

	sqrtroot := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrtroot; i++ {
		if primeList[i] == true {
			index := int(math.Pow(float64(i), 2))
			for index < n {
				primeList[index] = false
				index += i
			}
		}
	}

	for i := range primeList {
		if i > 1 && primeList[i] == true {
			list = append(list, i)
		}
	}

	return
}

// algorithm Sieve of Eratosthenes is
//     input: an integer n > 1.
//     output: all prime numbers from 2 through n.

//     let A be an array of Boolean values, indexed by integers 2 to n,
//     initially all set to true.

//     for i = 2, 3, 4, ..., not exceeding √n do
//         if A[i] is true
//             for j = i2, i2+i, i2+2i, i2+3i, ..., not exceeding n do
//                 A[j] := false

//     return all i such that A[i] is true.

// # =============================================================================
// # Problem 5
// #
// # Implement a function that meets the specifications below.
// #
// # def cipher(map_from, map_to, code):
// #     """ map_from, map_to: strings where each contain
// #                           N unique lowercase letters.
// #         code: string (assume it only contains letters also in map_from)
// #         Returns a tuple of (key_code, decoded).
// #         key_code is a dictionary with N keys mapping str to str where
// #         each key is a letter in map_from at index i and the corresponding
// #         value is the letter in map_to at index i.
// #         decoded is a string that contains the decoded version
// #         of code using the key_code mapping. """
// #     # Your code here
// #
// # For example,
// #
// # cipher("abcd", "dcba", "dab") returns (order of entries in dictionary may not
// #       be the same) ({'a':'d', 'b': 'c', 'd': 'a', 'c': 'b'}, 'adc')
// # =============================================================================

func Cipher(from, to, target string) (dict map[string]string, res string) {
	dict = make(map[string]string, 0)

	for i := 0; i < len(from); i++ {
		dict[string(from[i])] = string(to[i])
	}

	for i := 0; i < len(target); i++ {
		res += dict[string(target[i])]
	}

	return
}

func main() {}
