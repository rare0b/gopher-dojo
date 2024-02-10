package main

import (
	"bufio"
	"fmt"
	"github.com/tjarratt/babble"
	"os"
	"time"
)

func tick(duration int, ticker time.Duration, done chan<- bool) {
	for remaining := duration; remaining >= 0; remaining-- {
		<-time.Tick(ticker)
		print("\033[s")
		print("\033[1;1H")
		fmt.Printf("\r残り %d 秒", remaining)
		print("\033[u")
	}
	done <- true
	close(done)
}

func monitorStdin(inputChan chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputChan <- scanner.Text()
	}
}

func calculateScore(words []string, inputs []string) int {
	score := 0

	length := len(words)
	if len(inputs) < length {
		length = len(inputs)
	}

	for i := 0; i < length; i++ {
		if words[i] == inputs[i] {
			score++
		}
	}

	return score
}

func main() {
	durationSecond := 10
	tickerSecond := time.Second
	tickDone := make(chan bool)
	go tick(durationSecond, tickerSecond, tickDone)

	inputChan := make(chan string)
	go monitorStdin(inputChan)

	babbler := babble.NewBabbler()
	babbler.Count = 1
	println(babbler.Babble())

	var words []string
	var inputs []string
	print("\033[1;1H")
	fmt.Printf("\r残り %d 秒", durationSecond)
Loop:
	for {
		// 1行目
		print("\033[2;1H> ")
		word := babbler.Babble()
		words = append(words, word)
		print(word)

		// 2行目
		print("\033[3;1H")
		print("\033[K")
		print("> ")

		select {
		case <-tickDone:
			println("\n終了")
			close(inputChan)
			break Loop
		case input := <-inputChan:
			inputs = append(inputs, input)
		}
	}

	score := calculateScore(words, inputs)
	println(score, "点")
	fmt.Printf("%+v\n", words)
	fmt.Printf("%+v\n", inputs)
}
