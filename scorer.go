package mono

import (
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
)

type Scorer interface {
    Spaces_scored() bool
    Reset()
    Set_file(string)
    Get_scored_count() int
    Score(string, bool) int
}

type Trigraph_Scorer struct {
    scored_count_ int
    spaces_scored_ bool
    trigram_filename_ string
    trigram_table_ [27][27][27]float64
}

func (scorer Trigraph_Scorer) Spaces_scored() bool {
    return scorer.spaces_scored_
}

func (scorer Trigraph_Scorer) Reset() {
    scorer.scored_count_ = 0
}

func (scorer *Trigraph_Scorer) Set_file(filename string) {
    if filename == scorer.trigram_filename_ {
        scorer.scored_count_ = 0
        return
    }
    scorer.trigram_filename_ = filename
    scorer.spaces_scored_ = false
    // read file 
    f, err := os.Open(scorer.trigram_filename_)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)

    for ; scanner.Scan(); {
        // format is XXX : NNNNN PCT
        parts := strings.Split(strings.ToLower(scanner.Text()), " ")

        trigram := parts[0]
        c1 := trigram[0] - byte('a')
        if c1 == byte('z') + 1 {
            scorer.spaces_scored_ = true
        }
        c2 := trigram[1] - byte('a')
        c3 := trigram[2] - byte('a')

        trigram_score, err := strconv.ParseFloat(parts[3], 64)
        if err == nil {
            scorer.trigram_table_[c1][c2][c3] = trigram_score
        } else {
            scorer.trigram_table_[c1][c2][c3] = 0.0
        }
    }
    err = scanner.Err()
    if err != nil {
        log.Fatal(err)
    }
}

func (scorer *Trigraph_Scorer) Score(plaintext string, debug bool) int {
    scorer.scored_count_++
    score := 0.0
    l := len(plaintext)
    for i := 0; i < l - 2; i++ {
        c1 := plaintext[i] - byte('a')
        c2 := plaintext[i + 1] - byte('a')
        c3 := plaintext[i + 2] - byte('a')
        score += scorer.trigram_table_[c1][c2][c3]
    }
    return (int)(100.0*score)
}

func (scorer Trigraph_Scorer) Get_scored_count() int {
    return scorer.scored_count_
}

type Tetragraph_Scorer struct {
    scored_count_ int
    spaces_scored_ bool
    tetragram_filename_ string
    tetragram_table_ [27][27][27][27]float64
}

func (scorer Tetragraph_Scorer) Spaces_scored() bool {
    return scorer.spaces_scored_
}

func (scorer Tetragraph_Scorer) Reset() {
    scorer.scored_count_ = 0
}

func (scorer *Tetragraph_Scorer) Set_file(filename string) {
    if filename == scorer.tetragram_filename_ {
        scorer.scored_count_ = 0
        return
    }
    scorer.tetragram_filename_ = filename
    scorer.spaces_scored_ = false
    // read file 
    f, err := os.Open(scorer.tetragram_filename_)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)

    for ; scanner.Scan(); {
        // format is XXXX : NNNNN PCT
        parts := strings.Split(strings.ToLower(scanner.Text()), " ")

        tetragram := parts[0]
        c1 := tetragram[0] - byte('a')
        if c1 == byte('z') + 1 {
            scorer.spaces_scored_ = true
        }
        c2 := tetragram[1] - byte('a')
        c3 := tetragram[2] - byte('a')
        c4 := tetragram[3] - byte('a')

        tetragram_score, err := strconv.ParseFloat(parts[3], 64)
        if err == nil {
            scorer.tetragram_table_[c1][c2][c3][c4] = tetragram_score
        } else {
            scorer.tetragram_table_[c1][c2][c3][c4] = 0.0
        }
    }
    err = scanner.Err()
    if err != nil {
        log.Fatal(err)
    }
}

func (scorer *Tetragraph_Scorer) Score(plaintext string, debug bool) int {
    scorer.scored_count_++
    score := 0.0
    l := len(plaintext)
    for i := 0; i < l - 3; i++ {
        c1 := plaintext[i] - byte('a')
        c2 := plaintext[i + 1] - byte('a')
        c3 := plaintext[i + 2] - byte('a')
        c4 := plaintext[i + 3] - byte('a')
        score += scorer.tetragram_table_[c1][c2][c3][c4]
    }
    return (int)(100.0*score)
}

func (scorer Tetragraph_Scorer) Get_scored_count() int {
    return scorer.scored_count_
}

type Ngraph_Scorer struct {
    scored_count_ int
    spaces_scored_ bool
    ngraph_filename_ string
    ngraph_length_ int
    ngraph_table_ map[string]float64
    prev_plaintext_ string
    prev_score_  float64
    prev_ngraph_score_ []float64
}

func (scorer Ngraph_Scorer) Spaces_scored() bool {
    return scorer.spaces_scored_
}

func (scorer Ngraph_Scorer) Reset() {
    scorer.scored_count_ = 0
}

func (scorer *Ngraph_Scorer) Set_file(filename string) {
    if filename == scorer.ngraph_filename_ {
        scorer.scored_count_ = 0
        return
    }
    scorer.ngraph_filename_ = filename
    scorer.spaces_scored_ = false
    // read file 
    f, err := os.Open(scorer.ngraph_filename_)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)

    scorer.ngraph_table_ = make(map[string]float64)
    for ; scanner.Scan(); {
        // format is XXXXX : NNNNN PCT
        parts := strings.Split(strings.ToLower(scanner.Text()), " ")

        ngraph := parts[0]
        scorer.ngraph_length_ = len(ngraph)

        c1 := ngraph[0] - byte('a')
        if c1 == byte('z') + 1 {
            scorer.spaces_scored_ = true
        }

        ngraph_score, err := strconv.ParseFloat(parts[3], 64)
        if err == nil {
            scorer.ngraph_table_[ngraph] = ngraph_score
        } else {
            scorer.ngraph_table_[ngraph] = 0.0
        }
    }
    err = scanner.Err()
    if err != nil {
        log.Fatal(err)
    }
}

func (scorer *Ngraph_Scorer) Score(plaintext string, debug bool) int {
    scorer.scored_count_++
    score := 0.0
    l := len(plaintext)
    for i := 0; i < l - scorer.ngraph_length_ + 1; i++ {
        ngraph := plaintext[i:i+scorer.ngraph_length_]
        score += scorer.ngraph_table_[ngraph]
    }
    return (int)(100.0*score)
}

func (scorer Ngraph_Scorer) Get_scored_count() int {
    return scorer.scored_count_
}


