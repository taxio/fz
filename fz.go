package fz

import "sort"

type Candidate interface {
	String() string
}

type Calculator interface {
	Calculate(s1, s2 string) (int, error)
}

type Result struct {
	Cand  Candidate
	Score int
}

type FuzzyFinder struct {
	alg Algorithm

	// used by calc alignment
	gap      int
	match    int
	mismatch int
}

func New() FuzzyFinder {
	return FuzzyFinder{
		alg:      AlgorithmLevenshtein,
		gap:      2,
		match:    2,
		mismatch: 1,
	}
}

func (f *FuzzyFinder) Find(query string, cands []Candidate, opts ...Option) ([]Result, error) {
	for _, opt := range opts {
		opt(f)
	}

	rets, err := f.calcBaseScore(query, cands)
	if err != nil {
		return nil, err
	}

	return rets, nil
}

func (f *FuzzyFinder) calcBaseScore(query string, cands []Candidate) (rets []Result, err error) {
	var calc Calculator
	switch f.alg {
	case AlgorithmLevenshtein:
		calc = NewLevenshteinCalculator()
	case AlgorithmGlobalAlign:
		calc = NewGlobalAlignmentCalculator(f.gap, f.match, f.mismatch)
	case AlgorithmLocalAlign:
		calc = NewLocalAlignmentCalculator(f.gap, f.match, f.mismatch)
	}

	for _, c := range cands {
		score, err := calc.Calculate(query, c.String())
		if err != nil {
			return nil, err
		}
		rets = append(rets, Result{
			Cand:  c,
			Score: score,
		})
	}

	if f.alg == AlgorithmLevenshtein {
		sort.Slice(rets, func(i, j int) bool { return abs(rets[i].Score) < abs(rets[j].Score) })
	} else {
		sort.Slice(rets, func(i, j int) bool { return rets[i].Score > rets[j].Score })
	}

	return rets, nil
}

type Algorithm int

const (
	AlgorithmLevenshtein Algorithm = iota + 1
	AlgorithmGlobalAlign
	AlgorithmLocalAlign
)
