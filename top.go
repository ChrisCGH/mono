package mono

type Top struct {
	top_                           []Alphabet
	score_                         []float64
	min_top_score_, max_top_score_ float64
	max_top_                       int
}

func NewTop(max_top int) Top {
	return Top{
		top_:           make([]Alphabet, max_top),
		score_:         make([]float64, max_top),
		min_top_score_: -1.0,
		max_top_score_: -1.0,
		max_top_:       max_top}
}

func (t *Top) Clear() {
	for i := 0; i < len(t.score_); i++ {
		t.score_[i] = -1.0
	}
	t.min_top_score_ = 1.0
	t.max_top_score_ = 1.0
}

func (t Top) Size() int {
	return len(t.top_)
}

func (t Top) Alphabet(i int) Alphabet {
	return t.top_[i]
}

func (t Top) Qualifies(score int) bool {
	if float64(score) > t.min_top_score_ {
		return true
	}
	return false
}

func (t *Top) Add(top Alphabet, score int) {
	i := 0
	for i < len(t.top_) && t.score_[i] >= float64(score) {
		i++
	}
	if i > 0 && float64(score) == t.score_[i-1] {
		return
	}
	if i < len(t.top_) {
		for j := t.max_top_ - 1; j > i; j-- {
			t.top_[j] = t.top_[j-1]
			t.score_[j] = t.score_[j-1]
		}
		t.top_[i] = top
		t.score_[i] = float64(score)
		t.min_top_score_ = t.score_[t.max_top_-1]
		t.max_top_score_ = t.score_[0]
	}
}

func (t Top) Equal(t2 Top) bool {
	if t.max_top_ != t2.max_top_ ||
		t.min_top_score_ != t2.min_top_score_ ||
		t.max_top_score_ != t2.max_top_score_ {
		return false
	}
	for i := 0; i < len(t2.score_); i++ {
		if t.score_[i] != t2.score_[i] {
			return false
		}
	}
	return true
}

func (t Top) Score(i int) int {
	return int(t.score_[i])
}
