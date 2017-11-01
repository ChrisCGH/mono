package mono

import (
	"fmt"
)

type Fixed_Key struct {
	fixed_        [26]byte
	dummy_        byte
	not_fixed_    [27]byte
	index_        [26]int
	number_fixed_ int
}

func NewFixed_Key() Fixed_Key {
	fk := Fixed_Key{dummy_: 0}
	for i := 0; i < len(fk.fixed_); i++ {
		fk.fixed_[i] = byte(' ')
	}
	(&fk).set_index()

	return fk
}

func (f Fixed_Key) Length() int {
	return len(f.fixed_)
}

func (f Fixed_Key) NotFixedLength() int {
	return len(f.not_fixed_) - 1
}

func (f Fixed_Key) Fixed(i int) byte {
	return f.fixed_[i]
}

func (f Fixed_Key) NotFixed(i int) byte {
	return f.not_fixed_[i]
}

func (f *Fixed_Key) set_index() {
	i := 0
	f.number_fixed_ = 0
	for i = 0; i < len(f.fixed_); i++ {
		f.index_[i] = -1
	}
	for i = 0; i < len(f.fixed_); i++ {
		if f.fixed_[i] != byte(' ') {
			f.index_[f.fixed_[i]-byte('a')] = i
			f.number_fixed_++
		}
	}
	i = 0
	for c := 'a'; c <= 'z'; c++ {
		if !f.Is_set(byte(c)) {
			f.not_fixed_[i] = byte(c)
			i++
		}
	}
	f.not_fixed_[i] = 0
}

func (f Fixed_Key) Is_set(pt byte) bool {
	return (f.Get_ct(pt) != byte(' '))
}

func (f Fixed_Key) Get_pt(ct byte) byte {
	i := ct - byte('A')
	if i < 0 || int(i) >= len(f.fixed_) {
		return ' '
	}
	return f.fixed_[i]
}

func (f Fixed_Key) Get_ct(pt byte) byte {
	i := pt - byte('a')
	if i < 0 || int(i) >= len(f.fixed_) {
		return ' '
	}
	if f.index_[i] < 0 {
		return ' '
	}
	return byte(f.index_[i]) + byte('A')
}

func (f *Fixed_Key) Set(pt, ct byte) {
	if ct < byte('A') || ct > byte('Z') {
		return
	}
	if pt < byte('a') || ct > byte('z') {
		return
	}
	i := ct - byte('A')
	if i < 0 || int(i) >= len(f.fixed_) {
		return
	}
	if f.Is_set(pt) {
		f.Clear(f.Get_ct(pt), false)
	}
	f.fixed_[i] = pt
	f.set_index()
}

func (f *Fixed_Key) Clear(ct byte, do_index bool) {
	i := ct - byte('A')
	if i < 0 || int(i) >= len(f.fixed_) {
		return
	}
	if f.fixed_[i] != byte(' ') {
		f.fixed_[i] = byte(' ')
		if do_index {
			f.set_index()
		}
	}
}

func (f Fixed_Key) Number_fixed() int {
	return f.number_fixed_
}

func (f Fixed_Key) Display() {
	fmt.Printf("number fixed = %d\n", f.number_fixed_)
	for i := 0; i < len(f.fixed_); i++ {
		fmt.Printf("%s", string(f.fixed_[i]))
	}
	fmt.Println("")
	fmt.Printf("Not fixed : [%s]\n", string(f.not_fixed_[:26]))
	for i := 0; i < len(f.index_); i++ {
		fmt.Printf("%d ", f.index_[i])
	}
	fmt.Println("")
}

func (f Fixed_Key) Get_key() string {
	key := make([]byte, 26)
	for pt := byte('a'); pt <= byte('z'); pt++ {
		ct := f.Get_ct(pt)
		if ct == byte(' ') {
			key[pt-byte('a')] = byte('.')
		} else {
			key[pt-byte('a')] = ct
		}
	}
	return string(key[:26])
}

func (f Fixed_Key) check() int {
	if f.dummy_ != 0 {
		return -1
	}
	num_fixed := 0
	for i := 0; i < len(f.fixed_); i++ {
		if f.index_[i] != -1 && f.index_[i] < 0 && int(f.index_[i]) >= len(f.fixed_) {
			return -2
		}
		if f.index_[i] != -1 {
			num_fixed++
		}
	}
	if num_fixed != f.number_fixed_ {
		return -3
	}
	return 1
}

func (f *Fixed_Key) clear() {
	for i := 0; i < len(f.fixed_); i++ {
		f.fixed_[i] = byte(' ')
	}
	f.set_index()
}
