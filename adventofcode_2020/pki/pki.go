package pki

var mkiDiff int = 20201227

func Calc(subj int, loopSize int) int {
	val := 1
	for i := 0; i < loopSize; i++ {
		val = val * subj
		val = val % mkiDiff
	}
	return val
}

func FindLoopSizeForSubj(subj int, pk int) (int, bool) {
	max := 100000000
	val := 1
	for ls := 0; ls <= max; ls++ {
		val = val * subj
		val = val % mkiDiff
		if val == pk {
			return ls + 1, true
		}
	}
	return 0, false
}

func FindLoopSizeForSubjNaiv(subj int, pk int) (int, bool) {
	max := 100000
	for ls := 1; ls <= max; ls++ {
		if Calc(subj, ls) == pk {
			return ls, true
		}
	}
	return 0, false
}

func FindLoopSizeAndSubj(pk int) (int, int, bool) {
	max := 10000
	for subj := 1; subj <= max; subj++ {
		if ls, ok := FindLoopSizeForSubj(subj, pk); ok {
			return ls, subj, true
		}
	}
	return 0, 0, false
}
