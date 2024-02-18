package main

import (
	"net/http"
)

func main() {
	url := "https://raw.githubusercontent.com/golang/go/master/README.md"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//out, err := os.Create("a.out")
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, err = io.Copy(out, resp.Body)
	//if err != nil {
	//	panic(err)
	//}

	//body, err := io.ReadAll(resp.Body)
	//print(string(body))
}
