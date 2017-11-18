package mono

import (
	"testing"
)

func TestNewCrib(t *testing.T) {
	c := NewCrib("", "", 0)
	expected := Crib{ct_: "", crib_: "", pt_: "", position_: 0, possible_: false, original_ct_: "", fixed_key_: nil}
	if c != expected {
		t.Errorf("c should be %v, but is actually %v\n", expected, c)
	}

	c = NewCrib("N EASSTH EIN GIVIC ANDI N EWORRID ENORSC", "a smooth sea", 0)
	expected_ct := "NEASSTHEINGIVICANDINEWORRIDENORSC"
	expected_crib := "asmoothsea"
	expected_position := 0
	expected_possible := true
	expected_original_ct := "N EASSTH EIN GIVIC ANDI N EWORRID ENORSC"
	if c.ct_ != expected_ct {
		t.Errorf("c.ct_ should be %s, but is actually %s\n", expected_ct, c.ct_)
	}
	if c.Get_ct() != expected_ct {
		t.Errorf("c.Get_ct() should be %s, but is actually %s\n", expected_ct, c.Get_ct())
	}

	if c.crib_ != expected_crib {
		t.Errorf("c.crib_ should be %s, but is actually %s\n", expected_crib, c.crib_)
	}
	if c.Get_crib_string() != expected_crib {
		t.Errorf("c.Get_crib_string() should be %s, but is actually %s\n", expected_crib, c.Get_crib_string())
	}

	if c.position_ != expected_position {
		t.Errorf("c.position_ should be %d, but is actually %d\n", expected_position, c.position_)
	}
	if c.Get_position() != expected_position {
		t.Errorf("c.Get_position() should be %d, but is actually %d\n", expected_position, c.Get_position())
	}

	if c.possible_ != expected_possible {
		t.Errorf("c.possible_ should be %d, but is actually %d\n", expected_possible, c.possible_)
	}
	if c.Is_possible() != expected_possible {
		t.Errorf("c.Is_possible() should be %d, but is actually %d\n", expected_possible, c.Is_possible())
	}

	if c.original_ct_ != expected_original_ct {
		t.Errorf("c.original_ct_ should be %s, but is actually %s\n", expected_original_ct, c.original_ct_)
	}
	if c.fixed_key_ == nil {
		t.Error("c.fixed_key should be non-nil")
	}
	if c.Get_fixed_key() == nil {
		t.Error("c.Get_fixed_key() should be non-nil")
	}
}

func TestMove(t *testing.T) {
	cc := NewCrib("N EASSTH EIN GIVIC ANDI N EWORRID ENORSC", "a smooth sea", 0)
	c := &cc
	c.Move_left()
	if c.position_ != 0 {
		t.Error("c.position_ should be 0 after Move_left()")
	}
	if !c.Is_possible() {
		t.Error("Should be possible after Move_left()")
	}
	c.Move_right()
	if c.position_ != 1 {
		t.Error("c.position_ should be 1 after Move_right()")
	}
	if c.Is_possible() {
		t.Error("Should not be possible after Move_right()")
	}
	c.Next_right()
	if c.position_ != 1 {
		t.Error("c.position_ should be 1 after Next_right()")
	}
	if c.Is_possible() {
		t.Error("Should not be possible after Next_right()")
	}
	cc = NewCrib("N EASSTH EIN GIVIC ANDI N EWORRID ENORSC", "a smooth", 0)
	c = &cc
	c.Next_right()
	if c.position_ != 20 {
		t.Error("c.position_ should be 20 after Next_right()")
	}
	if !c.Is_possible() {
		t.Error("Should be possible after Next_right()")
	}
	c.Next_right()
	if c.position_ != 20 {
		t.Error("c.position_ should be 20 after Next_right()")
	}
	if !c.Is_possible() {
		t.Error("Should be possible after Next_right()")
	}
	c.Next_left()
	if c.position_ != 0 {
		t.Error("c.position_ should be 0 after Next_left()")
	}
	if !c.Is_possible() {
		t.Error("Should be possible after Next_left()")
	}
}

func TestGet_pt(t *testing.T) {
	cc := NewCrib("N EASSTH EIN GIVIC ANDI N EWORRID ENORSC", "a smooth", 0)
	c := &cc
	expected_pt := "asmooths.a.....ma..as......sa..o."
	if c.Get_pt() != expected_pt {
		t.Errorf("c.Get_pt() should be %s, but actually is %s\n", expected_pt, c.Get_pt())
	}
	c.Next_right()
	expected_pt = ".a.....at..t.t...ht.asmootha.mo.."
	if c.Get_pt() != expected_pt {
		t.Errorf("c.Get_pt() should be %s, but actually is %s\n", expected_pt, c.Get_pt())
	}
}

func TestClear(t *testing.T) {
	cc := NewCrib("N EASSTH EIN GIVIC ANDI N EWORRID ENORSC", "a smooth", 0)
	c := &cc
	c.Next_right()
	c.Clear()
	if c.crib_ != "" {
		t.Errorf("c.crib_ should be blank but is %s\n", c.crib_)
	}
	if c.position_ != 0 {
		t.Errorf("c.position_ should be zero but is %d\n", c.position_)
	}
}

func TestPossible_positions(t *testing.T) {
	possible_positions := Possible_positions("N EASSTH EIN GIVIC ANDI N EWORRID ENORSC", "a smooth")
	if _, ok := possible_positions[0]; !ok {
		t.Error("0 should be a possible position")
	}
	if _, ok := possible_positions[20]; !ok {
		t.Error("20 should be a possible position")
	}
	if _, ok := possible_positions[15]; ok {
		t.Error("15 should not be a possible position")
	}

}
