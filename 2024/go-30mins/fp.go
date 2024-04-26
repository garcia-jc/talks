package main

import "fmt"

func main() {
	var (
		suma  = func(a, b int) int { return a + b }
		resta = func(a, b int) int { return a - b }
		mult  = func(a, b int) int { return a * b }
		div   = func(a, b int) int { return a / b }
	)
	a, b := 10, 5
	opName := "+"
	var op func(int, int) int
	switch opName {
	case "+":
		op = suma
	case "-":
		op = resta
	case "/":
		op = div
	case "*":
		op = mult
	}
	fmt.Println(a, opName, b, "=", op(a, b))
}
