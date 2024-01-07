package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		str := ""
		if i%2 == 0 {
			str = "偶数"
		} else {
			str = "奇数"
		}
		fmt.Printf("%d-%s\n", i, str)
	}
}
