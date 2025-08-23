package tests

import (
	"helloapp/app/task"
	"reflect"
	"testing"
)

func TestSingleNumber_1(t *testing.T) {
	got := task.SingleNumber([]int{2, 1, 4, 1, 2})
	want := 4

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(\"123\", 0) = %v; want %v", got, want)
	}
}

func TestSingleNumber_2(t *testing.T) {
	got := task.SingleNumber([]int{1})
	want := 1

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(\"123\", 0) = %v; want %v", got, want)
	}
}

func TestSingleNumber_3(t *testing.T) {
	got := task.SingleNumber([]int{2, 2, 1})
	want := 1

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(\"123\", 0) = %v; want %v", got, want)
	}
}

func TestSingleNumber_4(t *testing.T) {
	got := task.SingleNumber([]int{10, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8})
	want := 10

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(\"123\", 0) = %v; want %v", got, want)
	}
}

func BenchmarkSolution_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = task.SingleNumber([]int{10, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8})
	}
}
