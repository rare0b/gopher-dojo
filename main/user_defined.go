package main

import (
	"fmt"
	"math/rand"
)

type Scores []int

func (scores *Scores) Sum() int {
	sum := 0
	for _, v := range *scores {
		sum += v
	}
	return sum
}

func main() {
	var scores Scores
	for i := 0; i < 5; i++ {
		scores = append(scores, rand.Intn(101))
	}
	fmt.Printf("%+v\n", scores)
	println(scores.Sum())
}
