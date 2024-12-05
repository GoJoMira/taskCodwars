package tests

import (
	"helloapp/app/task"
	"reflect"
	"testing"
)

func TestRemoveDuplicates_1(t *testing.T) {
	var testNumber []int = []int{0, 1, 1, 1, 3, 3, 6, 10}
	got := task.RemoveDuplicates(testNumber)
	want := 5
	want1 := []int{0, 1, 3, 6, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveDuplicates() = %v; want %v", got, want)
	}

	if !reflect.DeepEqual(testNumber, want1) {
		t.Errorf("RemoveDuplicates() = %v; want %v", testNumber, want1)
	}
}
