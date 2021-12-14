package main

import (
	"fmt"
	"testing"
)


func TestPolymer_ApplyRules(t *testing.T) {
	rules := map[string]string {
		"CH": "B",
		"HH" : "N",
		"CB" : "H",
		"NH" : "C",
		"HB" : "C",
		"HC" : "B",
		"HN" : "C",
		"NN" : "C",
		"BH" : "H",
		"NC" : "B",
		"NB" : "B",
		"BN" : "B",
		"BB" : "N",
		"BC" : "B",
		"CC" : "N",
		"CN" : "C",
	}

	polymer := CreatePolymer("NNCB")

	polymer.ApplyRules(rules)
	fmt.Printf("%v", polymer.pairCount)
}
