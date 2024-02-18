package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	start := time.Now()
	cmd := exec.Command("go", "run", "C:\\Users\\rare\\git\\gopher-dojo\\main\\download.go")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running download:", err)
		return
	}
	duration := time.Since(start)
	fmt.Println("通常のダウンロード時間:", duration)

	start = time.Now()
	cmd = exec.Command("go", "run", "C:\\Users\\rare\\git\\gopher-dojo\\main\\split_download.go")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running split_download:", err)
		return
	}
	duration = time.Since(start)
	fmt.Println("分割ダウンロード時間:", duration)
}
