package tests

import (
	"helloapp/app/task"
	"reflect"
	"testing"
)

func TestPlayPass_1(t *testing.T) {
	got := task.PlayPass("BORN IN 2015!", 1)
	want := "!4897 Oj oSpC"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("PlayPass_1() = %v; want %v", got, want)
	}
}

func TestPlayPass_2(t *testing.T) {
	got := task.PlayPass("AAABBCCY", 1)
	want := "zDdCcBbB"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("PlayPass_2() = %v; want %v", got, want)
	}
}

func TestPlayPass_3(t *testing.T) {
	got := task.PlayPass("I LOVE YOU!!!", 1)
	want := "!!!vPz fWpM J"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("PlayPass_3() = %v; want %v", got, want)
	}
}

func TestPlayPass_4(t *testing.T) {
	got := task.PlayPass("I LOVE YOU ^^", 0)
	want := "^^ uOy eVoL I"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("PlayPass_4() = %v; want %v", got, want)
	}
}
