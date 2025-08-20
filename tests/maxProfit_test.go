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
