package tests

import (
	"helloapp/app/task"
	"testing"
)

func TestEncryptThis_1(t *testing.T) {
	got := task.EncryptThis("")
	want := ""

	if got != want {
		t.Errorf("EncryptThis(\"\") = %v; want %v", got, want)
	}
}

func TestEncryptThis_2(t *testing.T) {
	got := task.EncryptThis("Hello")
	want := "72olle"

	if got != want {
		t.Errorf("CEncryptThis(\"Hello\") = %v; want %v", got, want)
	}
}

func TestEncryptThis_3(t *testing.T) {
	got := task.EncryptThis("An")
	want := "65n"

	if got != want {
		t.Errorf("EncryptThis(\"An\") = %v; want %v", got, want)
	}
}

func TestEncryptThis_4(t *testing.T) {
	got := task.EncryptThis("A wise old owl lived in an oak")
	want := "65 119esi 111dl 111lw 108dvei 105n 97n 111ka"

	if got != want {
		t.Errorf("EncryptThis(\"A wise old owl lived in an oak\") = %v; want %v", got, want)
	}
}
