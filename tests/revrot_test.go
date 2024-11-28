package tests

import (
	"helloapp/app/task"
	"reflect"
	"testing"
)

func TestRevrot_1(t *testing.T) {
	got := task.Revrot("123", 0)
	want := ""

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(\"123\", 0) = %v; want %v", got, want)
	}
}

func TestRevrot_2(t *testing.T) {
	got := task.Revrot("", 8)
	want := ""

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(\"\", 8) = %v; want %v", got, want)
	}
}

func TestRevrot_3(t *testing.T) {
	got := task.Revrot("", 0)
	want := ""

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(\"\", 0) = %v; want %v", got, want)
	}
}

func TestRevrot_4(t *testing.T) {
	got := task.Revrot("1234", 5)
	want := ""

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(\"1234\", 5) = %v; want %v", got, want)
	}
}
