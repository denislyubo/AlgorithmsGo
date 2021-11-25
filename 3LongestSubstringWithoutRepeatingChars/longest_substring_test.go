package longestsubstringwithoutrepeatingchars

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	if lengthOfLongestSubstring("abcabcbb") != 3 {
		t.Error("incorrect solution")
	}
}

func TestLengthOfLongestSubstringAnother(t *testing.T) {
	if lengthOfLongestSubstring("pwwkew") != 3 {
		t.Error("incorrect solution")
	}
}

func TestLengthOfLongestSubstringGap(t *testing.T) {
	if lengthOfLongestSubstring(" ") != 1 {
		t.Error("incorrect solution")
	}
}

func TestLengthOfLongestSubstringAAB(t *testing.T) {
	if lengthOfLongestSubstring("aab") != 2 {
		t.Error("incorrect solution")
	}
}
