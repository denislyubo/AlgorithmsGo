package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	var prefix string = ""

out:
	for j := 0; j < len(strs[0]); j++ {
		curSymbol := strs[0][j]

		for i := 1; i < len(strs); i++ {
			if j < len(strs[i]) {
				if strs[i][j] != curSymbol {
					break out
				}
			} else {
				break out
			}
		}

		prefix += string(curSymbol)
	}

	return prefix
}
