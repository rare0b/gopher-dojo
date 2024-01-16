package main

func swap2(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}
