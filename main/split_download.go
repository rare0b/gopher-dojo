package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func getFileSize(url string) (int, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	contentLength := resp.Header.Get("Content-Length")
	if contentLength == "" {
		return 0, fmt.Errorf("Content-Length header is missing")
	}

	size, err := strconv.Atoi(contentLength)
	if err != nil {
		return 0, err
	}

	return size, nil
}

func generateRangeRequests(maxGoroutines int, fileSize int) []RangeRequest {
	partSize := fileSize / maxGoroutines
	firstRangeRequest := RangeRequest{
		start: 0,
		end:   partSize - 1, // ちょうどpartSizeバイト分取れるように-1する
	}
	dRangeRequests := []RangeRequest{firstRangeRequest}

	for i := 1; i < maxGoroutines; i++ {
		dRangeRequest := RangeRequest{
			start: dRangeRequests[i-1].end + 1,
			end:   dRangeRequests[i-1].end + partSize,
		}
		if i == maxGoroutines-1 {
			dRangeRequest.end = fileSize - 1
		}
		dRangeRequests = append(dRangeRequests, dRangeRequest)
	}

	return dRangeRequests
}

type RangeRequest struct {
	start int
	end   int
}

func DownloadRange(rangeRequest RangeRequest) {

}

func main() {
	url := "https://raw.githubusercontent.com/golang/go/master/README.md"

	resp, err := http.Head(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	contentLength := resp.Header.Get("Content-Length")
	fileSize, err := strconv.Atoi(contentLength)
	if err != nil {
		panic(err)
	}

	// GOMAXPROCSの値は自動値のまま
	maxGoroutines := runtime.GOMAXPROCS(0)
	rangeRequests := generateRangeRequests(maxGoroutines, fileSize)

	var wg sync.WaitGroup
	for i := 0; i < maxGoroutines; i++ {
		wg.Add(1)
		//TODO
	}

	out, err := os.Create("a.out")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
