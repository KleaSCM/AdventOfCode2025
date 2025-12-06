/**
 * Advent of Code 2025 - Day 5: Cafeteria Inventory
 *
 * This program identifies fresh ingredient IDs by checking
 * which available IDs fall within the specified fresh ranges.
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

// represents a range of fresh ingredient IDs from Start to End inclusive
type IDRange struct {
	Start, End int64
}

// converts string like "3-5" to a struct
func ParseIDRange(rangeStr string) (IDRange, error) {
	parts := strings.Split(rangeStr, "-")
	if len(parts) != 2 {
		return IDRange{}, fmt.Errorf("invalid range format: %s", rangeStr)
	}

	start, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
	end, err2 := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)

	if err1 != nil || err2 != nil {
		return IDRange{}, fmt.Errorf("invalid numbers in range: %s", rangeStr)
	}

	if start > end {
		return IDRange{}, fmt.Errorf("start > end in range: %s", rangeStr)
	}

	return IDRange{Start: start, End: end}, nil
}

// checks if ingredient ID is within any fresh ranges
func IsFresh(id int64, ranges []IDRange) bool {
	for _, r := range ranges {
		if id >= r.Start && id <= r.End {
			return true
		}
	}
	return false
}

// processes input file and counts fresh ingredient IDs
func CountFreshIngredients(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var ranges []IDRange
	var availableIDs []int64
	foundBlankLine := false

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			foundBlankLine = true
			continue
		}

		if !foundBlankLine {
			// Parse ranges
			idRange, err := ParseIDRange(line)
			if err != nil {
				return 0, err
			}
			ranges = append(ranges, idRange)
		} else {
			// Parse available IDs
			id, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				return 0, err
			}
			availableIDs = append(availableIDs, id)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// Count fresh ingredients
	freshCount := 0
	for _, id := range availableIDs {
		if IsFresh(id, ranges) {
			freshCount++
		}
	}

	return freshCount, nil
}

// counts all unique fresh ingredient IDs by union of ranges (p2)
// returns int64 for massive counts (hundreds of billions of IDs)
func CountTotalFreshIngredients(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var ranges []IDRange
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			break // stop at blank line, only need ranges for p2
		}

		// Parse ranges
		idRange, err := ParseIDRange(line)
		if err != nil {
			return 0, err
		}
		ranges = append(ranges, idRange)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// Count unique IDs in the union of all ranges
	return CountUniqueIDsInRanges(ranges), nil
}

// counts unique IDs covered by union of ranges
// returns int64 to prevent overflow with large ranges
func CountUniqueIDsInRanges(ranges []IDRange) int64 {
	if len(ranges) == 0 {
		return 0
	}

	// Sort ranges by start position for easier merging
	sortedRanges := make([]IDRange, len(ranges))
	copy(sortedRanges, ranges)

	//bubble because yes
	for i := 0; i < len(sortedRanges)-1; i++ {
		for j := 0; j < len(sortedRanges)-1-i; j++ {
			if sortedRanges[j].Start > sortedRanges[j+1].Start {
				sortedRanges[j], sortedRanges[j+1] = sortedRanges[j+1], sortedRanges[j]
			}
		}
	}

	// Merge overlapping ranges
	merged := []IDRange{sortedRanges[0]}

	for i := 1; i < len(sortedRanges); i++ {
		current := sortedRanges[i]
		last := &merged[len(merged)-1]

		if current.Start <= last.End+1 { // Overlapping or adjacent
			if current.End > last.End {
				last.End = current.End
			}
		} else {
			merged = append(merged, current)
		}
	}

	// Count total unique IDs
	// FIX: Use int64 instead of int to prevent overflow with large ranges
	// ISSUE: Previously used int(r.End - r.Start + 1) which cast int64 to int,
	// causing silent overflow for ranges covering hundreds of billions of IDs
	var total int64
	for _, r := range merged {
		total += r.End - r.Start + 1
	}

	return total
}

func main() {
	// P1: Count fresh ingredients from available list
	freshCount, err := CountFreshIngredients("input/input.txt")
	if err != nil {
		fmt.Printf("Error processing ingredients: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Number of fresh ingredients (p1): %d\n", freshCount)

	// p2: count total unique fresh ingredient IDs in ranges
	totalFreshCount, err := CountTotalFreshIngredients("input/input.txt")
	if err != nil {
		fmt.Printf("Error processing total fresh ingredients: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total fresh ingredient IDs (p2): %d\n", totalFreshCount)
}
