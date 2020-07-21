// Write a program that does the following in order:
// 1. Asks the user to enter a number “x”
// 2. Asks the user to enter a number “y”
// 3. Prints out number “x”, raised to the power “y”.
// 4. Prints out the log (base 2) of “x”.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var inputReader *bufio.Reader
var input string
var x, y float64
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter number x: ")
	input, err = inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("The x was: %s", err)
	}
	input = strings.TrimSuffix(input, "\n")

	if x, err = strconv.ParseFloat(input, 64); err != nil {
		fmt.Println(err, "x is not an integer.")
	}

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter number y: ")
	input, err = inputReader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	if err != nil {
		fmt.Printf("The y was: %s", err)
	}
	if y, err = strconv.ParseFloat(input, 64); err != nil {
		fmt.Println(input, "y is not an integer.")
	}

	fmt.Printf("%f**%f = %f\n", x, y, math.Pow(x, y))
	fmt.Printf("log(%f) = %f\n", x, math.Log2(x))
}
