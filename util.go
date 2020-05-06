package fz

func min(vn ...int) int {
	m := vn[0]
	for _, v := range vn[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func max(vn ...int) int {
	m := vn[0]
	for _, v := range vn[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
