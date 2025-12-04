package day01

import (
	"fmt"
	"log"
)

const (
	positive int = iota
	negative
)

func Part2() {
	rotations, err := parseInput()
	if err != nil {
		log.Fatalf("Day 1 Part 1: grabbing inputs: %s", err)
	}

	blorb := calculateCoversZero(rotations)

	log.Printf("Day 1 Part 2: Thingy touches on 0 '%d' times.", blorb)
}

func calculateCoversZero(rotations []int) int {
	current := 50
	blorb := 0

	//direction := positive
	//collector := 0

	//chunks := make([]int, 0)

	//fmt.Printf("collector is 0, direction is positive\n")

	//for _, rotation := range rotations {
	//	// rotation negative, direction negative
	//	if rotation < 0 && direction == negative {
	//		collector += rotation
	//		//fmt.Printf("rotation is negative %d, direction is negative, collector is %d\n", rotation, collector)
	//		continue
	//	}
	//
	//	// rotation positive, direction negative, change
	//	if rotation > 0 && direction == negative {
	//		chunks = append(chunks, collector)
	//		//fmt.Printf("rotation is positive %d, direction is negative, collector is %d, adding it to chunk\n", rotation, collector)
	//		collector = rotation
	//		direction = positive
	//		//fmt.Printf(" L--- collector reset to %d, direction reset to positive\n", rotation)
	//		continue
	//	}
	//
	//	// rotation negative, direction positive, change
	//	if rotation < 0 && direction == positive {
	//		chunks = append(chunks, collector)
	//		//fmt.Printf("rotation is negative %d, direction is positive, collector is %d, adding it to chunk\n", rotation, collector)
	//
	//		collector = rotation
	//		direction = negative
	//		//fmt.Printf(" L--- collector reset to %d, direction reset to negative\n", rotation)
	//		continue
	//	}
	//
	//	// rotation positive, direction positive
	//	if rotation > 0 && direction == positive {
	//		collector += rotation
	//		//fmt.Printf("rotation is positive %d, direction is positive, collector is %d\n", rotation, collector)
	//		continue
	//	}
	//}
	//chunks = append(chunks, collector)

	for i, rotation := range rotations {
		fmt.Printf("\n%3d: current: %d, rotation: %d => %d\n", i, current, rotation, current+rotation)
		newBlorb, newCurrent := coversZero(current, rotation)
		blorb += newBlorb
		//fmt.Printf("current: %d, rotation: %d, newBlorb: %d, allblorbs: %d\n\n", current, rotation, newBlorb, blorb)
		current = newCurrent
	}
	//
	//fmt.Printf("this is the chunked: %v\n"+
	//	"len of original: %d\n"+
	//	"len of chunked: %d\n", chunks, len(rotations), len(chunks))

	return blorb
}

// coversZero calculates how many times a rotation touched the number 0. Return arguments are
// turns, newCurrent
func coversZero(current, rotation int) (int, int) {
	if rotation == 0 {
		return 0, current
	}

	switch rotation < 0 {
	case true:
		return coversZeroNegative(current, rotation)
	case false:
		return coversZeroPositive(current, rotation)
	default:
		log.Fatalf("switch rotation < 0 should never have tripped this")
		return 0, 0
	}
}

// coversZeroNegative calculates the zero touching if we're rotating left. Return values are
// turns, newCurrent
func coversZeroNegative(current, rotation int) (int, int) {
	fmt.Printf("we're doing negative turns...")
	fullTurns := abs(rotation / 100)
	remainder := rotation % 100
	remainderFromCurrent := current % 100
	current += rotation

	if abs(remainder) < abs(remainderFromCurrent) {
		fmt.Printf("there were %d full turns, and no extra, because remainder of rotation %d is less than remainder of current %d\n", fullTurns, remainder, remainderFromCurrent)
		return fullTurns, current
	}

	if remainderFromCurrent == 0 {
		fmt.Printf("there were %d full turns, and 0 extra, because rotation started from a hundred position, so we only count full turns\n", fullTurns)
		return fullTurns, current
	}

	fmt.Printf("there were %d full turns, and 1 extra, because remainder of rotation %d is more than or equal to remainder of current %d\n", fullTurns, remainder, remainderFromCurrent)

	return fullTurns + 1, current
}

// coversZeroPositive calculates the zero touching if we're rotating right. Return values are
// turns, newCurrent
func coversZeroPositive(current, rotation int) (int, int) {
	fmt.Printf("we're doing positive turns...")
	fullTurns := abs(rotation / 100)
	remainder := rotation % 100
	remainderFromCurrent := 100 - (current % 100)
	current += rotation

	if abs(remainder) < abs(remainderFromCurrent) {
		fmt.Printf("there were %d full turns, and no extra, because remainder of rotation %d is less than remainder of current %d\n", fullTurns, remainder, remainderFromCurrent)
		return fullTurns, current
	}

	fmt.Printf("there were %d full turns, and 1 extra, because remainder of rotation %d is more than or equal to remainder of current %d\n", fullTurns, remainder, remainderFromCurrent)

	return fullTurns + 1, current
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
