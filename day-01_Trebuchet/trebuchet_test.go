package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputExample(t *testing.T) {
	buffer, err := os.ReadFile("input_example")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	result := Trebuchet(string(buffer))
	assert.Equal(t, 142, result)
}

func TestInputExamplePart2(t *testing.T) {
	buffer, err := os.ReadFile("input_part2")
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	result := Trebuchet(string(buffer))
	assert.Equal(t, 281, result)
}
