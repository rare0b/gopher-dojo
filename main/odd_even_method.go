package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	for i := 1; i <= 100; i++ {
		str := ""
		if isEven(i) {
			str = "偶数"
		} else {
			str = "奇数"
		}
		fmt.Printf("%d-%s\n", i, str)
	}
}
