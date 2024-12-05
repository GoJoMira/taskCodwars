package tests

import (
	"helloapp/app/task"
	"reflect"
	"testing"
)

func TestSolution_1(t *testing.T) {
	got := task.Solution(1)
	want := "I"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(1) = %v; want %v", got, want)
	}
}

func TestSolution_2(t *testing.T) {
	got := task.Solution(1000)
	want := "M"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(1000) = %v; want %v", got, want)
	}
}

func TestSolution_3(t *testing.T) {
	got := task.Solution(1666)
	want := "MDCLXVI"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(1666) = %v; want %v", got, want)
	}
}

func TestSolution_4(t *testing.T) {
	got := task.Solution(49)
	want := "XLIX"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(49) = %v; want %v", got, want)
	}
}

func TestSolution_5(t *testing.T) {
	got := task.Solution(77)
	want := "LXXVII"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(77) = %v; want %v", got, want)
	}
}

func TestSolution_6(t *testing.T) {
	got := task.Solution(2263)
	want := "MMCCLXIII"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(2263) = %v; want %v", got, want)
	}
}
