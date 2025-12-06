/**
 * Test suite for Advent of Code 2025 - Day 6: Trash Compactor Math (Parts 1 & 2)
 *
 * Tests verify worksheet parsing in both left-to-right (Part 1) and
 * right-to-left cephalopod format (Part 2), plus math problem solving.
 *
 * Author: KleaSCM
 * Email: KleaSCM@gmail.com
 */

package main

import (
	"os"
	"testing"
)

// tests individual math problem solving
func TestSolveProblem(t *testing.T) {
	tests := []struct {
		problem  MathProblem
		expected int
	}{
		//Addition problems
		{MathProblem{Numbers: []int{328, 64, 98}, Op: '+'}, 490},
		{MathProblem{Numbers: []int{64, 23, 314}, Op: '+'}, 401},
		{MathProblem{Numbers: []int{5, 10, 15}, Op: '+'}, 30},

		// Multiplication probs
		{MathProblem{Numbers: []int{123, 45, 6}, Op: '*'}, 33210},
		{MathProblem{Numbers: []int{51, 387, 215}, Op: '*'}, 4243455},
		{MathProblem{Numbers: []int{2, 3, 4}, Op: '*'}, 24},

		//ugly cases
		{MathProblem{Numbers: []int{5}, Op: '+'}, 5},
		{MathProblem{Numbers: []int{7}, Op: '*'}, 7},
		{MathProblem{Numbers: []int{}, Op: '+'}, 0},
	}

	for _, test := range tests {
		result := SolveProblem(test.problem)
		if result != test.expected {
			t.Errorf("SolveProblem(%v) = %d; expected %d", test.problem, result, test.expected)
		}
	}
}

