package day01

import (
	"fmt"
	"log"
)

const (
	modulus = 100
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

	for i, rotation := range rotations {
		fmt.Printf("\n%3d: current: %d, rotation: %d\n", i, current, rotation)
		newBlorb, newCurrent := coversZero(current, rotation)

		blorb += newBlorb
		fmt.Printf("-- touches: %2d, total: %4d\n\n", newBlorb, blorb)

		current = newCurrent
	}

	return blorb
}

// coversZero calculates how many times a rotation touched the number 0. Return arguments are
// turns, newCurrent
func coversZero(current, rotation int) (int, int) {
	if rotation == 0 {
		return 0, current
	}

	negative := rotation < 0

	fulls := abs(rotation / modulus)
	remainder := abs(rotation % modulus)
	lower, upper := limits(current)
	if current != 0 {
		if negative && remainder >= lower {
			fulls++
		}

		if !negative && remainder >= upper {
			fulls++
		}
	}

	current = moduloShiftKeepInRange(current, rotation)

	return fulls, current
}

// moduloShiftKeepInRange will wrap around the offset change from initial between
// 0 and modulus. The parts:
// initial: this is our starting value, where the offset is going to happen from
// (offset % modulus): we only care about the change that's less than the modulus, so full turns don't count
// + modulus: we want to deal with positive values only. (offset % modulus) can produce negative, for example -255 % 100 == -55. This one takes it back up to 45
// % modulus at the end: and in case the offset was positive, the previous + modulus would have taken us over the modulus, so this one restricts us back to be within range
func moduloShiftKeepInRange(initial, offset int) int {
	return ((initial + (offset % modulus)) + modulus) % modulus
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func limits(current int) (int, int) {
	return current, modulus - current
}
