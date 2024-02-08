package main

import (
	"fmt"
	"time"
)

func tick(duration int, ticker time.Duration, done chan<- bool) {
	for remaining := duration; remaining >= 0; remaining-- {
		<-time.Tick(ticker)
		fmt.Printf("\r残り%d秒", remaining)
	}
	done <- true
}

func main() {
	durationSecond := 60
	tickerSecond := time.Second
	tickDone := make(chan bool)
	go tick(durationSecond, tickerSecond, tickDone)

	select {
	case <-tickDone:
		fmt.Println("\n終了")
	}
}
