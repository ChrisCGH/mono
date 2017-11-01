package mono

import (
	"unicode"
)

type Crib struct {
	ct_          string
	crib_        string
	pt_          string
	position_    int
	possible_    bool
	original_ct_ string
	fixed_key_   *Fixed_Key
}

func NewCrib(ct string, crib string, position int) Crib {
	new_crib := Crib{}
	if ct != "" {
		new_crib.ct_ = copy_characters_upper(ct)
		new_crib.original_ct_ = ct
		if crib != "" {
			new_crib.crib_ = copy_characters_lower(crib)
		}
		new_crib.place_at(position)
	}
	return new_crib
}

func copy_characters_upper(s string) string {
	t := ""
	for _, r := range s {
		if unicode.IsLetter(r) {
			t += string(unicode.ToUpper(r))
		}
	}
	return t
}

func copy_characters_lower(s string) string {
	t := ""
	for _, r := range s {
		if unicode.IsLetter(r) {
			t += string(unicode.ToLower(r))
		}
	}
	return t
}

func (c Crib) Get_ct() string {
	return c.ct_
}

func (c Crib) Is_possible() bool {
	return c.possible_
}

func (c *Crib) Move_left() {
	if c.position_ <= 0 {
		return
	}
	c.place_at(c.position_ - 1)
}

func (c *Crib) Move_right() {
	if c.crib_ == "" {
		return
	}
	if len(c.crib_)+c.position_+1 > len(c.ct_) {
		return
	}
	c.place_at(c.position_ + 1)
}

func (c *Crib) Next_left() {
	pos := c.position_
	c.Move_left()
	for !c.possible_ && c.position_ > 0 {
		c.Move_left()
	}
	if !c.possible_ {
		c.place_at(pos)
	}
}

func (c *Crib) Next_right() {
	if c.crib_ == "" {
		return
	}
	pos := c.position_
	c.Move_right()
	for !c.possible_ && len(c.crib_)+c.position_ < len(c.ct_) {
		c.Move_right()
	}
	if !c.possible_ {
		c.place_at(pos)
	}
}

func (c *Crib) Clear() {
	c.crib_ = ""
	c.position_ = 0
	c.possible_ = false
	f := NewFixed_Key()
	c.fixed_key_ = &f
}

func (c *Crib) Get_pt() string {
	c.pt_ = ""
	for _, r := range c.ct_ {
		d := c.fixed_key_.Get_pt(byte(r))
		if d == byte(' ') {
			d = byte('.')
		}
		c.pt_ += string(d)
	}
	return c.pt_
}

func (c Crib) Get_crib_string() string {
	return c.crib_
}

func (c Crib) Get_fixed_key() *Fixed_Key {
	return c.fixed_key_
}

func (c Crib) Ct_has_changed(ct string) bool {
	if ct != "" && c.original_ct_ != ct {
		return true
	}
	if (ct == "" || c.original_ct_ == "") && ct != c.original_ct_ {
		return true
	}
	return false
}

func (c Crib) Get_original_ct() string {
	return c.original_ct_
}

func (c Crib) Get_position() int {
	return c.position_
}

func (c *Crib) place_at(position int) {
	if c.crib_ == "" {
		return
	}
	if position+len(c.crib_) > len(c.ct_) {
		return
	}
	if position < 0 {
		return
	}
	c.possible_ = false
	c.position_ = position
	f := NewFixed_Key()
	c.fixed_key_ = &f
	i := position
	r := c.ct_[i]
	for _, d := range c.crib_ {
		if (c.fixed_key_.Get_ct(byte(d)) != byte(' ') || c.fixed_key_.Get_pt(byte(r)) != byte(' ')) && byte(r) != c.fixed_key_.Get_ct(byte(d)) {
			return
		}
		c.fixed_key_.Set(byte(d), byte(r))
		if i < len(c.ct_)-1 {
			i++
			r = c.ct_[i]
		}
	}
	c.possible_ = true
}

func Possible_positions(ciphertext, the_crib string) map[int]string {
	the_crib_position := 0
	cc := NewCrib(ciphertext, the_crib, the_crib_position)
	if !cc.Is_possible() {
		(&cc).Next_right()
	}
	positions := make(map[int]string, 0)
	last_pos := -1
	for cc.Is_possible() {
		pos := cc.Get_position()
		if pos == last_pos {
			break
		}
		positions[pos] = (&cc).Get_pt()
		if pos+len(the_crib) >= len(ciphertext) {
			break
		}
		(&cc).Next_right()
		last_pos = pos
	}
	return positions
}
