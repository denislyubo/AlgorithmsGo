package permutationinstring

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Count, s2Count := [26]int{}, [26]int{}

	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	var matches int
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}
	if matches == 26 {
		return true
	}
	var indexL, indexR uint8
	for l, r := 0, len(s1); r < len(s2); {
		indexL = s2[l] - 'a'
		if s2Count[indexL] == s1Count[indexL] {
			matches--
		}
		s2Count[indexL]--
		if s2Count[indexL] == s1Count[indexL] {
			matches++
		}

		l++
		indexR = s2[r] - 'a'
		if s2Count[indexR] == s1Count[indexR] {
			matches--
		}
		s2Count[indexR]++
		if s2Count[indexR] == s1Count[indexR] {
			matches++
		}

		r++

		if matches == 26 {
			return true
		}
	}

	return false
}
