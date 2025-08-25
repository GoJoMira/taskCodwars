package tests

import (
	"helloapp/app/task"
	"reflect"
	"testing"
)

func TestIntersect_1(t *testing.T) {
	got := task.Intersect([]int{1, 2, 2, 1}, []int{2, 2})
	want := []int{2, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Intersect = %v; want %v", got, want)
	}
}

func TestIntersect_2(t *testing.T) {
	got := task.Intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	want := []int{9, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Intersect = %v; want %v", got, want)
	}
}

func TestIntersect_3(t *testing.T) {
	got := task.Intersect([]int{4, 5, 6, 9}, []int{4, 4, 5, 6, 6, 9, 8, 9})
	want := []int{4, 5, 6, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Intersect = %v; want %v", got, want)
	}
}

func TestIntersect_4(t *testing.T) {
	got := task.Intersect([]int{1, 2}, []int{1, 2, 2})
	want := []int{1, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Intersect = %v; want %v", got, want)
	}
}

func TestIntersect_5(t *testing.T) {
	got := task.Intersect([]int{3, 1, 2}, []int{1, 1})
	want := []int{1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Intersect = %v; want %v", got, want)
	}
}

func BenchmarkIntersect_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = task.Intersect([]int{1, 2, 2, 1}, []int{2, 2})
	}
}
