package mono

import (
	"fmt"
	"io"
)

type Mono struct {
	key_ string
}

func NewMono() Mono {
	m := Mono{"abcdefghijklmnopqrstuvwxyz"}
	return m
}

func (m *Mono) Set_key(key string) {
	m.key_ = key
}

func (m Mono) get_ct_key() string {
	CT := make([]byte, 26)
	for i := 0; i < 26; i++ {
		CT[m.key_[i]-byte('a')] = byte('A') + byte(i)
	}
	return string(CT)
}

func (m Mono) Decode(ct string) string {
	outtext := make([]byte, len(ct))
	j := 0
	for _, c := range ct {
		i := byte(c) - byte('A')
		if byte(c) == byte('z')+1 {
			outtext[j] = byte(c)
		} else {
			outtext[j] = m.key_[i]
		}
		j++
	}
	return string(outtext)
}

func (m Mono) Display(w io.Writer) {
	fmt.Fprintln(w, m.get_ct_key())
	fmt.Fprintln(w, "abcdefghijklmnopqrstuvwxyz")
}
