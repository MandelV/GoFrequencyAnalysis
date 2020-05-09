package gofrequencyanalysis

import (
	"io/ioutil"
	"testing"
)

//TestFrequency d
func TestGeneric(test *testing.T) {
	test.Log("Loading file")
	data, err := ioutil.ReadFile("./testing/sample1.txt")
	if err != nil {
		test.Fatal("Unable to open sample file")
	}
	test.Log("Initialize counter")
	counter := Counter{Text: string(data)}
	var freqAnalyse IFreqAnalysis = &counter
	test.Log("Make analysis")
	freqAnalyse.Analysis()

	if len(counter.Pairs) == 0 {
		test.Error("No pairs contained in counter")
	}

}
