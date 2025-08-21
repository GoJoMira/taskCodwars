package tests

import (
	"helloapp/app/task"
	"testing"
)

func TestMaxProfit_1(t *testing.T) {
	got := task.MaxProfit([]int{1, 2, 3, 4, 5})
	want := 4

	if got != want {
		t.Errorf("MaxProfit = %d; want %d", got, want)
	}
}

func TestMaxProfit_2(t *testing.T) {
	got := task.MaxProfit([]int{7, 1, 5, 3, 6, 4})
	want := 7

	if got != want {
		t.Errorf("MaxProfit = %d; want %d", got, want)
	}
}

func TestMaxProfit_3(t *testing.T) {
	got := task.MaxProfit([]int{7, 6, 4, 3, 1})
	want := 0

	if got != want {
		t.Errorf("MaxProfit = %d; want %d", got, want)
	}
}

func TestMaxProfit_4(t *testing.T) {
	got := task.MaxProfit([]int{10, 15, 16, 5, 1, 5, 7})
	want := 12

	if got != want {
		t.Errorf("MaxProfit = %d; want %d", got, want)
	}
}

func TestMaxProfit_5(t *testing.T) {
	got := task.MaxProfit([]int{1, 2})
	want := 1

	if got != want {
		t.Errorf("MaxProfit = %d; want %d", got, want)
	}
}
