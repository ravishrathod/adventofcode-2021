package main

import (
	"adventoccode2021/commons"
	"strconv"
	"strings"
)

func main() {

	lines, err := commons.ReadFile("input/day2.txt")
	if err != nil {
		panic(err)
	}
	commands, _ := parseToCommands(lines)
	part1(commands)
	part2(commands)
}

func part1(commands []command) {
	var x int
	var depth int
	
	for _, command :=  range commands {
		switch command.Direction {
		case "forward":
			x = x + command.Amplitude
		case "up":
			depth =  depth - command.Amplitude
		case "down":
			depth =  depth + command.Amplitude
		}
	}

	println(x*depth)
}

func part2(commands []command) {
	var x int
	var depth int
	var aim int

	for _, command := range commands {
		switch command.Direction {
		case "forward":
			x = x + command.Amplitude
			depth =  depth + aim * command.Amplitude
		case "up":
			aim = aim - command.Amplitude
		case "down":
			aim = aim + command.Amplitude
		}
	}

	println(x*depth)
}

func parseToCommands(lines []string) ([]command, error) {
	var commands []command

	for _, input := range lines {
		parts := strings.Split(input, " ")
		amplitude, _ := strconv.Atoi(parts[1])
		command := &command{
			Direction: parts[0],
			Amplitude: amplitude,
		}
		commands = append(commands, *command)
	}
	return commands, nil
}

type command struct {
	Direction string
	Amplitude int
}
