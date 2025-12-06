/**
 * Test suite for Advent of Code 2025 - Day 4: Printing Department
 *
 * Tests verify the forklift accessibility logic and grid operations.
 *
 * Author: KleaSCM
 * Email: KleaSCM@gmail.com
 */

package main

import (
	"os"
	"testing"
)

// adjacent
func TestCountAdjacentRolls(t *testing.T) {
	// Test grid from the problem description
	grid := Grid{
		[]rune("..@@.@@@@."),
		[]rune("@@@.@.@.@@"),
		[]rune("@@@@@.@.@@"),
		[]rune("@.@@@@..@."),
		[]rune("@@.@@@@.@@"),
		[]rune(".@@@@@@@.@"),
		[]rune(".@.@.@.@@@"),
		[]rune("@.@@@.@@@@"),
		[]rune(".@@@@@@@@."),
		[]rune("@.@.@@@.@."),
	}

	tests := []struct {
		row, col int
		expected int
	}{
		{0, 0, 2}, // '.' position with 2 adjacent '@'
		{0, 2, 3}, // '@' position with 3 adjacent '@'
		{4, 4, 8}, // center position with 8 adjacent '@'
		{9, 9, 2}, // corner with 2 adjacent '@'
	}

	for _, test := range tests {
		result := CountAdjacentRolls(grid, test.row, test.col)
		if result != test.expected {
			t.Errorf("CountAdjacentRolls(%d, %d) = %d; expected %d", test.row, test.col, result, test.expected)
		}
	}
}

// accessible
func TestIsAccessible(t *testing.T) {
	grid := Grid{
		[]rune("..@@.@@@@."),
		[]rune("@@@.@.@.@@"),
		[]rune("@@@@@.@.@@"),
		[]rune("@.@@@@..@."),
		[]rune("@@.@@@@.@@"),
		[]rune(".@@@@@@@.@"),
		[]rune(".@.@.@.@@@"),
		[]rune("@.@@@.@@@@"),
		[]rune(".@@@@@@@@."),
		[]rune("@.@.@@@.@."),
	}

	tests := []struct {
		row, col int
		expected bool
	}{
		{0, 2, true},  // '@' with 3 adjacent (< 4, accessible)
		{0, 3, true},  // '@' with 3 adjacent (< 4, accessible)
		{1, 0, true},  // '@' with 3 adjacent (< 4, accessible)
		{9, 1, false}, // '.' is not accessible
		{0, 0, false}, // '.' is not accessible
	}

	for _, test := range tests {
		result := IsAccessible(grid, test.row, test.col)
		if result != test.expected {
			t.Errorf("IsAccessible(%d, %d) = %v; expected %v", test.row, test.col, result, test.expected)
		}
	}
}

// count accessible
func TestCountAccessibleRolls(t *testing.T) {
	// Test grid from the problem description
	grid := Grid{
		[]rune("..@@.@@@@."),
		[]rune("@@@.@.@.@@"),
		[]rune("@@@@@.@.@@"),
		[]rune("@.@@@@..@."),
		[]rune("@@.@@@@.@@"),
		[]rune(".@@@@@@@.@"),
		[]rune(".@.@.@.@@@"),
		[]rune("@.@@@.@@@@"),
		[]rune(".@@@@@@@@."),
		[]rune("@.@.@@@.@."),
	}

	result := CountAccessibleRolls(grid)
	expected := 13

	if result != expected {
		t.Errorf("CountAccessibleRolls() = %d; expected %d", result, expected)
	}
}

// removal p2
func TestCountTotalRemovableRolls(t *testing.T) {
	// Test grid from the problem description
	grid := Grid{
		[]rune("..@@.@@@@."),
		[]rune("@@@.@.@.@@"),
		[]rune("@@@@@.@.@@"),
		[]rune("@.@@@@..@."),
		[]rune("@@.@@@@.@@"),
		[]rune(".@@@@@@@.@"),
		[]rune(".@.@.@.@@@"),
		[]rune("@.@@@.@@@@"),
		[]rune(".@@@@@@@@."),
		[]rune("@.@.@@@.@."),
	}

	result := CountTotalRemovableRolls(grid)
	expected := 43 // From the example in the problem

	if result != expected {
		t.Errorf("CountTotalRemovableRolls() = %d; expected %d", result, expected)
	}
}

// small grid
func TestCountTotalRemovableRollsSmall(t *testing.T) {
	// Simple test case
	grid := Grid{
		[]rune("..."),
		[]rune(".@."),
		[]rune("..."),
	}

	result := CountTotalRemovableRolls(grid)
	expected := 1 // Single '@' with 0 adjacent rolls (definitely < 4)

	if result != expected {
		t.Errorf("CountTotalRemovableRolls small test = %d; expected %d", result, expected)
	}
}

// full removal
func TestCountTotalRemovableRollsFullRemoval(t *testing.T) {
	// 3x3 grid of all '@' - corners and edges are initially accessible (< 4 adjacent)
	// After removing them, center becomes accessible (0 adjacent)
	grid := Grid{
		[]rune("@@@"),
		[]rune("@@@"),
		[]rune("@@@"),
	}

	result := CountTotalRemovableRolls(grid)
	expected := 9 // All 9 rolls get removed over multiple iterations

	if result != expected {
		t.Errorf("CountTotalRemovableRolls full removal = %d; expected %d", result, expected)
	}
}

// read input
func TestReadInput(t *testing.T) {
	// Create a temporary file with test data
	content := "..@@.\n@@@.@\n"

	tmpFile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	grid, err := ReadInput(tmpFile.Name())
	if err != nil {
		t.Fatalf("ReadInput failed: %v", err)
	}

	expected := Grid{
		[]rune("..@@."),
		[]rune("@@@.@"),
	}

	if len(grid) != len(expected) {
		t.Fatalf("Expected %d rows, got %d", len(expected), len(grid))
	}

	for i, exp := range expected {
		if len(grid[i]) != len(exp) {
			t.Errorf("Row %d length: expected %d, got %d", i, len(exp), len(grid[i]))
			continue
		}
		for j, val := range exp {
			if grid[i][j] != val {
				t.Errorf("Grid[%d][%d]: expected %c, got %c", i, j, val, grid[i][j])
			}
		}
	}
}
