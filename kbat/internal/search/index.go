package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
)

type Index struct {
	currentIdentifier int
	sourceDirectory   string
	Documents         map[string]*Document
	Index             map[string][]string
}

func NewIndex(sourceDirectory string) *Index {
	return &Index{
		sourceDirectory: filepath.Join(sourceDirectory, "_index"),
		Documents:       make(map[string]*Document),
		Index:           make(map[string][]string),
	}
}

type Document struct {
	ID              string
	Path            string
	WordFrequencies map[string]int
}

func NewDocument(path, text string) *Document {
	processed := preprocessText(text)
	wordFrequencies := make(map[string]int)

	for _, word := range processed {
		wordFrequencies[word] = wordFrequencies[word] + 1
	}

	return &Document{
		Path:            path,
		WordFrequencies: wordFrequencies,
	}
}

func (i *Index) nextIdentifier() string {
	i.currentIdentifier += 1
	return fmt.Sprintf("%07d", i.currentIdentifier)
}

// AddDocument adds a *Document to the index. The document will have the ID
// overwritten, and no field should be modified after insertion.
func (i *Index) AddDocument(d *Document) {
	d.ID = i.nextIdentifier()
	i.Documents[d.ID] = d
	for word := range d.WordFrequencies {
		i.Index[word] = append(i.Index[word], d.ID)
	}
}

func (i *Index) FromDisk() error {
	fcont, err := ioutil.ReadFile(filepath.Join(i.sourceDirectory, "index.json"))
	if err != nil {
		return err
	}
	return json.Unmarshal(fcont, i)
}

func (i *Index) ToDisk() error {
	jsonData, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(i.sourceDirectory, 0777); err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	return ioutil.WriteFile(filepath.Join(i.sourceDirectory, "index.json"), jsonData, 0644)
}

func (i *Index) documentFrequency(token string) int {
	return len(i.Index[token])
}

func (i *Index) inverseDocumentFrequency(token string) float64 {
	return math.Log10(float64(len(i.Documents)) / float64(i.documentFrequency(token)))
}
