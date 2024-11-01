package tests

import (
	"helloapp/app/task"
	"testing"
)

func TestAdd(t *testing.T) {
	got := task.Add(1, 2)
	want := 3

	if got != want {
		t.Errorf("Add(1, 2) = %d; want %d", got, want)
	}
}
