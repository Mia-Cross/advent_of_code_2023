package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func lineResult(line string) string {
	result := ""
	for _, c := range line {
		if unicode.IsDigit(c) {
			result += string(c)
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			result += string(line[i])
			break
		}
	}
	return result
}

func Trebuchet(str string) int {
	input := strings.Split(str, "\n")
	result := 0
	for _, line := range input {
		x, err := strconv.Atoi(lineResult(line))
		if err != nil {
			return -1
		}
		result += x
	}
	return result
}

func main() {
	buffer, err := os.ReadFile("input")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	result := Trebuchet(string(buffer))

	fmt.Printf("Result: %d", result)
}
