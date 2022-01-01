package bits

func bitsToNumber(bs []bool) int {
	if len(bs) == 0 {
		return 0
	}
	var n int
	// this is big endian, p.e. 100 => 4
	for i := 0; i < len(bs); i++ {
		b := bs[len(bs)-1-i]
		if !b {
			continue
		}
		n |= 1 << i
	}
	return n
}
