package test

import "unicode"

func IsPalindrome(s string) bool {
	var letters []rune
	// 提前进行内存分配，避免在append时，多次扩容分配内存
	letters = make([]rune, 0, len(s))

	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
