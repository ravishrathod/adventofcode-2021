package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubSystemParser_IncompleteLine(t *testing.T) {
	parser := CreateSubSystemParser("[({(<(())[]>[[{[]{<()<>>")
	parser.Parse()
	assert.False(t, parser.IsCorrupted())
}

func TestSubSystemParser_CorruptLine(t *testing.T) {
	corruptLine(t, "{([(<{}[<>[]}>{[]{[(<()>", "}")
	corruptLine(t, "[[<[([]))<([[{}[[()]]]", ")")
	corruptLine(t, "[{[{({}]{}}([{[{{{}}([]", "]")
	corruptLine(t, "[<(<(<(<{}))><([]([]()", ")")
	corruptLine(t, "<{([([[(<>()){}]>(<<{{", ">")
}

func TestSubSystemParser_IsIncompleteLine(t *testing.T) {
	inCompleteLine(t, "[({(<(())[]>[[{[]{<()<>>", []string{"}", "}", "]", "]", ")", "}", ")", "]"})
	inCompleteLine(t, "[(()[<>])]({[<{<<[]>>(", []string{")", "}", ">", "]", "}", ")"})
	inCompleteLine(t, "(((({<>}<{<{<>}{[]{[]{}", []string{"}", "}", ">", "}", ">", ")", ")", ")", ")"})
	inCompleteLine(t, "{<[[]]>}<{[{[{[]{()[[[]", []string{"]", "]", "}", "}", "]", "}", "]", "}", ">"})
	inCompleteLine(t, "<{([{{}}[<[[[<>{}]]]>[]]", []string{"]", ")", "}", ">"})
}

func inCompleteLine(t *testing.T, line string, completionSymbols []string) {
	parser := CreateSubSystemParser(line)
	parser.Parse()
	assert.True(t, parser.IsIncomplete())
	assert.Equal(t, completionSymbols, parser.CompletionSymbols())
}

func corruptLine(t *testing.T, line string, symbol string) {
	parser := CreateSubSystemParser(line)
	parser.Parse()
	assert.True(t, parser.IsCorrupted())
	assert.Equal(t, symbol, parser.IllegalSymbol())
}

