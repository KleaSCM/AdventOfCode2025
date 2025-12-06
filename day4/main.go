/**
 * Advent of Code 2025 - Day 4: Printing Department
 *
 * This program solves the forklift accessibility problem by counting
 * paper rolls (@) that have fewer than 4 adjacent rolls in their
 * 8 neighboring positions.
 *
 * Author: KleaSCM
 * Email: KleaSCM@gmail.com
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Grid represents the 2D grid of paper rolls
type Grid [][]rune

// reads the input file and returns a Grid
func ReadInput(filename string) (Grid, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid Grid
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for i, char := range line {
			row[i] = char
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

// counts the number of '@' in the 8 adjacent positions
func CountAdjacentRolls(grid Grid, row, col int) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Check all 8 directions
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			if grid[newRow][newCol] == '@' {
				count++
			}
		}
	}

	return count
}

// checks if a roll at position (row, col) is accessible
func IsAccessible(grid Grid, row, col int) bool {
	if grid[row][col] != '@' {
		return false
	}
	return CountAdjacentRolls(grid, row, col) < 4
}

// counts all accessible rolls in the grid
func CountAccessibleRolls(grid Grid) int {
	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if IsAccessible(grid, row, col) {
				count++
			}
		}
	}
	return count
}

// counts total rolls that can be removed through iterative process
// implements the p2 algorithm where accessible rolls are removed iteratively
func CountTotalRemovableRolls(grid Grid) int {
	totalRemoved := 0

	// Create a copy of the grid to modify
	currentGrid := make(Grid, len(grid))
	for i := range grid {
		currentGrid[i] = make([]rune, len(grid[i]))
		copy(currentGrid[i], grid[i])
	}

	for {
		// Find all currently accessible rolls
		accessible := make([][2]int, 0)

		for row := 0; row < len(currentGrid); row++ {
			for col := 0; col < len(currentGrid[row]); col++ {
				if IsAccessible(currentGrid, row, col) {
					accessible = append(accessible, [2]int{row, col})
				}
			}
		}

		// If no more accessible rolls, break
		if len(accessible) == 0 {
			break
		}

		// Remove all accessible rolls (change '@' to '.')
		for _, pos := range accessible {
			currentGrid[pos[0]][pos[1]] = '.'
		}

		// Add to total count
		totalRemoved += len(accessible)
	}

	return totalRemoved
}

func main() {
	grid, err := ReadInput("input/input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	// p1: count initially accessible rolls
	part1Result := CountAccessibleRolls(grid)
	fmt.Printf("Number of initially accessible rolls (p1): %d\n", part1Result)

	// p2: count total removable rolls through iterative process
	part2Result := CountTotalRemovableRolls(grid)
	fmt.Printf("Total removable rolls (p2): %d\n", part2Result)
}
