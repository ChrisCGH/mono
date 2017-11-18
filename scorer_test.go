package mono

import (
	"testing"
	"os"
	"fmt"
)

func writeTrigramFile(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create test trigram file")
	}
	fmt.Fprintln(f, "THE : 0 0.99")
	fmt.Fprintln(f, "ALL : 0 0.78")
	fmt.Fprintln(f, "ZZZ : 0 0.00")
}

func writeTrigramFileWithSpaces(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create test trigram file")
	}
	fmt.Fprintln(f, "AAA : 2 1.39032")
	fmt.Fprintln(f, "AA{ : 3 1.79578")
	fmt.Fprintln(f, "{Z{ : 1 0.697171")
}

func TestTrigraph_Scorer(t *testing.T) {
	test_tri_file := "test.tri"
	writeTrigramFile(t, test_tri_file)
	defer os.Remove(test_tri_file)
	test_tri_spaces_file := "test.tri.spaces"
	writeTrigramFileWithSpaces(t, test_tri_spaces_file)
	defer os.Remove(test_tri_spaces_file)
	scorer := &Trigraph_Scorer{}
	scorer.Set_file(test_tri_file)
	if scorer.trigram_filename_ != test_tri_file {
		t.Errorf("scorer.trigram_filename_ should be %s, but actually is %s\n", test_tri_file, scorer.trigram_filename_)
	}
	if scorer.Spaces_scored() {
		t.Error("scorer.Spaces_scored() should return false")
	}
	score := scorer.Score("thethethethethe", false)
	if score != 495 {
		t.Error("score should be 495")
	}
	score = scorer.Score("allthezzz", false)
	if score != 177 {
		t.Error("score should be 177")
	}
	if scorer.Get_scored_count() != 2 {
		t.Errorf("scorer.Get_scored_count() should be 2, but actually is %d\n", scorer.Get_scored_count())
	}
	scorer.Reset()
	score = scorer.Score("cheeseontoast", false)
	if score != 0 {
		t.Error("score should be 0")
	}
	if scorer.Get_scored_count() != 1 {
		t.Errorf("scorer.Get_scored_count() should be 1, but actually is %d\n", scorer.Get_scored_count())
	}
	scorer.Set_file(test_tri_spaces_file)
	if scorer.trigram_filename_ != test_tri_spaces_file {
		t.Errorf("scorer.trigram_filename_ should be %s, but actually is %s\n", test_tri_spaces_file, scorer.trigram_filename_)
	}
	if ! scorer.Spaces_scored() {
		t.Error("scorer.Spaces_scored() should return true")
	}
	score = scorer.Score("aaa{aa{z{aaa", false)
	if score != 706 {
		t.Errorf("score should be 706, but is %d", score)
	}
	if scorer.Get_scored_count() != 1 {
		t.Errorf("scorer.Get_scored_count() should be 1, but actually is %d\n", scorer.Get_scored_count())
	}
}

func writeTetragramFile(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create test tetragram file")
	}
	fmt.Fprintln(f, "THAT : 5307 8.92778")
	fmt.Fprintln(f, "THER : 5075 8.88308")
	fmt.Fprintln(f, "WITH : 4325 8.72317")
	fmt.Fprintln(f, "NTHE : 3957 8.63424")
	fmt.Fprintln(f, "DTHE : 3369 8.47337")
}

func writeTetragramFileWithSpaces(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create test tetragram file")
	}
	fmt.Fprintln(f, "AAAA : 1 0.697172")
	fmt.Fprintln(f, "AAAH : 1 0.697172")
	fmt.Fprintln(f, "AABO : 1 0.697172")
	fmt.Fprintln(f, "AAFT : 4 2.08347")
	fmt.Fprintln(f, "{ABA : 135 5.60245")
}

