package main

import "fmt"

type Stringer interface {
	String() string
}

type A string

func (a A) String() string {
	return fmt.Sprintf("%s", string(a))
}

type B string

func (b B) String() string {
	return fmt.Sprintf("%s", string(b))
}

type C string

func (c C) String() string {
	return fmt.Sprintf("%s", string(c))
}

func PrintType(s Stringer) {
	switch s.(type) {
	// 全部同じ処理だけど分岐の練習したいだけなので気にしない
	case A:
		fmt.Println(s.String(), "A")
	case B:
		fmt.Println(s.String(), "B")
	case C:
		fmt.Println(s.String(), "C")
	default:
		fmt.Println("N/A")
	}
}

func main() {
	var a A = "a"
	var b B = "b"
	var c C = "c"

	PrintType(a)
	PrintType(b)
	PrintType(c)
}
