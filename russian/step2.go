package russian

import (
	"github.com/emoland/snowball/snowballword"
)

// Step 2 is the removal of the "и" suffix.
//
func step2(word *snowballword.SnowballWord) bool {
	suffix, _ := word.RemoveFirstSuffixIn(word.RVstart, "и")
	if suffix != "" {
		return true
	}
	return false
}
