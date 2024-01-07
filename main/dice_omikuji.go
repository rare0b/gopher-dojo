package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(6) + 1
	fmt.Printf("dice: %d\n", n)

	if n == 6 {
		fmt.Println("大吉です！")
	} else if n == 5 || n == 4 {
		fmt.Println("中吉です。")
	} else if n == 3 || n == 2 {
		fmt.Println("吉です。")
	} else {
		fmt.Println("凶です…")
	}
}
