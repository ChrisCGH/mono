package mono

import (
	"fmt"
	"os"
	"time"
	"unicode"
)

type Mono_Solver struct {
	scorer_         Scorer
	mono_           Mono
	top_            Top
	fixed_          Fixed_Key
	max_top_        int
	max_iterations_ int
	timeout_        int

	verbose_     bool
	stuck_count_ int
	score_       int
	ciphertext_  string
	solution_    string
	iterations_  int
	key_         string
	elapsed_     time.Duration
}

func NewMono_Solver() Mono_Solver {
	ms := Mono_Solver{}
	ms.max_top_ = 1
	return ms
}

func (solver *Mono_Solver) Set_cipher_text(ciphertext string) {
	solver.ciphertext_ = ""
	space_last := false
	for _, r := range ciphertext {
		if unicode.IsLetter(r) {
			solver.ciphertext_ += string(unicode.ToUpper(r))
			space_last = false
		} else if solver.scorer_.Spaces_scored() && !space_last && unicode.IsSpace(r) {
			solver.ciphertext_ += string(byte('z') + 1)
			space_last = true
		}
	}
}

func (solver *Mono_Solver) Set_trigraph_scoring(trigram_file string) {
	solver.scorer_ = &Trigraph_Scorer{}
	solver.scorer_.Set_file(trigram_file)
}

func (solver *Mono_Solver) Set_tetragraph_scoring(tetragram_file string) {
	solver.scorer_ = &Tetragraph_Scorer{}
	solver.scorer_.Set_file(tetragram_file)
}

func (solver *Mono_Solver) Set_ngraph_scoring(ngraph_file string) {
	solver.scorer_ = &Ngraph_Scorer{}
	solver.scorer_.Set_file(ngraph_file)
}

func (solver *Mono_Solver) Set_verbose() {
	solver.verbose_ = true
}

func (solver *Mono_Solver) Set_max_iterations(max_iterations int) {
	solver.max_iterations_ = max_iterations
	if solver.max_iterations_ < 0 {
		solver.max_iterations_ = 0
	}
}

func (solver *Mono_Solver) Set_timeout(t int) {
	solver.timeout_ = t
	if solver.timeout_ < 0 {
		solver.timeout_ = 0
	}
}

func (solver *Mono_Solver) Set_fixed(f Fixed_Key) {
	solver.fixed_ = f
}

func (solver Mono_Solver) is_time_to_stop() bool {
	if solver.max_iterations_ > 0 && solver.scorer_.Get_scored_count() > solver.max_iterations_ {
		return true
	}
	if solver.timeout_ > 0 && solver.elapsed_.Seconds() > float64(solver.timeout_) {
		return true
	}
	return false
}

func (solver *Mono_Solver) Solve() int {
	best_local_max_score := -1000000
	score := -1000000
	done := false
	i := 0
	outtext := ""
	solver.scorer_.Reset()
	solver.top_.Clear()

	start_time := time.Now()

	random := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	best_local_max_alphabet := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
	solver.stuck_count_ = 0
	top1 := NewTop(solver.max_top_)

	for i < solver.max_top_ {
		random.Randomise(solver.fixed_)
		solver.mono_.Set_key(random.Alphabet())
		outtext = solver.mono_.Decode(solver.ciphertext_)
		score = solver.scorer_.Score(outtext, false)
		if top1.Qualifies(score) {
			top1.Add(random, score)
			i++
		}
	}

	for !done {
		top2 := NewTop(solver.max_top_)
		for g := 0; g < top1.Size(); g++ {
			local_max_found := false
			local_max_score := -1000000
			random = top1.Alphabet(g)
			for !local_max_found {
				random.Start_swaps()
				max_score := -1000000
				best_alphabet := NewAlphabet("abcdefghijklmnopqrstuvwxyz")
				for !random.End_swaps() {
					solver.mono_.Set_key(random.Alphabet())
					outtext = solver.mono_.Decode(solver.ciphertext_)
					score = solver.scorer_.Score(outtext, false)
					if score > max_score {
						max_score = score
						best_alphabet = random
						if max_score > best_local_max_score {
							best_local_max_score = max_score
							best_local_max_alphabet = best_alphabet
							solver.stuck_count_ = 0
							now_time := time.Now()
							elapsed := now_time.Sub(start_time)
							solver.score_ = best_local_max_score
							solver.solution_ = outtext
							solver.key_ = best_local_max_alphabet.Alphabet()
							solver.iterations_ = solver.scorer_.Get_scored_count()
							solver.elapsed_ = elapsed
							if solver.verbose_ {
								fmt.Printf("%d alphabets in %v\n", solver.scorer_.Get_scored_count(), elapsed)
								solver.mono_.Display(os.Stdout)
								fmt.Printf("Score = %d\n", best_local_max_score)
								fmt.Println(outtext)
							}
						}
					}
					if top2.Qualifies(score) {
						top2.Add(random, score)
					}
					if solver.top_.Qualifies(score) {
						solver.top_.Add(random, score)
					}
					random.Next_swap(solver.fixed_)
				}
				if max_score > local_max_score {
					local_max_score = max_score
					random = best_alphabet
					solver.mono_.Set_key(random.Alphabet())
				} else {
					solver.stuck_count_++
					local_max_found = true
				}
			}
		}
		if solver.is_time_to_stop() {
			now_time := time.Now()
			elapsed := now_time.Sub(start_time)
			solver.mono_.Set_key(best_local_max_alphabet.Alphabet())
			outtext = solver.mono_.Decode(solver.ciphertext_)
			if solver.verbose_ {
				fmt.Printf(">>>> %d alphabets in %v\n", solver.scorer_.Get_scored_count(), elapsed)
				solver.mono_.Display(os.Stdout)
				fmt.Printf("Score = %d\n", best_local_max_score)
				fmt.Println(outtext)
			}
			solver.score_ = best_local_max_score
			solver.solution_ = outtext
			solver.key_ = best_local_max_alphabet.Alphabet()
			solver.iterations_ = solver.scorer_.Get_scored_count()
			solver.elapsed_ = elapsed
			return 0
		}
		if top1.Equal(top2) {
			top1.Clear()
			i := 0
			for i < solver.max_top_ {
				random = best_local_max_alphabet
				random.Randomise1(solver.fixed_, solver.stuck_count_)
				solver.mono_.Set_key(random.Alphabet())
				outtext = solver.mono_.Decode(solver.ciphertext_)
				score = solver.scorer_.Score(outtext, false)
				if top1.Qualifies(score) {
					top1.Add(random, score)
					i++
				}
			}
		} else {
			top1 = top2
		}
	}
	return 0
}

func (solver Mono_Solver) Solution() string {
	return solver.solution_
}

func (solver Mono_Solver) Score() int {
	return solver.score_
}

func (solver Mono_Solver) Key() string {
	return solver.key_
}

func (solver Mono_Solver) Iterations() int {
	return solver.iterations_
}

func (solver Mono_Solver) Elapsed() time.Duration {
	return solver.elapsed_
}
