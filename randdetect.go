package randdetect

import "strings"

const (
	commonBigramsThreshold     = 0.3
	uncommonBigramsThreshold   = 0.5
	duplicatedBigramsThreshold = 0.33
)

// IsRandom returns true if a string looks like being randomly generated.
func IsRandom(s string) bool {
	// Consider only strings which are longer than 3 chars.
	if len(s) < 4 {
		return false
	}

	s = strings.ToLower(s)

	bigrams := []string{}
	rs := []rune(s)
	prev := string(rs[0])
	for _, c := range rs[1:] {
		if string(c) == " " {
			continue
		}
		bigrams = append(bigrams, prev+string(c))
		prev = string(c)
	}

	numBigrams := len(bigrams)
	numCommon := 0
	dups := map[string]int{}
	for _, b := range bigrams {
		if eb, ok := englishBigrams[b]; ok {
			if eb > commonBigramsThreshold {
				numCommon += 1
			}
		}
		dups[b] += 1
	}
	if len(dups) == 1 {
		return true
	}
	numUncommon := numBigrams - numCommon
	numDups := numBigrams - len(dups)

	// Number of uncommon bigrams dominates, in this case is a high chance that the string
	// might look random.
	if float32(numUncommon)/float32(numBigrams) > uncommonBigramsThreshold {
		return true
	}

	// Number of duplicated bigrams is more than one third in this case is a high chance that
	// the might look random.
	if float32(numDups)/float32(numBigrams) > duplicatedBigramsThreshold {
		return true
	}

	return false
}
