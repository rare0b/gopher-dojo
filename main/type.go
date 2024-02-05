package main

import "fmt"

type Func func(n any) int

func main() {
	double := Func(func(n any) int {
		number, ok := n.(int)
		if !ok {
			return 0
		}
		return number * 2
	})

	fmt.Println(double(5))
}
