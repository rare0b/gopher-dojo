package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		str := ""
		switch i % 2 {
		case 0:
			str = "偶数"
		case 1:
			str = "奇数"
		}
		fmt.Printf("%d-%s\n", i, str)
	}
}
