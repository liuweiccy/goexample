package test

import "testing"

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
