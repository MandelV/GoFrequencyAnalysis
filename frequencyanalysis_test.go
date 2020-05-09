package gofrequencyanalysis

import (
	"io/ioutil"
	"testing"
)

func initTests(test *testing.T) *Counter {
	test.Log("Loading file...")
	data, err := ioutil.ReadFile("./testing/sample1.txt")
	if err != nil {
		test.Fatal(err)
	}
	test.Log("Initialize Counter")
	counter := Counter{Text: string(data)}
	return &counter
}

//TestFrequency d
func TestNumberOfPairs(test *testing.T) {

	counter := initTests(test)
	var freqAnalyse IFreqAnalysis = counter
	test.Log("Make analysis")
	freqAnalyse.Analysis()

	if len(counter.Pairs) == 0 {
		test.Error("No pairs contained in counter")
	}
}
