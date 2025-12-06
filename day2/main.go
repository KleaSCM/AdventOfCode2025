/**
 * Advent of Code 2025 - Day 2: Invalid Product IDs
 *
 * This program identifies invalid product IDs in given ranges.
 * An ID is invalid if it consists of a repeated digit pattern.
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

// IDRange represents a range of product IDs from Start to End inclusive
type IDRange struct {
	Start, End int
}

// converts a string like "11-22" into a struct
func ParseIDRange(rangeStr string) (IDRange, error) {
	parts := strings.Split(rangeStr, "-")
	if len(parts) != 2 {
		return IDRange{}, fmt.Errorf("invalid range format: %s", rangeStr)
	}

	start, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	end, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

	if err1 != nil || err2 != nil {
		return IDRange{}, fmt.Errorf("invalid numbers in range: %s", rangeStr)
	}

	if start > end {
		return IDRange{}, fmt.Errorf("start > end in range: %s", rangeStr)
	}

	return IDRange{Start: start, End: end}, nil
}

// ParseIDRanges converts a comma-separated line of ranges into a slice of IDRange
func ParseIDRanges(line string) ([]IDRange, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}

	rangeStrings := strings.Split(line, ",")
	var ranges []IDRange

	for _, rangeStr := range rangeStrings {
		rangeStr = strings.TrimSpace(rangeStr)
		if rangeStr == "" {
			continue
		}

		idRange, err := ParseIDRange(rangeStr)
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, idRange)
	}

	return ranges, nil
}

// checks if an ID is invalid for p1
// an ID is invalid if it consists of exactly two identical halves
func IsInvalidIDPart1(id int) bool {
	s := strconv.Itoa(id)
	length := len(s)

	// Must be even length to be split into two equal halves
	if length%2 != 0 {
		return false
	}

	half := length / 2
	firstHalf := s[:half]
	secondHalf := s[half:]

	return firstHalf == secondHalf
}

// checks if an ID is invalid for p2
// an ID is invalid if it consists of two or more repetitions of the same digit pattern
func IsInvalidIDPart2(id int) bool {
	s := strconv.Itoa(id)
	length := len(s)

	// Try all possible pattern lengths that divide the total length
	for patternLen := 1; patternLen <= length/2; patternLen++ {
		if length%patternLen != 0 {
			continue
		}

		pattern := s[:patternLen]
		repeatCount := length / patternLen

		// Must repeat at least twice
		if repeatCount < 2 {
			continue
		}

		// Check if the entire string matches this repeated pattern
		allMatch := true
		for i := 0; i < length; i += patternLen {
			if s[i:i+patternLen] != pattern {
				allMatch = false
				break
			}
		}

		if allMatch {
			return true
		}
	}

	return false
}

// calculates the sum of all invalid IDs within the given ranges
func SumInvalidIDsInRanges(ranges []IDRange, isInvalidFunc func(int) bool) int64 {
	var sum int64

	for _, idRange := range ranges {
		for id := idRange.Start; id <= idRange.End; id++ {
			if isInvalidFunc(id) {
				sum += int64(id)
			}
		}
	}

	return sum
}

// reads the input file and processes all ID ranges
func ReadAndProcessRanges(filename string) ([]IDRange, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var allRanges []IDRange
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		ranges, err := ParseIDRanges(line)
		if err != nil {
			return nil, err
		}

		allRanges = append(allRanges, ranges...)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return allRanges, nil
}

func main() {
	// Read all ID ranges from input file
	ranges, err := ReadAndProcessRanges("input/input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	// p1: sum IDs that are exactly two identical halves
	part1Sum := SumInvalidIDsInRanges(ranges, IsInvalidIDPart1)
	fmt.Printf("Sum of invalid IDs (p1): %d\n", part1Sum)

	// p2: sum IDs that are two or more repetitions of any pattern
	part2Sum := SumInvalidIDsInRanges(ranges, IsInvalidIDPart2)
	fmt.Printf("Sum of invalid IDs (p2): %d\n", part2Sum)
}
