package main

import (
	"adventoccode2021/commons"
	"strings"
)

func main() {
	lines, err := commons.ReadFile("input/day14.txt")
	if err != nil {
		panic(err)
	}
	template := lines[0]
	rules := make(map[string]string)

	for _, line := range lines {
		if strings.Contains(line, " -> ") {
			entries := strings.Split(line, " -> ")
			rules[entries[0]] = entries[1]
		}
	}

	polymer := CreatePolymer(template)

	for i:=0;i<40;i++ {
		polymer.ApplyRules(rules)
	}
	max, min := polymer.GetMaxAndMinCounts()
	println(max-min)
}
