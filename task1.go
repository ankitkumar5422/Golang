package main

import (
	"fmt"
	"sort"
)

var Taxonomy_CPI_LIMIT = map[string]float64{
	"207L00000X": 2.5,
	"207N00000X": 2.5,
	"207P00000X": 2.9,
	"207R00000X": 1.6,
	"207RC0000X": 2.5,
	"207RC0200X": 3,
	"207RE0101X": 2.5,
	"207RG0100X": 3.7,
	"207RH0000X": 4,
	"207RH0003X": 4,
	"207RN0300X": 2.4,
	"207RP1001X": 2.8,
	"207RR0500X": 3.5,
	"207RX0202X": 4,
	"207T00000X": 3.5,
	"207U00000X": 3.4,
	"207V00000X": 2.2,
	"207W00000X": 2.6,
	"207X00000X": 2.7,
	"207Y00000X": 2.5,
	"208000000X": 1.8,
	"2084A0401X": 1.2,
	"2085B0100X": 4.2,
	"208600000X": 2,
	"208800000X": 3.2,
	"208D00000X": 1.0,
	"301D00000X": 1.4,
	"301D00001X": 1.4,
	"301D00002X": 2.1,
	"301D00004X": 1.5,
	"301D00005X": 1.4,
	"301D00007X": 4.1,
	"301D00008X": 4.2,
	"301D00009X": 2.5,
	"301D00010X": 4.1,
}

// sorting the string using slices
// create a function
func taxonomyCPI(taxonomyCPI map[string]float64, input []string) []string {

	sortedKeys := make([]string, len(input))
	copy(sortedKeys, input)

	sort.Slice(sortedKeys, func(i, j int) bool {
		return taxonomyCPI[sortedKeys[i]] < taxonomyCPI[sortedKeys[j]]
	})

	return sortedKeys
}

func main() {
	input := []string{"207RC0200X", "208800000X", "208600000X"}
	sorted := taxonomyCPI(Taxonomy_CPI_LIMIT, input)
	fmt.Println(sorted)
}
