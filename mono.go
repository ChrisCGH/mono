package mono

import (
    "fmt"
)

type Mono struct {
    key_ string
}

func NewMono () Mono {
    m := Mono{"abcdefghijklmnopqrstuvwxyz"}
    return m
}

func (m *Mono) Set_key (key string) {
    m.key_ = key
}

func (m Mono) Get_key () string {
    CT := make([]byte, 27)
    for i := 0; i < 26; i++ {
        CT[m.key_[i] - byte('a')] = byte('A') + byte(i)
    }
    CT[26] = 0
    return string(CT)
}

func (m Mono) Decode (ct string) string {
    outtext := make([]byte, len(ct))
    j := 0
    for _, c := range ct {
        i := byte(c) - byte('A')
        if byte(c) == byte('z') + 1 {
            outtext[j] = byte(c)
        } else {
            outtext[j] = m.key_[i]
        }
        j++
    }
    return string(outtext)
}

func (m Mono) Display () {
    fmt.Println(m.Get_key())
    fmt.Println("abcdefghijklmnopqrstuvwxyz")
}
