/**
 * Test suite for Advent of Code 2025 - Day 5: Cafeteria Inventory
 *
 * Tests verify the ingredient freshness checking and range processing.
 *
 * Author: KleaSCM
 * Email: KleaSCM@gmail.com
 */

package main

import (
	"os"
	"testing"
)

// range parsing function
func TestParseIDRange(t *testing.T) {
	tests := []struct {
		input    string
		expected IDRange
		hasError bool
	}{
		{"3-5", IDRange{Start: 3, End: 5}, false},
		{"10-14", IDRange{Start: 10, End: 14}, false},
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

// freshness checking function
func TestIsFresh(t *testing.T) {
	ranges := []IDRange{
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18}, // overlapping
	}

	tests := []struct {
		id       int64
		expected bool
	}{
		{1, false},  // not in any range
		{3, true},   // in range 3-5
		{5, true},   // in range 3-5
		{8, false},  // not in any range
		{10, true},  // in range 10-14
		{12, true},  // in overlapping ranges 10-14 and 12-18
		{15, true},  // in range 12-18
		{16, true},  // in range 16-20
		{20, true},  // in range 16-20
		{21, false}, // not in any range
	}

	for _, test := range tests {
		result := IsFresh(test.id, ranges)
		if result != test.expected {
			t.Errorf("IsFresh(%d, ranges) = %v; expected %v", test.id, result, test.expected)
		}
	}
}

// tests the complete ingredient counting
func TestCountFreshIngredients(t *testing.T) {
	// Create a temporary file with test data matching the problem example
	content := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	tmpFile, err := os.CreateTemp("", "test_ingredients_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	result, err := CountFreshIngredients(tmpFile.Name())
	if err != nil {
		t.Fatalf("CountFreshIngredients failed: %v", err)
	}

	expected := 3 // IDs 5, 11, 17 are fresh according to the example
	if result != expected {
		t.Errorf("CountFreshIngredients() = %d; expected %d", result, expected)
	}
}

// tests with empty range list
func TestCountFreshIngredientsEmptyRanges(t *testing.T) {
	content := `
1
5
8`

	tmpFile, err := os.CreateTemp("", "test_empty_ranges_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	result, err := CountFreshIngredients(tmpFile.Name())
	if err != nil {
		t.Fatalf("CountFreshIngredients failed: %v", err)
	}

	expected := 0 // no ranges defined, so no IDs are fresh
	if result != expected {
		t.Errorf("CountFreshIngredients empty ranges = %d; expected %d", result, expected)
	}
}

// tests with ranges but no available IDs
func TestCountFreshIngredientsNoIDs(t *testing.T) {
	content := `3-5
10-14
`

	tmpFile, err := os.CreateTemp("", "test_no_ids_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	result, err := CountFreshIngredients(tmpFile.Name())
	if err != nil {
		t.Fatalf("CountFreshIngredients failed: %v", err)
	}

	expected := 0 // no IDs to check
	if result != expected {
		t.Errorf("CountFreshIngredients no IDs = %d; expected %d", result, expected)
	}
}

// tests handling of large ingredient IDs
func TestLargeNumbers(t *testing.T) {
	ranges := []IDRange{
		{1000000000000, 2000000000000}, // large numbers
	}

	tests := []struct {
		id       int64
		expected bool
	}{
		{500000000000, false},  // below range
		{1500000000000, true},  // within range
		{2500000000000, false}, // above range
	}

	for _, test := range tests {
		result := IsFresh(test.id, ranges)
		if result != test.expected {
			t.Errorf("IsFresh large number(%d) = %v; expected %v", test.id, result, test.expected)
		}
	}
}

// tests behavior with overlapping ranges
func TestOverlappingRanges(t *testing.T) {
	ranges := []IDRange{
		{1, 10},
		{5, 15},  // overlaps with first range
		{12, 20}, // overlaps with second range
	}

	// All IDs from 1-20 should be fresh due to overlaps
	for id := int64(1); id <= 20; id++ {
		if !IsFresh(id, ranges) {
			t.Errorf("ID %d should be fresh due to overlapping ranges", id)
		}
	}

	// IDs outside ranges should not be fresh
	if IsFresh(21, ranges) {
		t.Errorf("ID 21 should not be fresh")
	}
}

// test the range union counting for Part 2
func TestCountUniqueIDsInRanges(t *testing.T) {
	tests := []struct {
		ranges   []IDRange
		expected int64
	}{
		// Example from problem: 3-5,10-14,16-20,12-18 should give 14 unique IDs
		{
			ranges: []IDRange{
				{3, 5}, {10, 14}, {16, 20}, {12, 18},
			},
			expected: 14, // 3,4,5,10,11,12,13,14,15,16,17,18,19,20
		},
		// No ranges
		{
			ranges:   []IDRange{},
			expected: 0,
		},
		// Single range
		{
			ranges:   []IDRange{{5, 10}},
			expected: 6, // 5,6,7,8,9,10
		},
		// Overlapping ranges
		{
			ranges:   []IDRange{{1, 5}, {3, 8}, {10, 12}},
			expected: 11, // 1-8,10-12 = 11 total (1,2,3,4,5,6,7,8,10,11,12)
		},
		// Adjacent ranges (should merge)
		{
			ranges:   []IDRange{{1, 5}, {6, 10}},
			expected: 10, // 1-10 continuous
		},
	}

	for _, test := range tests {
		result := CountUniqueIDsInRanges(test.ranges)
		if result != test.expected {
			t.Errorf("CountUniqueIDsInRanges(%v) = %d; expected %d", test.ranges, result, test.expected)
		}
	}
}

// tests the complete Part 2 processing
func TestCountTotalFreshIngredients(t *testing.T) {
	// create a temporary file with test ranges
	content := `3-5
10-14
16-20
12-18
`

	tmpFile, err := os.CreateTemp("", "test_ranges_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	result, err := CountTotalFreshIngredients(tmpFile.Name())
	if err != nil {
		t.Fatalf("CountTotalFreshIngredients failed: %v", err)
	}

	expected := int64(14) // From the example in the problem
	if result != expected {
		t.Errorf("CountTotalFreshIngredients() = %d; expected %d", result, expected)
	}
}
