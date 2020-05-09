package gofrequencyanalysis

/*
 * This package enable to calculate the occurence
 * of each character contained in given text.
 */

import (
	"sort"
)

//IFreqAnalysis : Interface
type IFreqAnalysis interface {
	Sort()
	Mapping()
	Frequency()
	Analysis()
}

//Pair : Represent some data around char, its number and occurence frequency.
type Pair struct {
	char  string
	count int
	freq  float32
}

/*Counter : Contain all data around text like :
 * The original text
 * the sizeof the text
 * mappedKeys which is the number of time a key appear in the text
 * array of Pair (see Pair)
 */
type Counter struct {
	text       string
	sizeText   int
	mappedKeys map[string]int
	pairs      []Pair
}

//Sort : Enable to sorting the Pairs
func (counter *Counter) Sort() {
	var pairs []Pair
	//Initialization of Pairs
	for k, v := range counter.mappedKeys {
		pairs = append(pairs, Pair{k, v, 0})
	}
	//Sorting
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})
	counter.pairs = pairs
}

//Mapping : Count each character in text
func (counter *Counter) Mapping() {

	counter.sizeText = len(counter.text)

	countedKeys := make(map[string]int)

	for i := 0; i < counter.sizeText; i++ {
		key := counter.text[i : i+1]
		countedKeys[key]++
	}
	counter.mappedKeys = countedKeys
}

//Frequency : Calculate the occurrence frequency
func (counter *Counter) Frequency() {

	for i := 0; i < len(counter.pairs); i++ {

		counter.pairs[i].freq = float32(counter.pairs[i].count) / float32(counter.sizeText) * float32(100)
	}
}

//Analysis : Automated the process to mapping, sorting and finally calcule frequency
func (counter *Counter) Analysis() {
	counter.Mapping()
	counter.Sort()
	counter.Frequency()
}
