package commons

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func LinetoInt(line string) []int {
	values := strings.Split(line, ",")
	var days []int
	for _, val := range values {
		day, _ := strconv.Atoi(val)
		days = append(days, day)
	}
	return days
}
