package main

func main() {
	n := []int{19, 86, 1, 12}
	sum := 0
	for _, v := range n {
		sum += v
	}
	println(sum)
}
