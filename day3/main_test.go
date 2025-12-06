/**
 * Test suite for Advent of Code 2025 - Day 3: Battery Banks
 *
 * Tests verify the joltage calculation and battery selection logic.
 *
 * Author: KleaSCM
 * Email: KleaSCM@gmail.com
 */

package main

import (
	"math/big"
	"os"
	"testing"
)

// parse bank
func TestParseBatteryBank(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
		hasError bool
	}{
		{"12345", []int{1, 2, 3, 4, 5}, false},
		{"987654321111111", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, false},
		{"", nil, false},                     // empty line
		{"   123   ", []int{1, 2, 3}, false}, // trimmed
		{"123a45", nil, true},                // invalid character
		{"1.234", nil, true},                 // invalid character
	}

	for _, test := range tests {
		result, err := ParseBatteryBank(test.input)

		if test.hasError {
			if err == nil {
				t.Errorf("ParseBatteryBank(%q) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ParseBatteryBank(%q) unexpected error: %v", test.input, err)
			}
			if len(result) != len(test.expected) {
				t.Errorf("ParseBatteryBank(%q) length = %d; expected %d", test.input, len(result), len(test.expected))
				continue
			}
			for i, expected := range test.expected {
				if result[i] != expected {
					t.Errorf("ParseBatteryBank(%q)[%d] = %d; expected %d", test.input, i, result[i], expected)
				}
			}
		}
	}
}

// max two digit
func TestFindMaxTwoDigitJoltage(t *testing.T) {
	tests := []struct {
		batteries []int
		expected  int
	}{
		{[]int{1, 2, 3, 4, 5}, 45},                               // 4 and 5 = 45 (positions 3,4)
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, 98}, // 9 and 8 = 98 (positions 0,1)
		{[]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, 89},    // 8 and 9 = 89 (positions 0,13)
		{[]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, 78}, // 7 and 8 = 78 (positions 13,14)
		{[]int{1}, 0},     // too few batteries
		{[]int{}, 0},      // empty bank
		{[]int{5, 5}, 55}, // same digits
	}

	for _, test := range tests {
		result := FindMaxTwoDigitJoltage(test.batteries)
		if result != test.expected {
			t.Errorf("FindMaxTwoDigitJoltage(%v) = %d; expected %d", test.batteries, result, test.expected)
		}
	}
}

// max subsequence
func TestFindMaxSubsequence(t *testing.T) {
	tests := []struct {
		batteries []int
		k         int
		expected  string
	}{
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, 12, "987654321111"},
		{[]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, 12, "811111111119"},
		{[]int{1, 2, 3, 4, 5}, 2, "45"},
		{[]int{1, 2, 3, 4, 5}, 5, "12345"}, // k >= n
		{[]int{1, 2, 3, 4, 5}, 0, ""},      // k = 0
		{[]int{}, 2, ""},                   // empty input
	}

	for _, test := range tests {
		result := FindMaxSubsequence(test.batteries, test.k)
		if result != test.expected {
			t.Errorf("FindMaxSubsequence(%v, %d) = %q; expected %q", test.batteries, test.k, result, test.expected)
		}
	}
}

// sum p1
func TestSumMaxJoltagesPart1(t *testing.T) {
	banks := [][]int{
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, // 98
		{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},    // 89
		{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, // 78
	}

	result := SumMaxJoltagesPart1(banks)
	expected := int64(98 + 89 + 78) // 265

	if result != expected {
		t.Errorf("SumMaxJoltagesPart1() = %d; expected %d", result, expected)
	}
}

// sum p2
func TestSumMaxJoltagesPart2(t *testing.T) {
	banks := [][]int{
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, // "987654321111"
		{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},    // "811111111119"
	}

	result := SumMaxJoltagesPart2(banks)

	// Calculate expected: 987654321111 + 811111111119
	expected := big.NewInt(0)
	expected.SetString("987654321111", 10)
	addend := big.NewInt(0)
	addend.SetString("811111111119", 10)
	expected.Add(expected, addend)

	if result.Cmp(expected) != 0 {
		t.Errorf("SumMaxJoltagesPart2() = %s; expected %s", result.String(), expected.String())
	}
}

// read banks
func TestReadBatteryBanks(t *testing.T) {
	// Create a temporary file with test data
	content := "12345\n67890\n"

	tmpFile, err := os.CreateTemp("", "test_batteries_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	banks, err := ReadBatteryBanks(tmpFile.Name())
	if err != nil {
		t.Fatalf("ReadBatteryBanks failed: %v", err)
	}

	expected := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 0},
	}

	if len(banks) != len(expected) {
		t.Fatalf("Expected %d banks, got %d", len(expected), len(banks))
	}

	for i, exp := range expected {
		if len(banks[i]) != len(exp) {
			t.Errorf("Bank %d length: expected %d, got %d", i, len(exp), len(banks[i]))
			continue
		}
		for j, val := range exp {
			if banks[i][j] != val {
				t.Errorf("Bank %d[%d]: expected %d, got %d", i, j, val, banks[i][j])
			}
		}
	}
}

// empty banks
func TestEmptyBanks(t *testing.T) {
	banks := [][]int{}

	part1Result := SumMaxJoltagesPart1(banks)
	part2Result := SumMaxJoltagesPart2(banks)

	if part1Result != 0 {
		t.Errorf("Empty banks part1 sum = %d; expected 0", part1Result)
	}

	expectedZero := big.NewInt(0)
	if part2Result.Cmp(expectedZero) != 0 {
		t.Errorf("Empty banks part2 sum = %s; expected 0", part2Result.String())
	}
}

// single bank
func TestSingleBatteryBank(t *testing.T) {
	banks := [][]int{
		{5}, // only one battery
	}

	part1Result := SumMaxJoltagesPart1(banks)
	if part1Result != 0 {
		t.Errorf("Single battery bank part1 = %d; expected 0", part1Result)
	}
}

// same digits
func TestAllSameDigits(t *testing.T) {
	banks := [][]int{
		{2, 2, 2, 2, 2}, // all 2s
	}

	part1Result := SumMaxJoltagesPart1(banks)
	expected := 22 // 2 and 2 = 22
	if part1Result != int64(expected) {
		t.Errorf("All same digits part1 = %d; expected %d", part1Result, expected)
	}
}
