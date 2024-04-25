package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	ans := Add(2, 2)
	want := 4
	if ans != 4 {
		t.Errorf("got %d, want %d", ans, want)
	}
}

func TestAddMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{15, 12, 27},
		{11, 0, 11},
		{2, -21, -19},
		{10, -11, -1},
		{-11, 20, 9},
	}

	for _, tt := range tests {
		ans := Add(tt.a, tt.b)
		if ans != tt.want {
			t.Errorf("got %d, want %d", ans, tt.want)
		}
	}

}

func TestSum(t *testing.T) {
	var tests = []struct {
		nums []int
		want  int
	}{
		{nums: []int{1, 2, 3}, want: 6},
		{nums: []int{2, 3, 4}, want: 9},
	}

	for _, tt := range tests {
		ans := Sum(tt.nums...)
		if ans != tt.want {
			t.Errorf("got %d, want %d", ans, tt.want)
		}
	}
}

func TestAbs(t *testing.T) {
	var tests = []struct {
		a int
		want int
	}{
		{a: -6, want: 6},
		{a: 9, want: 9},
		{a: 0, want: 0},
	}

	for _, tt := range tests {
		ans := Abs(tt.a)
		if ans != tt.want {
			t.Errorf("got %d, want %d", ans, tt.want)
		}
	}
}