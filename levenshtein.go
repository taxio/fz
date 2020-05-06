package fz

type LevenshteinDistance struct{}

func NewLevenshteinCalculator() Calculator {
	return &LevenshteinDistance{}
}

func (l *LevenshteinDistance) Calculate(s1, s2 string) (int, error) {
	m := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		m[i] = make([]int, len(s2)+1)
	}

	for i := 0; i <= len(s1); i++ {
		m[i][0] = i
	}

	for j := 0; j <= len(s2); j++ {
		m[0][j] = j
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			x := 1
			if s1[i-1] == s2[j-1] {
				x = 0
			}
			y1 := m[i-1][j] + 1
			y2 := m[i][j-1] + 1
			y3 := m[i-1][j-1] + x
			m[i][j] = min(y1, y2, y3)
		}
	}

	return m[len(s1)][len(s2)], nil
}
