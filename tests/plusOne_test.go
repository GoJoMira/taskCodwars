package tests

import (
	"helloapp/app/task"
	"reflect"
	"testing"
)

func TestPlusOne_1(t *testing.T) {
	got := task.PlusOne([]int{1, 2, 3})
	want := []int{1, 2, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveDuplicates() = %v; want %v", got, want)
	}
}

func TestPlusOne_2(t *testing.T) {
	got := task.PlusOne([]int{4, 3, 2, 1})
	want := []int{4, 3, 2, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveDuplicates() = %v; want %v", got, want)
	}
}

func TestPlusOne_3(t *testing.T) {
	got := task.PlusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6})
	want := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveDuplicates() = %v; want %v", got, want)
	}
}

func TestPlusOne_4(t *testing.T) {
	got := task.PlusOne([]int{9, 9})
	want := []int{1, 0, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveDuplicates() = %v; want %v", got, want)
	}
}

func BenchmarkPlusOne_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = task.PlusOne([]int{1, 2, 3})
	}
}
