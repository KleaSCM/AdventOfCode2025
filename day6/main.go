/**
 * Advent of Code 2025 - Day 6: Trash Compactor Math (P1,2)
 *
 * This program parses visual math worksheets in both left-to-right
 * (P1) and right-to-left cephalopod format (P2).
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

type CephalopodColumn struct {
	numbers   []int
	operation string // "addition" or "multiplication"
}

func main() {
	input := readLines("input.txt")

	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))
}

func readLines(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func solvePart1(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	numberOfColumns := len(extractNumbers(lines[0]))
	numberOfRows := len(lines) - 1

	var columns []CephalopodColumn
	for i := 0; i < numberOfColumns; i++ {
		var nums []int
		for j := 0; j < numberOfRows; j++ {
			numLine := extractNumbers(lines[j])
			nums = append(nums, numLine[i])
		}

		opSymbols := extractOperators(lines[numberOfRows])
		op := "multiplication"
		if opSymbols[i] == "+" {
			op = "addition"
		}

		columns = append(columns, CephalopodColumn{numbers: nums, operation: op})
	}

	sum := 0
	for _, col := range columns {
		sum += calculateColumnValue(col)
	}

	return sum
}

func extractNumbers(line string) []int {
	var nums []int
	for _, part := range strings.Fields(line) {
		n, _ := strconv.Atoi(part)
		nums = append(nums, n)
	}
	return nums
}

func extractOperators(line string) []string {
	var ops []string
	for _, c := range line {
		if c == '+' || c == '*' {
			ops = append(ops, string(c))
		}
	}
	return ops
}

func calculateColumnValue(col CephalopodColumn) int {
	if col.operation == "addition" {
		sum := 0
		for _, n := range col.numbers {
			sum += n
		}
		return sum
	} else {
		prod := 1
		for _, n := range col.numbers {
			prod *= n
		}
		return prod
	}
}

func solvePart2(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	// pad lines to same width
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		padded := line + strings.Repeat(" ", maxWidth-len(line))
		grid[i] = []rune(padded)
	}

	rows := len(grid)
	cols := maxWidth
	opRow := rows - 1

	isBlankColumn := func(c int) bool {
		for r := 0; r < rows; r++ {
			if grid[r][c] != ' ' {
				return false
			}
		}
		return true
	}

	type problem struct {
		columns  []int
		operator rune
	}

	var problems []problem
	var current []int

	for c := 0; c < cols; c++ {
		if isBlankColumn(c) {
			if len(current) > 0 {
				op := rune(0)
				for _, col := range current {
					if grid[opRow][col] == '+' || grid[opRow][col] == '*' {
						op = grid[opRow][col]
						break
					}
				}
				problems = append(problems, problem{columns: current, operator: op})
			}
			current = nil
		} else {
			current = append(current, c)
		}
	}
	if len(current) > 0 {
		op := rune(0)
		for _, col := range current {
			if grid[opRow][col] == '+' || grid[opRow][col] == '*' {
				op = grid[opRow][col]
				break
			}
		}
		problems = append(problems, problem{columns: current, operator: op})
	}

	result := 0
	for pi := len(problems) - 1; pi >= 0; pi-- {
		p := problems[pi]
		var numbers []int
		for i := len(p.columns) - 1; i >= 0; i-- {
			col := p.columns[i]
			numStr := ""
			for r := 0; r < opRow; r++ {
				if grid[r][col] != ' ' {
					numStr += string(grid[r][col])
				}
			}
			if numStr != "" {
				n, _ := strconv.Atoi(numStr)
				numbers = append(numbers, n)
			}
		}

		sum := 0
		prod := 1
		if p.operator == '+' {
			for _, n := range numbers {
				sum += n
			}
			result += sum
		} else {
			for _, n := range numbers {
				prod *= n
			}
			result += prod
		}
	}

	return result
}
