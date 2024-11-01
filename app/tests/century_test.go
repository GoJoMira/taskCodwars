package tests

import (
	"helloapp/app/task"
	"testing"
)

func TestCentury_1(t *testing.T) {
	got := task.Century(1990)
	want := 20

	if got != want {
		t.Errorf("Century(1990) = %d; want %d", got, want)
	}
}

func TestCentury_2(t *testing.T) {
	got := task.Century(1705)
	want := 18

	if got != want {
		t.Errorf("Century(1990) = %d; want %d", got, want)
	}
}

func TestCentury_3(t *testing.T) {
	got := task.Century(2000)
	want := 20

	if got != want {
		t.Errorf("Century(1990) = %d; want %d", got, want)
	}
}

func TestCentury_4(t *testing.T) {
	got := task.Century(1601)
	want := 17

	if got != want {
		t.Errorf("Century(1990) = %d; want %d", got, want)
	}
}

func TestCentury_5(t *testing.T) {
	got := task.Century(89)
	want := 1

	if got != want {
		t.Errorf("Century(1990) = %d; want %d", got, want)
	}
}

func TestCentury_6(t *testing.T) {
	got := task.Century(121499)
	want := 1215

	if got != want {
		t.Errorf("Century(1990) = %d; want %d", got, want)
	}
}
