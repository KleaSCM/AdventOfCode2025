/**
 * Test suite for Advent of Code 2025 - Day 1: North Pole Security Dial
 *
 * Tests verify the dial rotation simulation and password counting logic.
 *
 * Author: KleaSCM
 * Email: KleaSCM@gmail.com
 */

package main

import (
	"os"
	"testing"
)

// parse rotation
func TestParseRotation(t *testing.T) {
	tests := []struct {
		input    string
		expected Rotation
		hasError bool
	}{
		{"L68", Rotation{Direction: 'L', Distance: 68}, false},
		{"R48", Rotation{Direction: 'R', Distance: 48}, false},
		{"L1", Rotation{Direction: 'L', Distance: 1}, false},
		{"R100", Rotation{Direction: 'R', Distance: 100}, false},
		{"L", Rotation{}, true},    // Too short
		{"X50", Rotation{}, true},  // Invalid direction
		{"Labc", Rotation{}, true}, // Invalid distance
		{"", Rotation{}, true},     // Empty string
	}

	for _, test := range tests {
		result, err := ParseRotation(test.input)

		if test.hasError {
			if err == nil {
				t.Errorf("ParseRotation(%q) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ParseRotation(%q) unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ParseRotation(%q) = %+v; expected %+v", test.input, result, test.expected)
			}
		}
	}
}

// simulate p1
func TestSimulateDialPart1(t *testing.T) {
	// Example from the problem description
	rotations := []Rotation{
		{'L', 68}, {'L', 30}, {'R', 48}, {'L', 5}, {'R', 60},
		{'L', 55}, {'L', 1}, {'L', 99}, {'R', 14}, {'L', 82},
	}

	result := SimulateDialPart1(rotations)
	expected := 3

	if result != expected {
		t.Errorf("SimulateDialPart1() = %d; expected %d", result, expected)
	}
}

// simulate p2
func TestSimulateDialPart2(t *testing.T) {
	// Example from the problem description
	rotations := []Rotation{
		{'L', 68}, {'L', 30}, {'R', 48}, {'L', 5}, {'R', 60},
		{'L', 55}, {'L', 1}, {'L', 99}, {'R', 14}, {'L', 82},
	}

	result := SimulateDialPart2(rotations)
	expected := 6

	if result != expected {
		t.Errorf("SimulateDialPart2() = %d; expected %d", result, expected)
	}
}

// edge p2
func TestSimulateDialPart2EdgeCase(t *testing.T) {
	// Test case where a single large rotation causes multiple zeros
	rotations := []Rotation{
		{'R', 150}, // Should pass through zero multiple times
	}

	result := SimulateDialPart2(rotations)
	// Starting at 50, rotating R150 passes through 0 at step 50 and step 150
	// So we expect to hit 0 twice during this rotation
	expected := 2

	if result != expected {
		t.Errorf("SimulateDialPart2 edge case = %d; expected %d", result, expected)
	}
}

// dial wrap
func TestDialWrapping(t *testing.T) {
	// Test that position calculations handle wrapping
	rotations := []Rotation{
		{'R', 50}, // 50 + 50 = 0 (should wrap and count)
	}

	result := SimulateDialPart1(rotations)
	expected := 1 // Should end at 0

	if result != expected {
		t.Errorf("Dial wrapping test = %d; expected %d", result, expected)
	}
}

// negative wrap
func TestNegativePositionHandling(t *testing.T) {
	rotations := []Rotation{
		{'L', 60}, // 50 - 60 = -10 -> 90 (mod 100)
	}

	result := SimulateDialPart1(rotations)
	expected := 0 // Should end at 90, not 0

	if result != expected {
		t.Errorf("Negative position handling = %d; expected %d", result, expected)
	}
}

// read rotations
func TestReadRotations(t *testing.T) {
	// Create a temporary file with test data
	content := "L68\nR48\nL5\n"

	tmpFile, err := os.CreateTemp("", "test_rotations_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	rotations, err := ReadRotations(tmpFile.Name())
	if err != nil {
		t.Fatalf("ReadRotations failed: %v", err)
	}

	expected := []Rotation{
		{'L', 68},
		{'R', 48},
		{'L', 5},
	}

	if len(rotations) != len(expected) {
		t.Fatalf("Expected %d rotations, got %d", len(expected), len(rotations))
	}

	for i, exp := range expected {
		if rotations[i] != exp {
			t.Errorf("Rotation %d: expected %+v, got %+v", i, exp, rotations[i])
		}
	}
}

// empty rotations
func TestEmptyRotations(t *testing.T) {
	rotations := []Rotation{}

	part1Result := SimulateDialPart1(rotations)
	part2Result := SimulateDialPart2(rotations)

	if part1Result != 0 {
		t.Errorf("Empty rotations part1 = %d; expected 0", part1Result)
	}

	if part2Result != 0 {
		t.Errorf("Empty rotations part2 = %d; expected 0", part2Result)
	}
}

// zero distance
func TestZeroDistanceRotation(t *testing.T) {
	rotations := []Rotation{
		{'R', 0}, // Should not change position
		{'L', 0}, // Should not change position
	}

	part1Result := SimulateDialPart1(rotations)
	expected := 0 // Should not hit zero

	if part1Result != expected {
		t.Errorf("Zero distance rotations = %d; expected %d", part1Result, expected)
	}
}
