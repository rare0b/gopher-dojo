package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	results := []string{
		"大吉",
		"吉",
		"中吉",
		"小吉",
		"末吉",
		"凶",
		"大凶",
	}

	if r.FormValue("p") == "Gopher	" {
		fmt.Fprint(w, "Gopherさんの運勢は「大吉」です！")
	} else {
		fmt.Fprint(w, results[rand.Intn(len(results))])
	}
}

func main() {
	http.HandleFunc("/", omikujiHandler)
	http.ListenAndServe(":8080", nil)
}
