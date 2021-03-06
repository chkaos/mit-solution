package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// 2.1 Finger exercise: Write a program that examines three variables— x , y , and z —
// and prints the largest odd number among them. If none of them are odd, it
// should print a message to that effect.

// 2.4 Finger exercise: Write a program that asks the user to input 10 integers, and
// then prints the largest odd number that was entered. If no odd number was
// entered, it should print a message to that effect. 变式不大, 可以套用2.1

func GetBiggestOdd(x, y, z int) (res int, err error) {
	var hasOdd = false
	temp := []int{x, y, z}
	for _, val := range temp {
		if val%2 == 1 {
			hasOdd = true
			if val > res {
				res = val
			}
		}
	}

	if !hasOdd {
		return 0, errors.New("none of them are odd")
	}

	return

}

// 3.1 Finger exercise: Write a program that asks the user to enter an integer and
// prints two integers, root and pwr , such that 0 < pwr < 6 and root**pwr is equal
// to the integer entered by the user. If no such pair of integers exists, it should
// print a message to that effect.
func GetPair(target int) (root, pwa int, err error) {
	// pwr的范围应该是2-6，因为总是会有一个整数到pwr为1的答案，并且练习声明将存在没有这样的对的整数。
	for pwa = 2; pwa < 6; pwa++ {
		a := math.Pow(float64(target), 1.0/float64(pwa))
		if math.Ceil(a) == a {
			return int(a), pwa, nil
		}
	}
	return 0, 0, errors.New("No pair founded")
}

// 3.2 Finger exercise: Let s be a string that contains a sequence of decimal numbers
// separated by commas, e.g., s = '1.23,2.4,3.123' . Write a program that prints
// the sum of the numbers in s .

func Count(numStr string) (total float64) {
	s := strings.Split(numStr, ",")
	for num := range s {
		total += float64(num)
	}
	return
}

// 3.3 Finger exercise: What would the code in Figure 3.4 do if the statement x = 25
// were replaced by x = -25?
// x = 25
// epsilon = 0.01
// numGuesses = 0
// low = 0.0
// high = max(1.0, x)
// ans = (high + low)/2.0
// while abs(ans**2 - x) >= epsilon:
// print 'low =', low, 'high =', high, 'ans =', ans
// numGuesses += 1
// if ans**2 < x:
// low = ans
// else:
// high = ans
// ans = (high + low)/2.0
// print 'numGuesses =', numGuesses
// print ans, 'is close to square root of', x

func FindSquareRoot(target float64) (ans float64) {
	if target <= 0 {
		fmt.Printf("%f has not square root\n", target)
		return
	}
	epsilon := 0.01
	var numGuesses int
	var low float64
	high := math.Max(0, target)
	ans = (high + low) / 2.0
	for math.Abs(math.Pow(ans, 2)-target) >= epsilon {
		fmt.Printf("low = %f, high = %f, ans = %f\n", low, high, ans)
		numGuesses++
		if math.Pow(ans, 2) < target {
			low = ans
		} else {
			high = ans
		}
		ans = (high + low) / 2.0
	}
	fmt.Printf("numGuesses = %d\n", numGuesses)
	fmt.Printf("%f is close to square root of %f\n", ans, target)
	return
}

// 3.4 Finger exercise: What is the decimal equivalent of the binary number 10011 ? 10进制的 19

// 3,5 牛顿法求平方根
// #Newton-Raphson for square root
// #Find x such that x**2 - 24 is within epsilon of 0
// epsilon = 0.01
// k = 24.0
// guess = k/2.0
// while abs(guess*guess - k) >= epsilon:
// guess = guess - (((guess**2) - k)/(2*guess))
// print 'Square root of', k, 'is about', guess

func NewtonSquareRoot(target float64) (guess float64) {
	if target <= 0 {
		fmt.Printf("%f has not square root\n", target)
		return
	}
	epsilon := 0.01
	var numGuesses int
	guess = target / 2
	for math.Abs(math.Pow(guess, 2)-target) >= epsilon {
		fmt.Printf("current guess is %f\n", guess)
		numGuesses++
		guess = guess - (((math.Pow(guess, 2)) - target) / (2 * guess))
	}
	fmt.Printf("numGuesses = %d\n", numGuesses)
	fmt.Printf("%f is close to square root of %f\n", guess, target)
	return
}

// 4.1.1 Finger exercise: Write a function isIn that accepts two strings as arguments
// and returns True if either string occurs anywhere in the other, and False
// otherwise. Hint: you might want to use the built-in str operation in .
func isIn(string1, string2 string) bool {
	tem := make(map[rune]bool)
	for _, letter := range string1 {
		tem[letter] = true
	}
	for _, letter := range string2 {
		if tem[letter] {
			return true
		}
	}
	return false
}

func factorial(n int) int {
	if n == 1 {
		return n
	}

	return n * factorial(n-1)
}

type solver interface {
	play(int)
}

// towers is example of type satisfying solver interface
type towers struct {
	// an empty struct
}

// play is sole method required to implement solver type
func (t *towers) play(n int) {
	t.moveN(n, "A", "B", "C")
}

// recursive algorithm
func (t *towers) moveN(n int, from, to, via string) {
	if n > 0 {
		t.moveN(n-1, from, via, to)
		t.moveM(from, to)
		t.moveN(n-1, via, to, from)
	}
}

// 斐波那契
func fibL(n int) int {
	var fib = map[int]int{1: 1, 2: 1}

	if n == 1 || n == 2 {
		return 1
	}

	for i := 3; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}

func (t *towers) moveM(from, to string) {
	fmt.Println("Move disk from rod", from, "to rod", to)
}

func main() {
	if res, err := GetBiggestOdd(15, 3, 8); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("largest odd number: %d\n", res)
	}

	fmt.Println("-------------------------")

	if root, pwa, err := GetPair(16); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("found pair, root :%d, pwa:%d\n", root, pwa)
	}
	fmt.Println("-------------------------")

	total := Count("1.23,2.4,3.123")
	fmt.Printf("Total is :%f\n", total)

	fmt.Println("-------------------------")

	FindSquareRoot(-25)

	fmt.Println("-------------------------")
	NewtonSquareRoot(123456)

	fmt.Println("-------------------------")
	in := isIn("abc", "apple")
	fmt.Println(in)

	fmt.Println("-------------------------")
	var t solver
	t = new(towers) // type towers must satisfy solver interface
	t.play(4)

	fmt.Println("-------------------------")
	f := fibL(2)
	fmt.Println(f)
}
