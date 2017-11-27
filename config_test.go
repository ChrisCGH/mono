package mono

import (
	"fmt"
	"os"
	"testing"
)

func writeBadConfigFile(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create test config file")
	}
	fmt.Fprintln(f, "monorest:")
	fmt.Fprintln(f, "  data_files:")
	fmt.Fprintln(f, "    - name: test.tri")
	fmt.Fprintln(f, "    type: trigram")
	fmt.Fprintln(f, "	path: test.tri")
}

func writeConfigFile(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create test config file")
	}
	fmt.Fprintln(f, "monorest:")
	fmt.Fprintln(f, "  data_files:")
	fmt.Fprintln(f, "    - name: test.tri")
	fmt.Fprintln(f, "      type: trigram")
	fmt.Fprintln(f, "      path: test.tri")
}

func writeTestTrigramFile(t *testing.T, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal("Failed to create test trigram file")
	}
	fmt.Fprintln(f, "THE : 0 0.99")
	fmt.Fprintln(f, "ALL : 0 0.78")
	fmt.Fprintln(f, "ZZZ : 0 0.00")
}

func TestConfig(t *testing.T) {
	err := Init()
	if err != nil {
		t.Errorf("err should be nil, but actually is %v", err)
	}

	c := NewConfig()
	n, ty, p, e := c.FindDataFile("english.tri")
	if n != "" {
		t.Errorf("n should be \"\", but actually is %s", n)
	}
	if ty != "" {
		t.Errorf("ty should be \"\", but actually is %s", ty)
	}
	if p != "" {
		t.Errorf("p should be \"\", but actually is %s", p)
	}
	if e.Error() != "Data file english.tri not found" {
		t.Errorf("e should be \"Data file english.tri not found\", but actually is %s", e.Error())
	}

	test_tri_file := "test.tri"
	writeTestTrigramFile(t, test_tri_file)
	defer os.Remove(test_tri_file)
	config_file := "monorest.yml"
	writeConfigFile(t, config_file)
	defer os.Remove(config_file)
	err = Init()
	if err != nil {
		t.Errorf("err should be nil, but actually is %v", err)
	}
	c = NewConfig()
	n, ty, p, e = c.FindDataFile("test.tri")
	if n != "test.tri" {
		t.Errorf("n should be test.tri, but actually is %s", n)
	}
	if ty != "trigram" {
		t.Errorf("ty should be trigram, but actually is %s", ty)
	}
	if p != "test.tri" {
		t.Errorf("p should be test.tri, but actually is %s", p)
	}
	if e != nil {
		t.Errorf("e should be nil, but actually is %s", e.Error())
	}

	df := c.GetDataFiles()
	if len(df) != 1 {
		t.Errorf("len(df) should be 1, but actually is %d", len(df))
	}

	bad_config_file := "monorest.yml"
	writeBadConfigFile(t, bad_config_file)
	defer os.Remove(bad_config_file)
	err = Init()
	if err == nil {
		t.Errorf("err should be not nil, but actually is nil")
	}
	c = NewConfig()
}
