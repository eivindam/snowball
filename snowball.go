package snowball

import (
	"fmt"

	"github.com/eivindam/snowball/english"
	"github.com/eivindam/snowball/french"
	"github.com/eivindam/snowball/norwegian"
	"github.com/eivindam/snowball/russian"
	"github.com/eivindam/snowball/spanish"
	"github.com/eivindam/snowball/swedish"
)

const (
	VERSION string = "v0.3.4"
)

// Stem a word in the specified language.
//
func Stem(word, language string, stemStopWords bool) (stemmed string, err error) {

	var f func(string, bool) string
	switch language {
	case "english":
		f = english.Stem
	case "spanish":
		f = spanish.Stem
	case "french":
		f = french.Stem
	case "russian":
		f = russian.Stem
	case "swedish":
		f = swedish.Stem
	case "norwegian":
		f = norwegian.Stem
	default:
		err = fmt.Errorf("Unknown language: %s", language)
		return
	}
	stemmed = f(word, stemStopWords)
	return

}
