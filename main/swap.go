package main

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	n, m := swap(10, 20)
	println(n, m)
}
