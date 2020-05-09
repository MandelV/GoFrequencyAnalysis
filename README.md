# GoFrequencyAnalysis
![Go](https://github.com/MandelV/GoFrequencyAnalysis/workflows/Go/badge.svg?branch=master)

A little package developed with Go to make frequency analysis on text


# Data

First of all there is two structures Counter and Pair

#### Counter
Counter is a main structure, it contains the results and data about the text to be analysed.
```Go
type Counter struct {
	Text       string
	SizeText   int
	MappedKeys map[string]int
	Pairs      []Pair
}
```
The Text and its size are contained in it, but most importantly is MappedKeys and Pairs :

MappedKeys : will contain for each character the number of occurence in the text
Pairs : see next section.

#### Pairs

The structure "Pair" is useful because it enables to match a character with how many time it occur and frequency of its apparition.
```Go
type Pair struct {
	Char  string
	Count int
	Freq  float32
}
```

# How to use

So, to do some analysis you just need two things : a text and a counter which also contains your text.
Here an example :

```Go
text := "beautiful text"

counter := Counter{Text: text} //Not needed to set the size of text

var freqAnalysis IFreqAnalysis = &counter

freqAnalysis.Analysis()

for _, pair := range counter.Pairs {
	test.Log(pair.Char, pair.Count, pair.Freq)
}
/*
 *Output:
 * t 3 21.428572
 * e 2 14.285715
 * u 2 14.285715
 * l 1 7.1428576
 *   1 7.1428576
 * b 1 7.1428576
 * a 1 7.1428576
 * i 1 7.1428576
 * f 1 7.1428576
 * x 1 7.1428576
 */
```

# Method
#### Mapping()
Its a main method it map each character with its number of occurence 
#### Sort()
Ensure that the most occur characters will be in the firt indexes in the Pairs array
#### Frequency()
Calculate the frequence for each character
#### Analysis()
Its a helper that execute three previous method see above (same order)