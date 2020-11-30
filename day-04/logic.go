package main

// VerifyPassword verifies that pwd fulfills all the requirements
func VerifyPassword(pwd string, maxSame int) bool {
	// It is a six-digit number
	if len(pwd) != 6 {
		return false
	}
	// adjacent digits are the same
	foundSame := false
	idxSame := 0
	for !foundSame && idxSame < len(pwd)-1 {
		// look for repeating sequence
		currentSame := 1
		subIdxSame := idxSame + 1
		keepSearching := true
		for keepSearching && subIdxSame < len(pwd) {
			if pwd[idxSame] == pwd[subIdxSame] {
				currentSame++
				subIdxSame++
			} else {
				keepSearching = false
			}
		}
		// skip the repeating part
		idxSame = subIdxSame
		// check whether we found relevant repeating sequence
		if maxSame > 0 {
			foundSame = currentSame == maxSame
		} else {
			foundSame = currentSame > 1
		}
	}
	if !foundSame {
		return false
	}
	// Going from left to right, the digits never decrease;
	isIncreasingOrSame := true
	idxIncreasing := 1
	for isIncreasingOrSame && idxIncreasing < len(pwd) {
		isIncreasingOrSame = pwd[idxIncreasing-1] <= pwd[idxIncreasing]
		idxIncreasing++
	}
	if !isIncreasingOrSame {
		return false
	}

	return true
}
