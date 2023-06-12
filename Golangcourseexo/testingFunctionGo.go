package main

import (
	"errors"
	"fmt"
	"testing"
)

func match(pat string, str string) (bool, error) {
	if len(pat) != len(str) {
		return false, errors.New("pattern and string have different lengths")
	}
	for i := 0; i < len(pat); i++ {
		if pat[i] != str[i] {
			return false, nil
		}
	}
	return true, nil
}

func TestMatch(t *testing.T) {
	testCases := []struct {
		pat  string
		str  string
		want bool
	}{
		{"abc", "abc", true},
		{"abc", "abd", false},
		{"", "", true},
	}
	for _, tc := range testCases {
		got, _ := match(tc.pat, tc.str)
		if got != tc.want {
			t.Errorf("match(%q, %q) = %v; want %v", tc.pat, tc.str, got, tc.want)
		}
	}
}

func TestMatchDifferentLengths(t *testing.T) {
	_, err := match("abc", "abcd")
	if err == nil {
		t.Error("expected error when pattern and string have different lengths")
	}
}

func main() {

	choice, err := match("aa", "aa")

	fmt.Printf("The result for the bool comparison %t.\n Also, the return error %v\n", choice, err)

	fmt.Println("Running tests...")

	tests := []testing.InternalTest{
		{"TestMatch", TestMatch},
		{"TestMatchDifferentLengths", TestMatchDifferentLengths},
	}
	testing.Main(nil, tests, nil, nil)
}
