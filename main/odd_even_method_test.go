package main

import (
	"testing"
)

func Test_isEven(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{0}, true},
		{"1", args{1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEven(tt.args.n); got != tt.want {
				t.Errorf("isEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
