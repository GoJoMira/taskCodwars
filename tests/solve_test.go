package tests

import (
	"helloapp/app/task"
	"reflect"
	"testing"
)

func TestSolve_1(t *testing.T) {
	got := task.Solve([]int{2, 3, 5, 3, 7, 9, 5, 3, 7})
	want := []int{3, 3, 3, 5, 5, 7, 7, 2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Solve([]int) = %v; want %v", got, want)
	}
}

func TestSolve_2(t *testing.T) {
	got := task.Solve([]int{1, 2, 3, 0, 5, 0, 1, 6, 8, 8, 6, 9, 1})
	want := []int{1, 1, 1, 0, 0, 6, 6, 8, 8, 2, 3, 5, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Solve([]int) = %v; want %v", got, want)
	}
}
