package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	if longestCommonPrefix([]string{"flower", "flow", "flight"}) != "fl" {
		t.Error(`longestCommonPrefix([]string{"flower","flow","flight"}) == "fl"`)
	}

	if longestCommonPrefix([]string{"flower", "den"}) != "" {
		t.Error(`longestCommonPrefix([]string{"flower", "den"}) == ""`)
	}
}
