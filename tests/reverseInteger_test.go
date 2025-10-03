package tests

import (
	"helloapp/app/task"
	"testing"
)

func TestReverseInteger_1(t *testing.T) {
	got := task.ReverseInteger(-123)
	want := -321

	if got != want {
		t.Errorf("ReverseInteger = %d; want %d", got, want)
	}
}

func TestReverseInteger_2(t *testing.T) {
	got := task.ReverseInteger(1230)
	want := 321

	if got != want {
		t.Errorf("ReverseInteger = %d; want %d", got, want)
	}
}

func TestReverseInteger_3(t *testing.T) {
	got := task.ReverseInteger(12000)
	want := 21

	if got != want {
		t.Errorf("ReverseInteger = %d; want %d", got, want)
	}
}

func TestReverseInteger_4(t *testing.T) {
	got := task.ReverseInteger(12010)
	want := 1021

	if got != want {
		t.Errorf("ReverseInteger = %d; want %d", got, want)
	}
}

func TestReverseInteger_5(t *testing.T) {
	got := task.ReverseInteger(-1200100)
	want := -10021

	if got != want {
		t.Errorf("ReverseInteger = %d; want %d", got, want)
	}
}

func TestReverseInteger_6(t *testing.T) {
	got := task.ReverseInteger(1534236469)
	want := 0

	if got != want {
		t.Errorf("ReverseInteger = %d; want %d", got, want)
	}
}
