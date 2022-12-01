package poweroftwo

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}

	if n <= 0 {
		return false
	}

	for temp, res := 1, 0; ; temp++ {
		res = 1 << temp

		if res == n {
			return true
		} else if res > n {
			return false
		}
	}
}
