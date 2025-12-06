/**
 * Advent of Code 2025 - Day 1: North Pole Security Dial
 *
 * This program solves the dial password problem by simulating
 * rotations and counting how many times the dial points at zero.
 *
 * Author: KleaSCM
 * Email: KleaSCM@gmail.com
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Rotation represents a single dial rotation instruction
type Rotation struct {
	Direction byte // 'L' for left, 'R' for right
	Distance  int  // Number of clicks to rotate
}

// reads and parses rotation instructions from file
func ReadRotations(filename string) ([]Rotation, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rotations []Rotation
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		rotation, err := ParseRotation(line)
		if err != nil {
			return nil, err
		}
		rotations = append(rotations, rotation)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rotations, nil
}

// converts line like "L68" or "R48" to a struct
func ParseRotation(line string) (Rotation, error) {
	if len(line) < 2 {
		return Rotation{}, fmt.Errorf("invalid rotation format: %s", line)
	}

	direction := line[0]
	if direction != 'L' && direction != 'R' {
		return Rotation{}, fmt.Errorf("invalid direction: %c", direction)
	}

	distance, err := strconv.Atoi(line[1:])
	if err != nil {
		return Rotation{}, fmt.Errorf("invalid distance: %s", line[1:])
	}

	return Rotation{Direction: direction, Distance: distance}, nil
}

// counts how many times the dial ends at position 0
// after each complete rotation (p1)
func SimulateDialPart1(rotations []Rotation) int {
	position := 50
	zeroCount := 0

	for _, rotation := range rotations {
		if rotation.Direction == 'L' {
			position = (position - rotation.Distance) % 100
			if position < 0 {
				position += 100
			}
		} else if rotation.Direction == 'R' {
			position = (position + rotation.Distance) % 100
		}

		if position == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

// counts every time the dial points at 0 during any rotation,
// including intermediate positions (p2)
func SimulateDialPart2(rotations []Rotation) int {
	position := 50
	zeroCount := 0

	for _, rotation := range rotations {
		if rotation.Direction == 'L' {
			// Moving left (decreasing)
			for step := 1; step <= rotation.Distance; step++ {
				position = (position - 1) % 100
				if position < 0 {
					position += 100
				}
				if position == 0 {
					zeroCount++
				}
			}
		} else if rotation.Direction == 'R' {
			// Moving right (increasing)
			for step := 1; step <= rotation.Distance; step++ {
				position = (position + 1) % 100
				if position == 0 {
					zeroCount++
				}
			}
		}
	}

	return zeroCount
}

func main() {
	// p1: count zeros at end of rotations
	rotations, err := ReadRotations("input/input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	part1Result := SimulateDialPart1(rotations)
	fmt.Printf("Password (p1): %d\n", part1Result)

	// p2: count all zeros during rotations
	part2Result := SimulateDialPart2(rotations)
	fmt.Printf("Password (p2): %d\n", part2Result)
}
