package fz

type GlobalAlignment struct {
	gap      int
	match    int
	mismatch int
}

func NewGlobalAlignmentCalculator(gap, match, mismatch int) Calculator {
	return &GlobalAlignment{
		gap:      gap,
		match:    match,
		mismatch: mismatch,
	}
}

func (g *GlobalAlignment) Calculate(s1, s2 string) (int, error) {
	m := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		m[i] = make([]int, len(s2)+1)
		m[i][0] = i * (-g.gap)
	}
	for j := 0; j <= len(s2); j++ {
		m[0][j] = j * (-g.gap)
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			s1Gap := m[i-1][j] - g.gap
			s2Gap := m[i][j-1] - g.gap
			match := m[i-1][j-1]
			if s1[i-1] == s2[j-1] {
				match += g.match
			} else {
				match -= g.mismatch
			}
			m[i][j] = max(s1Gap, s2Gap, match)
		}
	}

	return m[len(s1)][len(s2)], nil
}

type LocalAlignment struct {
	gap      int
	match    int
	mismatch int
}

func NewLocalAlignmentCalculator(gap, match, mismatch int) Calculator {
	return &LocalAlignment{}
}

func (l *LocalAlignment) Calculate(s1, s2 string) (int, error) {
	m := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		m[i] = make([]int, len(s2)+1)
		m[i][0] = i * (-l.gap)
	}
	for j := 0; j <= len(s2); j++ {
		m[0][j] = j * (-l.gap)
	}

	score := 0
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			s1gap := m[i-1][j] - l.gap
			s2gap := m[i][j-1] - l.gap
			match := m[i-1][j-1]
			if s1[i-1] == s2[j-1] {
				match += l.match
			} else {
				match -= l.mismatch
			}
			m[i][j] = max(s1gap, s2gap, match, 0)
			score = max(score, m[i][j])
		}
	}

	return score, nil
}
