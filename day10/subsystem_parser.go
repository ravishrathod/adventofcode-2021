package main

import "adventoccode2021/commons"

type SubSystemParser struct {
	input                     string
	stack                     *commons.Stack
	openingSymbols            map[string]bool
	closingSymbols            map[string]bool
	openingToClosingSymbolMap map[string]string
	corrupted                 bool
	incomplete                bool
	illegalSymbol             string
	completionSymbols         []string
}

func CreateSubSystemParser(input string) *SubSystemParser {
	parser := &SubSystemParser{
		stack: &commons.Stack{},
		input: input,
	}
	parser.openingToClosingSymbolMap = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
	parser.openingSymbols = make(map[string]bool)
	parser.closingSymbols = make(map[string]bool)

	for k, v := range parser.openingToClosingSymbolMap {
		parser.openingSymbols[k] = true
		parser.closingSymbols[v] = true
	}
	return parser
}

func (this *SubSystemParser) Parse() {
	chars := []rune(this.input)
	for _, char := range chars {
		symbol := string(char)
		if this.openingSymbols[symbol] == true {
			this.stack.Push(symbol)
		} else if this.closingSymbols[symbol] == true {
			openingSymbol, err := this.stack.Peek()
			if err != nil {
				this.corrupted = true
				this.illegalSymbol = symbol
				return
			}
			if this.openingToClosingSymbolMap[openingSymbol] != symbol {
				this.corrupted = true
				this.illegalSymbol = symbol
				return
			}
			_, err = this.stack.Pop()
			if err != nil {
				panic(err)
			}
		}
	}
	this.incomplete = !this.stack.IsEmpty()
	if this.incomplete {
		this.computeCompletionSymbols()
	}
}

func (this *SubSystemParser) computeCompletionSymbols() {
	for unbalancedSymbol, err := this.stack.Pop(); err == nil; unbalancedSymbol, err = this.stack.Pop() {
		this.completionSymbols = append(this.completionSymbols, this.openingToClosingSymbolMap[unbalancedSymbol])
	}
}

func (this *SubSystemParser) IsCorrupted() bool {
	return this.corrupted
}

func (this *SubSystemParser) IsIncomplete() bool {
	return this.incomplete
}

func (this *SubSystemParser) CompletionSymbols() []string {
	return this.completionSymbols
}

func (this *SubSystemParser) IllegalSymbol() string {
	return this.illegalSymbol
}
