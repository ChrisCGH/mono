package mono

import (
	"bytes"
	"testing"
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
	var b bytes.Buffer
	a.Display(&b)
	expected := string(`a b c d e f g h i j k l m n o p q r s t u v w x y z 
`)
	if b.String() != expected {
		t.Errorf("Display() should have printed\n[%v]\n, but actually printed \n[%v]\n", expected, b.String())
	}
}

func TestRandomise(t *testing.T) {
	aa := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	a := &aa
	ff := NewFixed_Key()
	f := &ff
	a.Randomise(ff)
	if a.alphabet_size_ != 26 {
		t.Error("a.alphabet_size_ should be 26 but is %d", a.alphabet_size_)
	}
	aa = NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	a = &aa
	f.Set(byte('e'), byte('J'))
	f.Set(byte('x'), byte('Z'))
	a.Randomise(ff)
	m := NewMono()
	m.Set_key(a.Alphabet())
	if m.Decode("J") != "e" {
		t.Error("J should decode to e")
	}
	if m.Decode("Z") != "x" {
		t.Error("Z should decode to x")
	}

	a.Randomise1(ff, -1)
	m.Set_key(a.Alphabet())
	if m.Decode("J") != "e" {
		t.Error("J should decode to e")
	}
	if m.Decode("Z") != "x" {
		t.Error("Z should decode to x")
	}

	a.Randomise1(ff, 1)
	m.Set_key(a.Alphabet())
	if m.Decode("J") != "e" {
		t.Error("J should decode to e")
	}
	if m.Decode("Z") != "x" {
		t.Error("Z should decode to x")
	}

	a.Randomise1(ff, 24)
	m.Set_key(a.Alphabet())
	if m.Decode("J") != "e" {
		t.Error("J should decode to e")
	}
	if m.Decode("Z") != "x" {
		t.Error("Z should decode to x")
	}

	a.Randomise1(ff, 25)
	m.Set_key(a.Alphabet())
	if m.Decode("J") != "e" {
		t.Error("J should decode to e")
	}
	if m.Decode("Z") != "x" {
		t.Error("Z should decode to x")
	}
	d := map[byte]byte{
		byte('a'): byte('Y'),
		byte('b'): byte('X'),
		byte('c'): byte('W'),
		byte('d'): byte('V'),
		byte('f'): byte('U'),
		byte('g'): byte('T'),
	}
	for p, c := range d {
		f.Set(p, c)
	}

	aa = NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	a = &aa
	a.Randomise1(ff, 20)
	m.Set_key(a.Alphabet())
	for p, c := range d {
		if m.Decode(string(c)) != string(p) {
			t.Errorf("%s should decode to %s, but actually decoded to %s", string(c), string(p), m.Decode(string(c)))
		}
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
	ff := NewFixed_Key()
	f := &ff
	d := map[byte]byte{
		byte('a'): byte('Y'),
		byte('b'): byte('X'),
		byte('c'): byte('W'),
		byte('d'): byte('V'),
		byte('f'): byte('U'),
		byte('g'): byte('T'),
	}
	for p, c := range d {
		f.Set(p, c)
	}

	aa = NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	a = &aa
	a.Randomise(ff)
	m.Set_key(a.Alphabet())
	for p, c := range d {
		if m.Decode(string(c)) != string(p) {
			t.Errorf("%s should decode to %s, but actually decoded to %s", string(c), string(p), m.Decode(string(c)))
		}
	}
	a.Start_swaps()
	for !a.End_swaps() {
		a.Next_swap(ff)
		m.Set_key(a.Alphabet())
		for p, c := range d {
			if m.Decode(string(c)) != string(p) {
				t.Errorf("%s should decode to %s, but actually decoded to %s", string(c), string(p), m.Decode(string(c)))
			}
		}
	}
}
