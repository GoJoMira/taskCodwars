package tests

import (
	"helloapp/app/task"
	"reflect"
	"testing"
)

func TestRotate_1(t *testing.T) {
	var testNumber []int = []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	task.Rotate(testNumber, k)
	want := []int{5, 6, 7, 1, 2, 3, 4}

	if !reflect.DeepEqual(testNumber, want) {
		t.Errorf("Rotate() = %v; want %v", testNumber, want)
	}
}

func TestRotate_2(t *testing.T) {
	var testNumber []int = []int{-1, -100, 3, 99}
	k := 2
	task.Rotate(testNumber, k)
	want := []int{3, 99, -1, -100}

	if !reflect.DeepEqual(testNumber, want) {
		t.Errorf("Rotate() = %v; want %v", testNumber, want)
	}
}

func TestRotate_3(t *testing.T) {
	var testNumber []int = []int{1, 2, 3}
	k := 3
	task.Rotate(testNumber, k)
	want := []int{1, 2, 3}

	if !reflect.DeepEqual(testNumber, want) {
		t.Errorf("Rotate() = %v; want %v", testNumber, want)
	}
}

func TestRotate_4(t *testing.T) {
	var testNumber []int = []int{1, 2, 3}
	k := 4
	task.Rotate(testNumber, k)
	want := []int{3, 1, 2}

	if !reflect.DeepEqual(testNumber, want) {
		t.Errorf("Rotate() = %v; want %v", testNumber, want)
	}
}
