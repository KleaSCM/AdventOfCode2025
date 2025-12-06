# Day 6: Trash Compactor Math

## Problem Description

In the Advent of Code 2025 Day 6 challenge, you need to help cephalopod students check their math homework. The worksheet contains multiple math problems arranged in a visual grid format where:

- Numbers are arranged vertically in columns
- Each column represents one problem
- The bottom of each column contains the operation (+ or *)
- Problems are separated by spacing
- You need to perform the operation on all numbers in each column and sum the results

## Part 1: Left-to-Right Reading

The initial solution reads the worksheet left-to-right, where each column contains numbers that are operated on with the operation at the bottom.

## Solution Approach

1. Parse the input file with multiple lines of numbers and one line of operations
2. Split each line by whitespace to identify individual numbers and operations
3. Group numbers by their column positions to match them with operations
4. For each operation, collect all numbers in that column
5. Perform the mathematical operation (+ or *) on all numbers in each column
6. Sum all the individual problem results

## Part 1 Example

For the worksheet:
```
123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +
```

Results in 4 problems:
- 123 × 45 × 6 = 33,210
- 328 + 64 + 98 = 490
- 51 × 387 × 215 = 4,243,455
- 64 + 23 + 314 = 401

**Part 1 Total**: 33,210 + 490 + 4,243,455 + 401 = 4,277,556

## Part 2: Right-to-Left Reading

- **Input Parsing**: Uses `strings.Fields()` to split lines by whitespace
- **Column Alignment**: Matches operation positions with corresponding number columns
- **Number Parsing**: Handles multi-digit numbers using `strconv.Atoi()`
- **Operation Handling**: Supports both addition (+) and multiplication (*) operations
- **Error Handling**: Validates input format and handles parsing errors

## Testing

The solution includes comprehensive tests covering:
- Individual problem solving with various operations
- Complete worksheet parsing with example data
- Edge cases: empty worksheets, single numbers, large numbers
- File I/O operations with temporary test files

## Performance

- **Time Complexity**: O(L × W) where L is lines and W is numbers per line
- **Space Complexity**: O(P × N) where P is problems and N is numbers per problem
- **Execution**: Linear scan through input with minimal memory usage

## Part 2: Right-to-Left Reading

In part 2, cephalopod math is read right-to-left. Each column is processed from right to left, and numbers are formed by collecting digits at each position from the right, reading top to bottom as most significant to least significant.

For the example worksheet:
```
123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +
```

Reading right-to-left gives 4 problems:
- Rightmost: 4 + 431 + 623 = 1,058
- Third: 175 × 581 × 32 = 3,253,600
- Second: 8 + 248 + 369 = 625
- Leftmost: 356 × 24 × 1 = 8,544

**Total**: 1,058 + 3,253,600 + 625 + 8,544 = 3,263,827

The actual input contains hundreds of numbers across multiple lines with complex spacing patterns.

**Part 1 Result**: 4,805,473,544,166
**Part 2 Result**: 8,907,730,960,817