func TestTetragraph_Scorer(t *testing.T) {
	test_tet_file := "test.tet"
	writeTetragramFile(t, test_tet_file)
	defer os.Remove(test_tet_file)
	test_tet_spaces_file := "test.tet.spaces"
	writeTetragramFileWithSpaces(t, test_tet_spaces_file)
	defer os.Remove(test_tet_spaces_file)
	scorer := &Tetragraph_Scorer{}
	scorer.Set_file(test_tet_file)
	if scorer.tetragram_filename_ != test_tet_file {
		t.Errorf("scorer.tetragram_filename_ should be %s, but actually is %s\n", test_tet_file, scorer.tetragram_filename_)
	}
	if scorer.Spaces_scored() {
		t.Error("scorer.Spaces_scored() should return false")
	}
	score := scorer.Score("thatthatthattherwith", false)
	if score != 4438 {
		t.Errorf("score should be 4438, but actually is %d\n", score)
	}
	score = scorer.Score("withthernthed", false)
	if score != 2624 {
		t.Errorf("score should be 2624, but actually is %d\n", score)
	}
	if scorer.Get_scored_count() != 2 {
		t.Errorf("scorer.Get_scored_count() should be 2, but actually is %d\n", scorer.Get_scored_count())
	}
	scorer.Reset()
	score = scorer.Score("cheeseontoast", false)
	if score != 0 {
		t.Error("score should be 0")
	}
	if scorer.Get_scored_count() != 1 {
		t.Errorf("scorer.Get_scored_count() should be 1, but actually is %d\n", scorer.Get_scored_count())
	}
	scorer.Set_file(test_tet_spaces_file)
	if scorer.tetragram_filename_ != test_tet_spaces_file {
		t.Errorf("scorer.tetragram_filename_ should be %s, but actually is %s\n", test_tet_spaces_file, scorer.tetragram_filename_)
	}
	if ! scorer.Spaces_scored() {
		t.Error("scorer.Spaces_scored() should return true")
	}
	score = scorer.Score("aaft{aba{aaaa", false)
	if score != 838 {
		t.Errorf("score should be 838, but is %d", score)
	}
	if scorer.Get_scored_count() != 1 {
		t.Errorf("scorer.Get_scored_count() should be 1, but actually is %d\n", scorer.Get_scored_count())
	}
}

func writeNgramFile(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create test ngram file")
	}
	fmt.Fprintln(f, "THAT : 5307 8.92778")
	fmt.Fprintln(f, "THER : 5075 8.88308")
	fmt.Fprintln(f, "WITH : 4325 8.72317")
	fmt.Fprintln(f, "NTHE : 3957 8.63424")
	fmt.Fprintln(f, "DTHE : 3369 8.47337")
}

func writeNgramFileWithSpaces(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create test ngram file")
	}
	fmt.Fprintln(f, "THE{A : 1312 7.87648")
	fmt.Fprintln(f, "THE{B : 1465 7.98678")
	fmt.Fprintln(f, "THE{C : 3197 8.76714")
	fmt.Fprintln(f, "THE{D : 1595 8.0718")
	fmt.Fprintln(f, "{ABAB : 1 0.697172")
}

func TestNgraph_Scorer(t *testing.T) {
	test_ng_file := "test.5"
	writeNgramFile(t, test_ng_file)
	defer os.Remove(test_ng_file)
	test_ng_spaces_file := "test.ng.spaces"
	writeNgramFileWithSpaces(t, test_ng_spaces_file)
	defer os.Remove(test_ng_spaces_file)
	scorer := &Ngraph_Scorer{}
	scorer.Set_file(test_ng_file)
	if scorer.ngraph_filename_ != test_ng_file {
		t.Errorf("scorer.ngraph_filename_ should be %s, but actually is %s\n", test_ng_file, scorer.ngraph_filename_)
	}
	if scorer.Spaces_scored() {
		t.Error("scorer.Spaces_scored() should return false")
	}
	score := scorer.Score("thatthatthattherwith", false)
	if score != 4438 {
		t.Errorf("score should be 4438, but actually is %d\n", score)
	}
	score = scorer.Score("withthernthed", false)
	if score != 2624 {
		t.Errorf("score should be 2624, but actually is %d\n", score)
	}
	if scorer.Get_scored_count() != 2 {
		t.Errorf("scorer.Get_scored_count() should be 2, but actually is %d\n", scorer.Get_scored_count())
	}
	scorer.Reset()
	score = scorer.Score("cheeseontoast", false)
	if score != 0 {
		t.Error("score should be 0")
	}
	if scorer.Get_scored_count() != 1 {
		t.Errorf("scorer.Get_scored_count() should be 1, but actually is %d\n", scorer.Get_scored_count())
	}
	scorer.Set_file(test_ng_spaces_file)
	if scorer.ngraph_filename_ != test_ng_spaces_file {
		t.Errorf("scorer.ngraph_filename_ should be %s, but actually is %s\n", test_ng_spaces_file, scorer.ngraph_filename_)
	}
	if ! scorer.Spaces_scored() {
		t.Error("scorer.Spaces_scored() should return true")
	}
	score = scorer.Score("the{a{the{b{abab", false)
	if score != 1656 {
		t.Errorf("score should be 1656, but is %d", score)
	}
	if scorer.Get_scored_count() != 1 {
		t.Errorf("scorer.Get_scored_count() should be 1, but actually is %d\n", scorer.Get_scored_count())
	}
}
