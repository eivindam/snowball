package french

import (
	"github.com/eivindam/snowball/snowballword"
)

// Step 5 Undouble non-vowel endings
//
func step5(word *snowballword.SnowballWord) bool {

	suffix, _ := word.FirstSuffix("enn", "onn", "ett", "ell", "eill")
	if suffix != "" {
		word.RemoveLastNRunes(1)
	}
	return false
}
