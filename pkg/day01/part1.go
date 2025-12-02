package day01

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput() ([]int, error) {
	content, err := os.ReadFile("pkg/day01/input.txt")
	if err != nil {
		return nil, fmt.Errorf("reading input:txt: %w", err)
	}

	repl := strings.NewReplacer("L", "-", "R", "")

	normalized := repl.Replace(strings.TrimSpace(string(content)))
	lines := strings.Split(normalized, "\n")

	rotations := make([]int, len(lines))

	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("parsing %s to int: %w", line, err)
		}

		rotations[i] = n
	}

	return rotations, nil
}

func Part1() {
	current := 50
	rotations, err := parseInput()
	if err != nil {
		log.Fatalf("Day 1 Part 1: grabbing inputs: %w", err)
	}

	blorb := 0

	for _, rotation := range rotations {
		current += rotation
		if current%100 == 0 {
			blorb++
		}
	}

	log.Printf("Day 1 Part 1: Thingy stops on 0 '%d' times.", blorb)
}
