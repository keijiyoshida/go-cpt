package main

import "testing"

func Test_max(t *testing.T) {
	testCases := []struct{
		x int
		y int
		want int
	}{
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, 1},
	}
	for _, tc := range testCases {
		if got := max(tc.x, tc.y); got != tc.want {
			t.Errorf("max(%d, %d) = %d; want %d",
				tc.x, tc.y, got, tc.want)
		}
	}
}

func Test_min(t *testing.T) {
	testCases := []struct{
		x int
		y int
		want int
	}{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	for _, tc := range testCases {
		if got := min(tc.x, tc.y); got != tc.want {
			t.Errorf("min(%d, %d) = %d; want %d",
				tc.x, tc.y, got, tc.want)
		}
	}
}

func Test_abs(t *testing.T) {
	testCases := []struct{
		x int
		want int
	}{
		{-1, 1},
		{0, 0},
		{1, 1},
	}
	for _, tc := range testCases {
		if got := abs(tc.x); got != tc.want {
			t.Errorf("abs(%d) = %d; want %d",
				tc.x, got, tc.want)
		}
	}
}

func Test_joinInt64s(t *testing.T) {
	testCases := []struct{
		a []int64
		sep string
		want string
	}{
		{[]int64{}, " ", ""},
		{[]int64{1}, " ", "1"},
		{[]int64{1, 2}, " ", "1 2"},
		{[]int64{1, 2, 3}, " ", "1 2 3"},
	}
	for _, tc := range testCases {
		if got := joinInt64s(tc.a, tc.sep); got != tc.want {
			t.Errorf("joinInt64s(%d, %s) = %s; want %s",
				tc.a, tc.sep, got, tc.want)
		}
	}
}

func Test_joinInts(t *testing.T) {
	testCases := []struct{
		a []int
		sep string
		want string
	}{
		{[]int{}, " ", ""},
		{[]int{1}, " ", "1"},
		{[]int{1, 2}, " ", "1 2"},
		{[]int{1, 2, 3}, " ", "1 2 3"},
	}
	for _, tc := range testCases {
		if got := joinInts(tc.a, tc.sep); got != tc.want {
			t.Errorf("joinInts(%d, %s) = %s; want %s",
				tc.a, tc.sep, got, tc.want)
		}
	}
}

func Test_divUp(t *testing.T) {
	testCases := []struct{
		x int
		y int
		want int
	}{
		{0, 1, 0},
		{1, 1, 1},
		{2, 1, 2},
		{3, 1, 3},
		{0, 2, 0},
		{1, 2, 1},
		{2, 2, 1},
		{3, 2, 2},
		{4, 2, 2},
		{0, 3, 0},
		{1, 3, 1},
		{2, 3, 1},
		{3, 3, 1},
		{4, 3, 2},
		{5, 3, 2},
		{6, 3, 2},
	}
	for _, tc := range testCases {
		if got := divUp(tc.x, tc.y); got != tc.want {
			t.Errorf("divUp(%d, %d) = %d; want %d",
				tc.x, tc.y, got, tc.want)
		}
	}
}

func Test_gcd(t *testing.T) {
	testCases := []struct{
		x int
		y int
		want int
	}{
		{1, 1, 1},
		{1, 2, 1},
		{1, 3, 1},
		{2, 1, 1},
		{2, 2, 2},
		{2, 3, 1},
		{2, 4, 2},
		{2, 5, 1},
		{2, 6, 2},
		{3, 1, 1},
		{3, 2, 1},
		{3, 3, 3},
		{3, 4, 1},
		{3, 5, 1},
		{3, 6, 3},
		{375, 525, 75},
		{525, 375, 75},
	}
	for _, tc := range testCases {
		if got := gcd(tc.x, tc.y); got != tc.want {
			t.Errorf("gcd(%d, %d) = %d; want %d",
				tc.x, tc.y, got, tc.want)
		}
	}
}

func Test_lcm(t *testing.T) {
	testCases := []struct{
		x int
		y int
		want int
	}{
		{1, 1, 1},
		{1, 2, 2},
		{1, 3, 3},
		{2, 1, 2},
		{2, 2, 2},
		{2, 3, 6},
		{2, 4, 4},
		{3, 1, 3},
		{3, 2, 6},
		{3, 3, 3},
		{3, 4, 12},
		{3, 5, 15},
		{3, 6, 6},
		{121, 17, 2057},
		{17, 121, 2057},
	}
	for _, tc := range testCases {
		if got := lcm(tc.x, tc.y); got != tc.want {
			t.Errorf("lcm(%d, %d) = %d; want %d",
				tc.x, tc.y, got, tc.want)
		}
	}
}

func Test_pow(t *testing.T) {
	testCases := []struct{
		x int
		y int
		want int
	}{
		{1, 0, 1},
		{1, 1, 1},
		{1, 2, 1},
		{2, 0, 1},
		{2, 1, 2},
		{2, 2, 4},
		{2, 3, 8},
		{2, 10, 1024},
		{3, 0, 1},
		{3, 1, 3},
		{3, 2, 9},
		{3, 3, 27},
		{3, 10, 59049},
	}
	for _, tc := range testCases {
		if got := pow(tc.x, tc.y); got != tc.want {
			t.Errorf("pow(%d, %d) = %d,; want %d",
				tc.x, tc.y, got, tc.want)
		}
	}
}