// test the worksheet parsing with example data
func TestParseWorksheet(t *testing.T) {
	// Create testsheet
	content := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	tmpFile, err := os.CreateTemp("", "test_worksheet_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	problems, err := ParseWorksheetPart1(tmpFile.Name())
	if err != nil {
		t.Fatalf("ParseWorksheetPart1 failed: %v", err)
	}

	expected := []MathProblem{
		{Numbers: []int{123, 45, 6}, Op: '*'},
		{Numbers: []int{328, 64, 98}, Op: '+'},
		{Numbers: []int{51, 387, 215}, Op: '*'},
		{Numbers: []int{64, 23, 314}, Op: '+'},
	}

	if len(problems) != len(expected) {
		t.Fatalf("Expected %d problems, got %d", len(expected), len(problems))
	}

	for i, exp := range expected {
		if len(problems[i].Numbers) != len(exp.Numbers) {
			t.Errorf("Problem %d: expected %d numbers, got %d", i, len(exp.Numbers), len(problems[i].Numbers))
			continue
		}

		for j, num := range exp.Numbers {
			if problems[i].Numbers[j] != num {
				t.Errorf("Problem %d, number %d: expected %d, got %d", i, j, num, problems[i].Numbers[j])
			}
		}

		if problems[i].Op != exp.Op {
			t.Errorf("Problem %d: expected op %c, got %c", i, exp.Op, problems[i].Op)
		}
	}
}

// complete worksheet p1
func TestSolveWorksheetPart1(t *testing.T) {
	//Test example worksheet
	content := `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `

	tmpFile, err := os.CreateTemp("", "test_worksheet_p1_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	result, err := SolveWorksheetPart1(tmpFile.Name())
	if err != nil {
		t.Fatalf("SolveWorksheetPart1 failed: %v", err)
	}

	// Expected: 33210 + 490 + 4243455 + 401 = 4277556
	expected := 33210 + 490 + 4243455 + 401
	if result != expected {
		t.Errorf("SolveWorksheetPart1() = %d; expected %d", result, expected)
	}
}

// test struct
func TestMathProblemType(t *testing.T) {
	problem := MathProblem{
		Numbers: []int{1, 2, 3},
		Op:      '+',
	}

	if len(problem.Numbers) != 3 {
		t.Errorf("Expected 3 numbers, got %d", len(problem.Numbers))
	}

	if problem.Op != '+' {
		t.Errorf("Expected op '+', got %c", problem.Op)
	}
}

// test behavior with minimal input
func TestEmptyWorksheet(t *testing.T) {
	content := `
*`

	tmpFile, err := os.CreateTemp("", "test_empty_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	result, err := SolveWorksheet(tmpFile.Name())
	if err != nil {
		t.Fatalf("SolveWorksheet failed: %v", err)
	}

	//np maybe? idk
	if result != 0 {
		t.Errorf("Empty worksheet result = %d; expected 0", result)
	}
}

// single num
func TestSingleNumberProblems(t *testing.T) {
	content := `5
+`

	tmpFile, err := os.CreateTemp("", "test_single_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	result, err := SolveWorksheet(tmpFile.Name())
	if err != nil {
		t.Fatalf("SolveWorksheet failed: %v", err)
	}

	// Single num with any operation should return that num?
	expected := 5
	if result != expected {
		t.Errorf("Single number result = %d; expected %d", result, expected)
	}
}

// handling of larger num
func TestLargeNumbers(t *testing.T) {
	problem := MathProblem{
		Numbers: []int{100, 200, 300},
		Op:      '*',
	}

	result := SolveProblem(problem)
	expected := 100 * 200 * 300 // 6,000,000
	if result != expected {
		t.Errorf("Large numbers result = %d; expected %d", result, expected)
	}
}

// right-to-left parsing with example
func TestParseWorksheetPart2(t *testing.T) {
	//create test worksheet content
	content := `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `

	tmpFile, err := os.CreateTemp("", "test_worksheet_part2_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	problems, err := ParseWorksheetPart2(tmpFile.Name())
	if err != nil {
		t.Fatalf("ParseWorksheetPart2 failed: %v", err)
	}

	// Expected right-to-left processing order:
	// 0. rightmost column: 4 + 431 + 623
	// 1. 3rd column: 175 * 581 * 32
	// 2. 2nd column: 8 + 248 + 369
	// 3. leftmost column: 356 * 24 * 1

	expected := []MathProblem{
		{Numbers: []int{4, 431, 623}, Op: '+'},  // rightmost column
		{Numbers: []int{175, 581, 32}, Op: '*'}, // 3rd column
		{Numbers: []int{8, 248, 369}, Op: '+'},  // 2nd column
		{Numbers: []int{356, 24, 1}, Op: '*'},   // eftmost column
	}

	if len(problems) != len(expected) {
		t.Fatalf("Expected %d problems, got %d", len(expected), len(problems))
	}

	for i, exp := range expected {
		if len(problems[i].Numbers) != len(exp.Numbers) {
			t.Errorf("Problem %d: expected %d numbers, got %d", i, len(exp.Numbers), len(problems[i].Numbers))
			continue
		}

		for j, num := range exp.Numbers {
			if problems[i].Numbers[j] != num {
				t.Errorf("Problem %d, number %d: expected %d, got %d", i, j, num, problems[i].Numbers[j])
			}
		}

		if problems[i].Op != exp.Op {
			t.Errorf("Problem %d: expected op %c, got %c", i, exp.Op, problems[i].Op)
		}
	}
}

// P2 worksheet
func TestSolveWorksheetPart2(t *testing.T) {
	// Test example workshee t
	content := `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `

	tmpFile, err := os.CreateTemp("", "test_worksheet_solve_part2_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	result, err := SolveWorksheet(tmpFile.Name())
	if err != nil {
		t.Fatalf("SolveWorksheet failed: %v", err)
	}

	// Expected: 356*24*1 + 8+248+369 + 175*581*32 + 4+431+623 = 8544 + 625 + 3253600 + 1058 = 3263827
	expected := 356*24*1 + 8 + 248 + 369 + 175*581*32 + 4 + 431 + 623
	if result != expected {
		t.Errorf("SolveWorksheet() = %d; expected %d", result, expected)
	}
}
