package main

import (
	"fmt"
	"github.com/rare0b/gopher-dojo/greeting"
	greeting1 "github.com/tenntenn/greeting"
	greeting2 "github.com/tenntenn/greeting/v2"
	"time"
)

func main() {
	fmt.Println(greeting.Do())
	fmt.Println(greeting1.Do())
	fmt.Println(greeting2.Do(time.Now()))
}
