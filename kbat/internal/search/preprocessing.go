package search

import (
	"regexp"
	"strings"
)

//go:generate cog -r $GOFILE

var commonWords = map[string]struct{}{
	/* [[[cog
		import cog
		common_words = ["the", "be", "to", "of", "and", "a", "in", "that", "have", "i", "it", "for", "not", "on", "with", "he", "as", "you", "do", "at", "this", "but", "his", "by", "from"]
		for word in common_words:
			cog.outl('"' + word + '": {},')
	]]] */
	"the": {},
	"be": {},
	"to": {},
	"of": {},
	"and": {},
	"a": {},
	"in": {},
	"that": {},
	"have": {},
	"i": {},
	"it": {},
	"for": {},
	"not": {},
	"on": {},
	"with": {},
	"he": {},
	"as": {},
	"you": {},
	"do": {},
	"at": {},
	"this": {},
	"but": {},
	"his": {},
	"by": {},
	"from": {},
	// [[[end]]]
}

var punctuationRegexp = regexp.MustCompile("[" + regexp.QuoteMeta("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~") + "]")

func removePunctuation(x string) string {
	return punctuationRegexp.ReplaceAllString(x, "")
}

func textToTokens(text string) *[]string {
	x := strings.Split(text, " ")
	return &x
}

func transformTokensToLowercase(x *[]string) {
	for i, y := range *x {
		(*x)[i] = strings.ToLower(y)
	}
}

func removeCommonWords(x *[]string) {
	n := 0
	for _, y := range *x {
		if _, ok := commonWords[y]; !ok {
			(*x)[n] = y
			n += 1
		}
	}
	(*x) = (*x)[:n]
}

func preprocessText(text string) []string {
	text = removePunctuation(text)
	tokens := textToTokens(text)
	transformTokensToLowercase(tokens)
	removeCommonWords(tokens)
	return *tokens
}
