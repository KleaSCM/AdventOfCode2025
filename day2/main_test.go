/**
 * Test suite for Advent of Code 2025 - Day 2: Invalid Product IDs
 *
 * Tests verify the ID validation logic and range processing.
 *
 * Author: KleaSCM
 * Email: KleaSCM@gmail.com
 */

package main

import (
	"os"
	"testing"
)

// parse range
func TestParseIDRange(t *testing.T) {
	tests := []struct {
		input    string
		expected IDRange
		hasError bool
	}{
		{"11-22", IDRange{Start: 11, End: 22}, false},
		{"95-115", IDRange{Start: 95, End: 115}, false},
		{"1-1", IDRange{Start: 1, End: 1}, false},
		{"100-99", IDRange{}, true},   // start > end
		{"11", IDRange{}, true},       // missing dash
		{"11-22-33", IDRange{}, true}, // too many parts
		{"abc-123", IDRange{}, true},  // invalid numbers
		{"", IDRange{}, true},         // empty string
	}

	for _, test := range tests {
		result, err := ParseIDRange(test.input)

		if test.hasError {
			if err == nil {
				t.Errorf("ParseIDRange(%q) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ParseIDRange(%q) unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ParseIDRange(%q) = %+v; expected %+v", test.input, result, test.expected)
			}
		}
	}
}

// parse ranges
func TestParseIDRanges(t *testing.T) {
	tests := []struct {
		input    string
		expected []IDRange
		hasError bool
	}{
		{"11-22,95-115", []IDRange{{11, 22}, {95, 115}}, false},
		{"11-22", []IDRange{{11, 22}}, false},
		{"", nil, false},                       // empty line
		{"11-22,", []IDRange{{11, 22}}, false}, // trailing comma
		{",11-22", []IDRange{{11, 22}}, false}, // leading comma
		{"11-22,invalid", nil, true},           // invalid range
	}

	for _, test := range tests {
		result, err := ParseIDRanges(test.input)

		if test.hasError {
			if err == nil {
				t.Errorf("ParseIDRanges(%q) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ParseIDRanges(%q) unexpected error: %v", test.input, err)
			}
			if len(result) != len(test.expected) {
				t.Errorf("ParseIDRanges(%q) length = %d; expected %d", test.input, len(result), len(test.expected))
				continue
			}
			for i, expected := range test.expected {
				if result[i] != expected {
					t.Errorf("ParseIDRanges(%q)[%d] = %+v; expected %+v", test.input, i, result[i], expected)
				}
			}
		}
	}
}

// invalid p1
func TestIsInvalidIDPart1(t *testing.T) {
	tests := []struct {
		id       int
		expected bool
	}{
		{11, true},    // "11" = "1" repeated twice
		{22, true},    // "22" = "2" repeated twice
		{99, true},    // "99" = "9" repeated twice
		{1010, true},  // "1010" = "10" repeated twice
		{123, false},  // odd length
		{1234, false}, // "12" != "34"
		{111, false},  // not exactly two repetitions
		{1, false},    // too short
		{0, false},    // zero
	}

	for _, test := range tests {
		result := IsInvalidIDPart1(test.id)
		if result != test.expected {
			t.Errorf("IsInvalidIDPart1(%d) = %v; expected %v", test.id, result, test.expected)
		}
	}
}

// invalid p2
func TestIsInvalidIDPart2(t *testing.T) {
	tests := []struct {
		id       int
		expected bool
	}{
		{11, true},        // "11" = "1" repeated twice
		{111, true},       // "111" = "1" repeated three times
		{123123, true},    // "123123" = "123" repeated twice
		{123123123, true}, // "123123123" = "123" repeated three times
		{121212, true},    // "121212" = "12" repeated three times
		{99, true},        // "99" = "9" repeated twice
		{123, false},      // not repeated
		{1234, false},     // not repeated pattern
		{1, false},        // too short
		{0, false},        // zero
	}

	for _, test := range tests {
		result := IsInvalidIDPart2(test.id)
		if result != test.expected {
			t.Errorf("IsInvalidIDPart2(%d) = %v; expected %v", test.id, result, test.expected)
		}
	}
}

// sum invalid
func TestSumInvalidIDsInRanges(t *testing.T) {
	ranges := []IDRange{
		{10, 15}, // contains 11
		{20, 25}, // contains 22
	}

	// p1 only 11 and 22 are invalid
	part1Sum := SumInvalidIDsInRanges(ranges, IsInvalidIDPart1)
	expectedPart1 := int64(11 + 22)
	if part1Sum != expectedPart1 {
		t.Errorf("SumInvalidIDsInRanges (part 1) = %d; expected %d", part1Sum, expectedPart1)
	}

	// p2 11 and 22 are still invalid
	part2Sum := SumInvalidIDsInRanges(ranges, IsInvalidIDPart2)
	expectedPart2 := int64(11 + 22)
	if part2Sum != expectedPart2 {
		t.Errorf("SumInvalidIDsInRanges (part 2) = %d; expected %d", part2Sum, expectedPart2)
	}
}

// empty ranges
func TestEmptyRanges(t *testing.T) {
	ranges := []IDRange{}

	part1Sum := SumInvalidIDsInRanges(ranges, IsInvalidIDPart1)
	part2Sum := SumInvalidIDsInRanges(ranges, IsInvalidIDPart2)

	if part1Sum != 0 {
		t.Errorf("Empty ranges part1 sum = %d; expected 0", part1Sum)
	}

	if part2Sum != 0 {
		t.Errorf("Empty ranges part2 sum = %d; expected 0", part2Sum)
	}
}

// single range
func TestSingleNumberRange(t *testing.T) {
	ranges := []IDRange{
		{11, 11}, // single number that is invalid
		{12, 12}, // single number that is valid
	}

	part1Sum := SumInvalidIDsInRanges(ranges, IsInvalidIDPart1)
	expected := int64(11) // only 11 is invalid
	if part1Sum != expected {
		t.Errorf("Single number range sum = %d; expected %d", part1Sum, expected)
	}
}

// large numbers
func TestLargeNumbers(t *testing.T) {
	// Test with a large number from the example
	id := 1188511885
	if !IsInvalidIDPart1(id) {
		t.Errorf("Large ID %d should be invalid for part 1", id)
	}

	if !IsInvalidIDPart2(id) {
		t.Errorf("Large ID %d should be invalid for part 2", id)
	}
}

// read ranges
func TestReadAndProcessRanges(t *testing.T) {
	// Create a temporary file with test data
	content := "11-22,95-115\n998-1012\n"

	tmpFile, err := os.CreateTemp("", "test_ranges_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	ranges, err := ReadAndProcessRanges(tmpFile.Name())
	if err != nil {
		t.Fatalf("ReadAndProcessRanges failed: %v", err)
	}

	expected := []IDRange{
		{11, 22}, {95, 115},
		{998, 1012},
	}

	if len(ranges) != len(expected) {
		t.Fatalf("Expected %d ranges, got %d", len(expected), len(ranges))
	}

	for i, exp := range expected {
		if ranges[i] != exp {
			t.Errorf("Range %d: expected %+v, got %+v", i, exp, ranges[i])
		}
	}
}
