package search

import (
	"sort"
)

type Result struct {
	Document *Document
	Score    float64
}

type Results []*Result

func (r Results) Len() int {
	return len(r)
}

func (r Results) Less(i, j int) bool {
	return r[i].Score > r[j].Score
}

func (r Results) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (i *Index) Search(query string) Results {

	processedQuery := preprocessText(query)

	var resultIDs [][]string
	{
		for _, token := range processedQuery {
			x, found := i.Index[token]
			if found {
				resultIDs = append(resultIDs, x)
			}
		}
	}

	var documents []*Document
	{
		seen := make(map[string]struct{})
		for _, x := range resultIDs {
			for _, y := range x {
				if _, found := seen[y]; !found {
					documents = append(documents, i.Documents[y])
				}
			}
		}
	}

	return i.rank(processedQuery, documents)
}

func (i *Index) rank(processedQuery []string, documents []*Document) Results {

	var o Results

	for _, document := range documents {
		var score float64
		for _, token := range processedQuery {
			tf := float64(document.WordFrequencies[token])
			idf := i.inverseDocumentFrequency(token)
			score += tf * idf
		}
		o = append(o, &Result{
			Document: document,
			Score:    score,
		})
	}

	sort.Sort(o)

	return o
}
