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

func TestRevrot_5(t *testing.T) {
	got := task.Revrot("66443875", 4)
	want := "44668753"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(\"66443875\", 4) = %v; want %v", got, want)
	}
}

func TestRevrot_6(t *testing.T) {
	got := task.Revrot("123456779", 8)
	want := "23456771"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(\"123456779\", 8) = %v; want %v", got, want)
	}
}

func TestRevrot_7(t *testing.T) {
	got := task.Revrot("123456987654", 6)
	want := "234561876549"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Revrot(\"123456987654\", 6) = %v; want %v", got, want)
	}
}
