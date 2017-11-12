package mono

import (
	"testing"
	"bytes"
)

func TestNewMono(t *testing.T) {
	m := NewMono()
	if m.key_ != "abcdefghijklmnopqrstuvwxyz" {
		t.Errorf("m.key_ should be abcdefghijklmnopqrstuvwxyz, is %s", m.key_)
	}
}

func TestSet_key(t *testing.T) {
	m := NewMono()
	m.Set_key("zyxwvutsrqponmlkjihgfedcba")
	if m.key_ != "zyxwvutsrqponmlkjihgfedcba" {
		t.Errorf("m.key_ should be zyxwvutsrqponmlkjihgfedcba, is %s", m.key_)
	}
}

func TestDecode(t *testing.T) {
	m := NewMono()
	m.Set_key("zyxwvutsrqponmlkjihgfedcba")
	if m.Decode("ABC") != "zyx" {
		t.Errorf("ABC should decode to zyx")
	}
}

func TestDisplay(t *testing.T) {
	m := NewMono()
	m.Set_key("zyxwvutsrqponmlkjihgfedcba")
	var b bytes.Buffer
	m.Display(&b)
	expected := string(`ZYXWVUTSRQPONMLKJIHGFEDCBA
abcdefghijklmnopqrstuvwxyz
`)
	if b.String() != expected {
		t.Errorf("Display should have printed \n[%v]\n, but actually printed \n[%v]\n", expected, b.String())
    }
}
