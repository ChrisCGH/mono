package mono

import (
    "fmt"
    "time"
    "unicode"
    "math/rand"
)

type Alphabet struct {
    alphabet_  string
    alphabet_size_ int
    c1_ byte
    c2_ byte
    byte_alphabet_ []byte
}

func NewAlphabet (a string) Alphabet {
    return Alphabet{alphabet_: a, alphabet_size_: len(a)}
}

func (a Alphabet) Display() {
    for _, r := range a.alphabet_ {
        fmt.Printf("%s ", string(r))
    }
    fmt.Printf("\n")
}

var first_time = true

func (a *Alphabet) Randomise(fixed_key Fixed_Key) {
    alphabet := make([]byte, len(a.alphabet_))
    copy(alphabet, a.alphabet_)
    if first_time {
        rand.Seed(time.Now().Unix())
        first_time = false
    }
    byte_alphabet := make([]byte, a.alphabet_size_ + 1)
    for k := 0; k < a.alphabet_size_; k++ {
        byte_alphabet[k] = byte(' ')
    }
    a.alphabet_ = "                              "
    j := a.alphabet_size_
    i := 0
    index := 0
    if fixed_key.Number_fixed() > 0 {
        for k := 0; k < fixed_key.Length(); k++ {
            byte_alphabet[k] = fixed_key.Fixed(k)
        }
        j = a.alphabet_size_ - fixed_key.Number_fixed()
        for k := 0; k < fixed_key.NotFixedLength(); k++ {
            alphabet[k] = fixed_key.NotFixed(k)
        }
        for i = 0; i < a.alphabet_size_ - fixed_key.Number_fixed(); i++ {
            num := rand.Int() % j
            for index < a.alphabet_size_ && byte_alphabet[index] != byte(' ') {
                index++
            }
            byte_alphabet[index] = byte(unicode.ToLower(rune(alphabet[num])))
            for k := 0; k < j - num - 1; k++ {
                alphabet[num + k] = alphabet[num + k + 1]
            }
            alphabet[j - 1] = 0
            j--
        }
        byte_alphabet[a.alphabet_size_] = 0
    } else {
        for i := 0; i < a.alphabet_size_; i++ {
            num := rand.Int() % j
            for index < a.alphabet_size_ && byte_alphabet[index] != byte(' ') {
                index++
            }
            byte_alphabet[index] = byte(unicode.ToLower(rune(alphabet[num])))
            for k := 0; k < j - num - 1; k++ {
                alphabet[num + k] = alphabet[num + k + 1]
            }
            alphabet[j - 1] = 0
            j--
        }
    }
    byte_alphabet[a.alphabet_size_] = 0
    a.alphabet_ = string(byte_alphabet)
}

func (a *Alphabet) Randomise1(fixed_key Fixed_Key, stuck_count int) {
    if stuck_count < 1 {
        stuck_count = 1
    }
    if stuck_count > 24 {
        stuck_count = 24
    }
    a.Randomise(fixed_key)
    byte_alphabet := make([]byte, a.alphabet_size_)
    for k := 0; k < a.alphabet_size_; k++ {
        byte_alphabet[k] = byte(a.alphabet_[k])
    }
    MAX_SWAPS := rand.Int() % stuck_count + 3
    for number_of_swaps := 0; number_of_swaps < MAX_SWAPS; number_of_swaps++ {
        i := rand.Int() % a.alphabet_size_
        j := rand.Int() % a.alphabet_size_
        for i == j || (fixed_key.Number_fixed() > 0 && (fixed_key.Is_set(byte_alphabet[i]) || fixed_key.Is_set(byte_alphabet[j]))) {
            i = rand.Int() % a.alphabet_size_
            j = rand.Int() % a.alphabet_size_
        }
        c := byte_alphabet[i]
        byte_alphabet[i] = byte_alphabet[j]
        byte_alphabet[j] = c
    }
    a.alphabet_ = string(byte_alphabet)
}

func (a Alphabet) Alphabet() string {
    return a.alphabet_
}

func (a *Alphabet) Start_swaps() {
    a.c1_ = 0
    a.c2_ = 0
    a.byte_alphabet_ = make([]byte, a.alphabet_size_)
    for k := 0; k < a.alphabet_size_; k++ {
        a.byte_alphabet_[k] = byte(a.alphabet_[k])
    }
}

func (a *Alphabet) Next_swap(fixed_key Fixed_Key) {
    if a.c2_ > a.c1_ && int(a.c2_) < a.alphabet_size_ {
        c := a.byte_alphabet_[a.c1_]
        a.byte_alphabet_[a.c1_] = a.byte_alphabet_[a.c2_]
        a.byte_alphabet_[a.c2_] = c
    }
    a.c2_++
    if int(a.c2_) >= a.alphabet_size_ {
        a.c1_++
        a.c2_ = a.c1_ + 1
    }
    for int(a.c2_) < a.alphabet_size_ && fixed_key.Number_fixed() > 0 && (fixed_key.Is_set(a.byte_alphabet_[a.c1_]) || fixed_key.Is_set(a.byte_alphabet_[a.c2_])) {
        a.c2_++
        if int(a.c2_) >= a.alphabet_size_ {
            a.c1_++
            a.c2_ = a.c1_ + 1
        }
    }
    if int(a.c2_) < a.alphabet_size_  {
        c := a.byte_alphabet_[a.c1_]
        a.byte_alphabet_[a.c1_] = a.byte_alphabet_[a.c2_]
        a.byte_alphabet_[a.c2_] = c
    }
    a.alphabet_ = string(a.byte_alphabet_)
}

func (a Alphabet) End_swaps() bool {
    return int(a.c2_) >= a.alphabet_size_
}
