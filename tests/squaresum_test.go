package tests

import (
	"helloapp/app/task"
	"testing"
)

func TestSquaruSum_1(t *testing.T) {
	got := task.SquareSum([]int{1, 2})
	want := 5

	if got != want {
		t.Errorf("SquareSum([]int{1, 2}) = %d; want %d", got, want)
	}
}

func TestSquaruSum_2(t *testing.T) {
	got := task.SquareSum([]int{0, 3, 4, 5})
	want := 50

	if got != want {
		t.Errorf("SquareSum([]int{0, 3, 4, 5}) = %d; want %d", got, want)
	}
}

func TestSquaruSum_3(t *testing.T) {
	got := task.SquareSum([]int{})
	want := 0

	if got != want {
		t.Errorf("SquareSum([]int{}) = %d; want %d", got, want)
	}
}
