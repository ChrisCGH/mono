package mono

import (
	"testing"
)

func TestNewTop(t *testing.T) {
	for i := 1; i < 3; i++ {
		top := NewTop(i)
		if top.max_top_ != i {
			t.Errorf("top.max_top_ should be %d, but actually is %d\n", i, top.max_top_)
		}
		if top.min_top_score_ != -1.0 {
			t.Errorf("top.min_top_score_ should be -1.0, but actually is %.2f\n", top.min_top_score_)
		}
		if top.max_top_score_ != -1.0 {
			t.Errorf("top.max_top_score_ should be -1.0, but actually is %.2f\n", top.max_top_score_)
		}
		if len(top.top_) != i {
			t.Errorf("len(top.top_) should be %d, but actually is %d\n", i, len(top.top_))
		}
		if len(top.score_) != i {
			t.Errorf("len(top.score_) should be %d, but actually is %d\n", i, len(top.score_))
		}
		if top.Size() != i {
			t.Errorf("top.Size() should be %d, but actually is %d\n", i, top.Size())
		}
	}
}

func TestAdd(t *testing.T) {
	top := NewTop(3)
	if top.Size() != 3 {
		t.Errorf("top.Size() should be %d, but actually is %d\n", 3, top.Size())
	}
	f := NewFixed_Key()
	alp1 := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	(&alp1).Randomise(f)
	score1 := 100
	alp2 := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	(&alp2).Randomise(f)
	score2 := 200
	alp3 := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	(&alp3).Randomise(f)
	score3 := 300
	alp4 := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	(&alp4).Randomise(f)
	score4 := 400
	alp5 := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	(&alp5).Randomise(f)
	score5 := 250
	alp6 := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	(&alp6).Randomise(f)
	score6 := 150
	alp7 := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	(&alp7).Randomise(f)
	score7 := 250
	alp8 := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	(&alp8).Randomise(f)
	score8 := 299

	if !top.Qualifies(score1) {
		t.Errorf("Score %d should qualify\n", score1)
	}
	(&top).Add(alp1, score1)
	if top.min_top_score_ != 0.0 {
		t.Errorf("top.min_top_score_ should be 0, but actually is %.2f\n", top.min_top_score_)
	}
	if top.max_top_score_ != float64(score1) {
		t.Errorf("top.max_top_score_ should be %d, but actually is %.2f\n", score1, top.max_top_score_)
	}
	if top.Alphabet(0).Alphabet() != alp1.Alphabet() {
		t.Errorf("top.Alphabet(0) should be %v, but actually is %v\n", alp1, top.Alphabet(0))
	}
	empty_alphabet := Alphabet{}
	if top.Alphabet(1).Alphabet() != empty_alphabet.Alphabet() {
		t.Errorf("top.Alphabet(1) should be %v, but actually is %v\n", empty_alphabet, top.Alphabet(1))
	}
	if top.Alphabet(-1).Alphabet() != empty_alphabet.Alphabet() {
		t.Errorf("top.Alphabet(-1) should be %v, but actually is %v\n", empty_alphabet, top.Alphabet(-1))
	}
	if top.Alphabet(3).Alphabet() != empty_alphabet.Alphabet() {
		t.Errorf("top.Alphabet(3) should be %v, but actually is %v\n", empty_alphabet, top.Alphabet(3))
	}

	(&top).Add(alp2, score2)
	if top.min_top_score_ != 0.0 {
		t.Errorf("top.min_top_score_ should be 0, but actually is %.2f\n", top.min_top_score_)
	}
	if top.max_top_score_ != float64(score2) {
		t.Errorf("top.max_top_score_ should be %d, but actually is %.2f\n", score2, top.max_top_score_)
	}

	(&top).Add(alp3, score3)
	if top.min_top_score_ != float64(score1) {
		t.Errorf("top.min_top_score_ should be %d, but actually is %.2f\n", score1, top.min_top_score_)
	}
	if top.max_top_score_ != float64(score3) {
		t.Errorf("top.max_top_score_ should be %d, but actually is %.2f\n", score3, top.max_top_score_)
	}

	(&top).Add(alp4, score4)
	if top.min_top_score_ != float64(score2) {
		t.Errorf("top.min_top_score_ should be %d, but actually is %.2f\n", score2, top.min_top_score_)
	}
	if top.max_top_score_ != float64(score4) {
		t.Errorf("top.max_top_score_ should be %d, but actually is %.2f\n", score4, top.max_top_score_)
	}

	if !top.Qualifies(score5) {
		t.Errorf("Score %d should qualify\n", score5)
	}
	(&top).Add(alp5, score5)
	if top.min_top_score_ != float64(score5) {
		t.Errorf("top.min_top_score_ should be %d, but actually is %.2f\n", score5, top.min_top_score_)
	}
	if top.max_top_score_ != float64(score4) {
		t.Errorf("top.max_top_score_ should be %d, but actually is %.2f\n", score4, top.max_top_score_)
	}
	if top.Qualifies(score6) {
		t.Errorf("Score %d should not qualify\n", score6)
	}
	(&top).Add(alp6, score6)
	if top.min_top_score_ != float64(score5) {
		t.Errorf("top.min_top_score_ should be %d, but actually is %.2f\n", score5, top.min_top_score_)
	}
	if top.max_top_score_ != float64(score4) {
		t.Errorf("top.max_top_score_ should be %d, but actually is %.2f\n", score4, top.max_top_score_)
	}

	if top.Score(0) != score4 {
		t.Errorf("top.Score(0) should be %d, but actually is %.2f\n", score4, top.Score(0))
	}
	if top.Alphabet(0).Alphabet() != alp4.Alphabet() {
		t.Errorf("top.Alphabet(0) should be %v, but actually is %v\n", alp4, top.Alphabet(0))
	}
	if top.Score(1) != score3 {
		t.Errorf("top.Score(1) should be %d, but actually is %.2f\n", score3, top.Score(1))
	}
	if top.Alphabet(1).Alphabet() != alp3.Alphabet() {
		t.Errorf("top.Alphabet(1) should be %v, but actually is %v\n", alp3, top.Alphabet(1))
	}
	if top.Score(2) != score5 {
		t.Errorf("top.Score(2) should be %d, but actually is %.2f\n", score5, top.Score(2))
	}
	if top.Alphabet(2).Alphabet() != alp5.Alphabet() {
		t.Errorf("top.Alphabet(2) should be %v, but actually is %v\n", alp5, top.Alphabet(2))
	}
	if top.Score(-1) != 0 {
		t.Error("top.Score(-1) should be 0, but actually is %d\n", top.Score(-1))
	}
	if top.Score(3) != 0 {
		t.Error("top.Score(3) should be 0, but actually is %d\n", top.Score(3))
	}
	(&top).Add(alp7, score7)
	if top.Alphabet(2).Alphabet() != alp5.Alphabet() {
		t.Errorf("top.Alphabet(2) should be %v, but actually is %v\n", alp5, top.Alphabet(2))
	}

	top2 := NewTop(3)
	if top2.Equal(top) {
		t.Error("top2 should not be equal to top\n")
	}
	(&top2).Add(alp1, score1)
	(&top2).Add(alp1, score2)
	(&top2).Add(alp1, score3)
	(&top2).Add(alp1, score4)
	(&top2).Add(alp1, score5)
	if !top2.Equal(top) {
		t.Error("top2 should be equal to top\n")
	}

	top3 := NewTop(3)
	(&top3).Add(alp1, score1)
	(&top3).Add(alp1, score2)
	(&top3).Add(alp1, score8)
	(&top3).Add(alp1, score4)
	(&top3).Add(alp1, score5)
	if top2.Equal(top3) {
		t.Error("top2 should not be equal to top3\n")
	}

	(&top2).Clear()
	for i := 0; i < 3; i++ {
		if top2.Score(i) != -1 {
			t.Errorf("top2.Score(i) should be %d, but actually is %.2f\n", -1, top2.Score(i))
		}
		if top2.Alphabet(i).Alphabet() != empty_alphabet.Alphabet() {
			t.Errorf("top2.Alphabet(i) should be %v, but actually is %v\n", empty_alphabet, top2.Alphabet(i))
		}
	}
}
