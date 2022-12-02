package reversebits

func reverseBits(n uint32) uint32 {
	var m uint32
	num := 32

	for num > 0 {
		if n%2 == 1 && n > 0 {
			m = m*2 + 1
		} else {
			m = m*2 + 0
		}

		if n > 0 {
			n = n >> 1
		}
		num--
	}

	return m
}
