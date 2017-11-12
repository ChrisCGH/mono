package mono

import (
	"testing"
	"fmt"
)

func TestNewAlphabet(t *testing.T) {
	a := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
    if a.alphabet_ != "abcdefghijklmnopqrstuvwxyz" {
		t.Error("a.alphabet_ should be %s, but is %s", "abcdefghijklmnopqrstuvwxyz", a.alphabet_)
	}
	if a.alphabet_size_ != 26 {
		t.Error("a.alphabet_size_ should be 26 but is %d", a.alphabet_size_)
	}
    if a.Alphabet() != "abcdefghijklmnopqrstuvwxyz" {
		t.Error("a.Alphabet() should be %s, but is %s", "abcdefghijklmnopqrstuvwxyz", a.Alphabet())
	}
	if a.c1_ != 0 {
		t.Error("a.c1_ should be 0 but is %d", a.c1_)
	}
	if a.c2_ != 0 {
		t.Error("a.c2_ should be 0 but is %d", a.c2_)
	}
}

func TestRandomise(t *testing.T) {
	aa := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	a := &aa
	f := NewFixed_Key()
	a.Randomise(f)
	if a.alphabet_size_ != 26 {
		t.Error("a.alphabet_size_ should be 26 but is %d", a.alphabet_size_)
	}
	aa = NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	a = &aa
	f.Set(byte('e'), byte('J'))
	f.Set(byte('x'), byte('Z'))
	a.Randomise(f)
	m := NewMono()
	m.Set_key(a.Alphabet())
	if m.Decode("J") != "e" {
		t.Error("J should decode to e")
	}
	if m.Decode("Z") != "x" {
		t.Error("Z should decode to x")
	}

	a.Randomise1(f, 1)
	m.Set_key(a.Alphabet())
	if m.Decode("J") != "e" {
		t.Error("J should decode to e")
	}
	if m.Decode("Z") != "x" {
		t.Error("Z should decode to x")
	}

	a.Randomise1(f, 24)
	m.Set_key(a.Alphabet())
	if m.Decode("J") != "e" {
		t.Error("J should decode to e")
	}
	if m.Decode("Z") != "x" {
		t.Error("Z should decode to x")
	}
}

func TestSwaps(t *testing.T) {
	aa := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	a := &aa
	a.Start_swaps()
	if a.c1_ != 0 {
		t.Error("a.c1_ should be 0, is %d", a.c1_)
	}
	if a.c2_ != 0 {
		t.Error("a.c2_ should be 0, is %d", a.c2_)
	}
	m := NewMono()
	f := NewFixed_Key()
	f.Set(byte('e'), byte('J'))
	f.Set(byte('x'), byte('Z'))
	a.Randomise(f)
	m.Set_key(a.Alphabet())
	if m.Decode("J") != "e" {
		t.Error("J should decode to e")
	}
	if m.Decode("Z") != "x" {
		t.Error("Z should decode to x")
	}
	a.Start_swaps()
	for ! a.End_swaps() {
		a.Next_swap(f)
		m.Set_key(a.Alphabet())
		if m.Decode("J") != "e" {
			t.Error("J should decode to e")
			fmt.Println(a.Alphabet())
		}
		if m.Decode("Z") != "x" {
			t.Error("Z should decode to x")
		}
	}
}

