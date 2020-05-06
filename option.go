package fz

type Option func(f *FuzzyFinder)

func UseLevenshteinAlg() Option {
	return func(f *FuzzyFinder) {
		f.alg = AlgorithmLevenshtein
	}
}

func UseGlobalAlignAlg() Option {
	return func(f *FuzzyFinder) {
		f.alg = AlgorithmGlobalAlign
	}
}

func UseLocalAlignAlg() Option {
	return func(f *FuzzyFinder) {
		f.alg = AlgorithmLocalAlign
	}
}

func SetAlignParam(gap, match, mismatch int) Option {
	return func(f *FuzzyFinder) {
		f.gap = gap
		f.match = match
		f.mismatch = mismatch
	}
}

func SetMaxGoRoutine(n uint32) Option {
	return func(f *FuzzyFinder) {
		f.maxGoroutine = n
	}
}
