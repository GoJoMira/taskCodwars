package tests

import (
	"helloapp/app/task"
	"reflect"
	"testing"
)

func TestRotateImage_1(t *testing.T) {
	got := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	task.RotateImage(got)
	want := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Rotate() = %v; want %v", got, want)
	}
}

func TestRotateImage_2(t *testing.T) {
	got := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	task.RotateImage(got)
	want := [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Rotate() = %v; want %v", got, want)
	}
}
