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

func BenchmarkPlusOne_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = task.PlusOne([]int{1, 2, 3})
	}
}
