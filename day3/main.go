/**
 * Advent of Code 2025 - Day 3: Battery Banks
 *
 * This program calculates maximum joltage from battery banks
 * by selecting optimal battery combinations.
 *
 * Author: KleaSCM
 * Email: KleaSCM@gmail.com
 */

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// converts line of digits to battery joltage values
func ParseBatteryBank(line string) ([]int, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}

	batteries := make([]int, len(line))
	for i, char := range line {
		if char < '0' || char > '9' {
			return nil, fmt.Errorf("invalid character '%c' in battery bank", char)
		}
		batteries[i] = int(char - '0')
	}

	return batteries, nil
}

// max 2-digit joltage from any two batteries
func FindMaxTwoDigitJoltage(batteries []int) int {
	if len(batteries) < 2 {
		return 0
	}

	maxJolt := 0
	for i := 0; i < len(batteries); i++ {
		for j := i + 1; j < len(batteries); j++ {
			jolt := 10*batteries[i] + batteries[j]
			if jolt > maxJolt {
				maxJolt = jolt
			}
		}
	}

	return maxJolt
}

// max numeric subsequence of length k using monotonic stack
func FindMaxSubsequence(batteries []int, k int) string {
	n := len(batteries)
	if k <= 0 {
		return ""
	}
	if k >= n {
		// Convert all batteries to string
		var result strings.Builder
		for _, battery := range batteries {
			result.WriteString(strconv.Itoa(battery))
		}
		return result.String()
	}

	// Convert batteries to string digits for easier manipulation
	digits := make([]byte, n)
	for i, battery := range batteries {
		digits[i] = byte('0' + battery)
	}

	stack := make([]byte, 0, k)

	for i := 0; i < n; i++ {
		c := digits[i]
		remaining := n - i

		// Remove smaller digits from stack if we have enough remaining digits
		for len(stack) > 0 && stack[len(stack)-1] < c && len(stack)-1+remaining >= k {
			stack = stack[:len(stack)-1]
		}

		// Add current digit if stack isn't full
		if len(stack) < k {
			stack = append(stack, c)
		}
	}

	// Ensure we have exactly k digits
	if len(stack) > k {
		stack = stack[:k]
	}

	return string(stack)
}

// total maximum joltage for p1 (2 batteries per bank)
func SumMaxJoltagesPart1(banks [][]int) int64 {
	var total int64
	for _, bank := range banks {
		maxJolt := FindMaxTwoDigitJoltage(bank)
		total += int64(maxJolt)
	}
	return total
}

// total maximum joltage for p2 (12 batteries per bank, big ints)
func SumMaxJoltagesPart2(banks [][]int) *big.Int {
	total := big.NewInt(0)

	for _, bank := range banks {
		maxSeq := FindMaxSubsequence(bank, 12)
		if maxSeq != "" {
			num := new(big.Int)
			if _, ok := num.SetString(maxSeq, 10); !ok {
				fmt.Printf("Warning: Failed to parse %q as big integer\n", maxSeq)
				continue
			}
			total.Add(total, num)
		}
	}

	return total
}

// reads and parses all battery banks from file
func ReadBatteryBanks(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var banks [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		bank, err := ParseBatteryBank(line)
		if err != nil {
			return nil, err
		}

		if bank != nil {
			banks = append(banks, bank)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return banks, nil
}

func main() {
	// Read all battery banks from input file
	banks, err := ReadBatteryBanks("input/input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	// p1: sum of maximum 2-digit joltages
	part1Total := SumMaxJoltagesPart1(banks)
	fmt.Printf("Total output joltage (p1): %d\n", part1Total)

	// p2: sum of maximum 12-digit joltages
	part2Total := SumMaxJoltagesPart2(banks)
	fmt.Printf("Total output joltage (p2): %s\n", part2Total.String())
}
