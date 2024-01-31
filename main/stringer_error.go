package main

import (
	"errors"
	"fmt"
)

func ToStringer(v interface{}) (Stringer, error) {
	stringerV, ok := v.(Stringer)
	if !ok {
		return nil, errors.New("Stringerに変換できませんでした")
	}
	return stringerV, nil
}

type NotStringer string

func main() {
	a := A("a")
	n := NotStringer("Not Stringer")
	stringerA, err := ToStringer(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stringerA)
	stringerN, err := ToStringer(n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stringerN)
}
