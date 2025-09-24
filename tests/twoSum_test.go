package tests

import (
	"helloapp/app/task"
	"reflect"
	"testing"
)

func TestTwoSum_1(t *testing.T) {
	got := task.TwoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Solve([]int) = %v; want %v", got, want)
	}
}

func TestTwoSum_2(t *testing.T) {
	got := task.TwoSum([]int{2, 7, 11, 15}, 18)
	want := []int{1, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Solve([]int) = %v; want %v", got, want)
	}
}

func TestTwoSum_3(t *testing.T) {
	got := task.TwoSum([]int{2, 7, 11, 15, 15}, 30)
	want := []int{3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Solve([]int) = %v; want %v", got, want)
	}
}
