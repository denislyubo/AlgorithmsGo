package numberof1its

func hammingWeight(num uint32) int {
	var count int
	for k := 0; k < 32; k++ {
		if (num & (1 << k)) == (1 << k) {
			count++
		}
	}

	return count
}
