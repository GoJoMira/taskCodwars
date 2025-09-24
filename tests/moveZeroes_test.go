package tests

import (
	"helloapp/app/task"
	"reflect"
	"testing"
)

func TestMoveZeroes_1(t *testing.T) {
	got := []int{0, 1, 0, 3, 12}
	task.MoveZeroes(got)
	want := []int{1, 3, 12, 0, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MoveZeroes = %v; want %v", got, want)
	}
}

func TestMoveZeroes_2(t *testing.T) {
	got := []int{0}
	task.MoveZeroes(got)
	want := []int{0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MoveZeroes = %v; want %v", got, want)
	}
}

func BenchmarkMoveZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got := []int{0, 1, 0, 3, 12}
		task.MoveZeroes(got)
	}
}
