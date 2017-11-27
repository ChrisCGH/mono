package mono

import (
	"bytes"
	"testing"
)

func TestNewFixed_Key(t *testing.T) {
	f := NewFixed_Key()
	if len(f.fixed_) != 26 {
		t.Errorf("f.fixed_ should be 26 bytes\n")
	}
	for _, b := range f.fixed_ {
		if b != byte(' ') {
			t.Errorf("f.fixed_ should be all spaces\n")
		}
	}
	for _, i := range f.index_ {
		if i != -1 {
			t.Errorf("f.index_ should be all -1\n")
		}
	}
}

func TestLength(t *testing.T) {
	f := NewFixed_Key()
	if f.Length() != len(f.fixed_) {
		t.Errorf("f.Length() should be %d, but actually is %d\n", len(f.fixed_), f.Length())
	}
}

func TestNotFixedLength(t *testing.T) {
	f := NewFixed_Key()
	if f.NotFixedLength() != len(f.not_fixed_)-1 {
		t.Errorf("f.Length() should be %d, but actually is %d\n", len(f.not_fixed_)-1, f.NotFixedLength())
	}
}

func TestRest(t *testing.T) {
	ff := NewFixed_Key()
	f := &ff
	if f.Is_set(byte('e')) {
		t.Error("e should not be set in fixed key")
	}
	f.Set(byte('a'), byte('9'))
	if f.Is_set(byte('a')) {
		t.Error("a should not be set in fixed key")
	}
	f.Set(byte('A'), byte('A'))
	if f.Is_set(byte('a')) {
		t.Error("a should not be set in fixed key")
	}
	check := ff.fixed_
	f.clear(byte('A') - 1)
	if f.fixed_ != check {
		t.Error("f should be unchanged")
	}

	f.Set(byte('e'), byte('J'))
	if !f.Is_set(byte('e')) {
		t.Errorf("e should be set in fixed key\n")
	}
	if f.Get_ct(byte('9')) != byte(' ') {
		t.Errorf("9 should return ' '")
	}
	if f.Get_ct(byte('e')) != byte('J') {
		t.Errorf("e should be set to J\n")
	}
	if f.Get_pt(byte('J')) != byte('e') {
		t.Errorf("J should be set to e\n")
	}
	f.Set(byte('x'), byte('Z'))
	if !f.Is_set(byte('x')) {
		t.Errorf("x should be set in fixed key\n")
	}
	if f.Get_ct(byte('x')) != byte('Z') {
		t.Errorf("x should be set to Z\n")
	}
	if f.Get_pt(byte('Z')) != byte('x') {
		t.Errorf("Z should be set to x\n")
	}
	if f.Is_set(byte('f')) {
		t.Error("f should not be set in fixed key")
	}
	if f.Get_ct(byte('f')) != byte(' ') {
		t.Error("f should be set to ' ' in fixed key")
	}
	if f.Get_pt(byte('F')) != byte(' ') {
		t.Error("F should be set to ' ' in fixed key")
	}
	if f.Get_pt(byte('0')) != byte(' ') {
		t.Error("0 should be set to ' ' in fixed key")
	}
	if f.Number_fixed() != 2 {
		t.Errorf("f.Number_fixed() should return 2, but actually returned %d\n", f.Number_fixed())
	}
	var b bytes.Buffer
	f.Display(&b)
	expected := `number fixed = 2
         e               x
Not fixed : [abcdfghijklmnopqrstuvwyz]
-1 -1 -1 -1 9 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 25 -1 -1 
`
	if b.String() != expected {
		t.Errorf("f.Display() should output \n%s\n, but actually output \n%s\n", expected, b.String())
	}
}
