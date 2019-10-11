package test

import (
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("kayak") {
		t.Error("IsPalindrome(kayak) == false")
	}

	if !IsPalindrome("qwertyytrewq") {
		t.Error("IsPalindrome(qwertyytrewq) == false")
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("ssfashae") {
		t.Error("IsPalindrome(ssfashae) == true")
	}
}

func TestChinesePalindrome(t *testing.T) {
	input := "刘威刘"
	if !IsPalindrome(input) {
		t.Errorf("IsPalindrome(%q) == false", input)
	}
}

func TestIsPalindrome2(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"ab", false},
		{"aba", true},
		{"aa", true},
		{"", true},
		{"", true},
		{"", true},
		{"", true},
		{"", true},
		{"", true},
		{"", true},
		{"", true},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v\n", test.input, got)
		}
	}
}

func randomPalindrome(rand *rand.Rand) string {
	n := rand.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rand.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestIsPalindromeRandom(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Rondom seed: %d\n", seed)
	rand := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000000; i++ {
		p := randomPalindrome(rand)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false\n", p)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("abssba")
	}
}
