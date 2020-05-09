package gofrequencyanalysis

/*
 * This package enable to calculate the occurence
 * of each character contained in given Text.
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
	Char  string
	Count int
	Freq  float32
}

/*Counter : Contain all data around Text like :
 * The original Text
 * the sizeof the Text
 * mappedKeys which is the number of time a key appear in the Text
 * array of Pair (see Pair)
 */
type Counter struct {
	Text       string
	SizeText   int
	MappedKeys map[string]int
	Pairs      []Pair
}

//Sort : Enable to sorting the Pairs
func (Counter *Counter) Sort() {
	var Pairs []Pair
	//Initialization of Pairs
	for k, v := range Counter.MappedKeys {
		Pairs = append(Pairs, Pair{k, v, 0})
	}
	//Sorting
	sort.Slice(Pairs, func(i, j int) bool {
		return Pairs[i].Count > Pairs[j].Count
	})
	Counter.Pairs = Pairs
}

//Mapping : Count each character in Text
func (Counter *Counter) Mapping() {

	Counter.SizeText = len(Counter.Text)

	CountedKeys := make(map[string]int)

	for i := 0; i < Counter.SizeText; i++ {
		key := Counter.Text[i : i+1]
		CountedKeys[key]++
	}
	Counter.MappedKeys = CountedKeys
}

//Frequency : Calculate the occurrence frequency
func (Counter *Counter) Frequency() {

	for i := 0; i < len(Counter.Pairs); i++ {

		Counter.Pairs[i].Freq = float32(Counter.Pairs[i].Count) / float32(Counter.SizeText) * float32(100)
	}
}

//Analysis : Automated the process to mapping, sorting and finally calcule frequency
func (Counter *Counter) Analysis() {
	Counter.Mapping()
	Counter.Sort()
	Counter.Frequency()
}
