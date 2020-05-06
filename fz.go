package fz

import (
	"sort"
	"sync"
)

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

	mu           *sync.Mutex
	maxGoroutine uint32
}

func New() FuzzyFinder {
	return FuzzyFinder{
		alg:          AlgorithmLevenshtein,
		gap:          2,
		match:        2,
		mismatch:     1,
		mu:           &sync.Mutex{},
		maxGoroutine: 1,
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

	limit := make(chan struct{}, f.maxGoroutine)
	wg := sync.WaitGroup{}
	for _, c := range cands {
		wg.Add(1)
		go func(query string, cand Candidate) {
			limit <- struct{}{}
			defer wg.Done()
			score, gErr := calc.Calculate(query, cand.String())
			if gErr != nil {
				err = gErr
			}
			f.mu.Lock()
			rets = append(rets, Result{
				Cand:  cand,
				Score: score,
			})
			f.mu.Unlock()
			<-limit
		}(query, c)
	}
	wg.Wait()
	if err != nil {
		return nil, err
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
