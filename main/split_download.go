package main

import (
	"fmt"
	"io"
	"net/http"
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

func generateRangeRequests(maxGoroutines int, fileSize int) []rangeRequest {
	partSize := fileSize / maxGoroutines
	firstRangeRequest := rangeRequest{
		start: 0,
		end:   partSize - 1, // ちょうどpartSizeバイト分取れるように-1する
	}
	dRangeRequests := []rangeRequest{firstRangeRequest}

	for i := 1; i < maxGoroutines; i++ {
		dRangeRequest := rangeRequest{
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

type rangeRequest struct {
	start int
	end   int
}

func downloadRange(rangeRequest rangeRequest, url string, ch chan<- downloadResult, wg *sync.WaitGroup) {
	defer wg.Done()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ch <- downloadResult{
			data: nil,
			err:  err,
		}
		return
	}
	rangeHeader := fmt.Sprintf("bytes=%d-%d", rangeRequest.start, rangeRequest.end)
	req.Header.Set("Range", rangeHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ch <- downloadResult{
			data: nil,
			err:  err,
		}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- downloadResult{
			data: nil,
			err:  err,
		}
		return
	}

	ch <- downloadResult{
		data: body,
		err:  nil,
	}
	return
}

type downloadResult struct {
	data []byte
	err  error
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
	fmt.Printf("%+v\n", rangeRequests)
	println(rangeRequests[0].start, rangeRequests[0].end)

	resultChan := make(chan downloadResult, maxGoroutines)
	wg := &sync.WaitGroup{}
	for i := 0; i < maxGoroutines; i++ {
		wg.Add(1)
		go downloadRange(rangeRequests[i], url, resultChan, wg)
	}

	wg.Wait()
	close(resultChan)

	for result := range resultChan {
		if result.err != nil {
			println(result.err)
		} else {
			println(string(result.data))
		}
	}
}
