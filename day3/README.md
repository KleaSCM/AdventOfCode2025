# Day 3: Battery Banks

## Problem Description

In the Advent of Code 2025 Day 3 challenge, you need to power an escalator using battery banks. Each bank contains batteries with joltage ratings (1-9) arranged in a sequence. You must select batteries to maximize the joltage output.

## Solution Approach

**Part 1**: Select exactly 2 batteries from each bank to form the maximum possible 2-digit number.
**Part 2**: Select exactly 12 batteries from each bank to form the maximum possible 12-digit number.

The joltage is calculated by concatenating the selected battery joltages in their original order.

## Implementation Details

- **Part 1**: Brute force checks all pairs of batteries to find the maximum 2-digit combination
- **Part 2**: Uses monotonic stack algorithm to select the maximum subsequence of 12 digits
- **Monotonic Stack**: Removes smaller digits when larger digits are available later
- **Big Integers**: Part 2 uses math/big for handling large 12-digit numbers
- **Error Handling**: Validates input format and handles edge cases

## Testing

The solution includes comprehensive tests covering:
- Battery bank parsing with various inputs and error conditions
- Maximum joltage calculation for both parts with example data
- Edge cases like empty banks, single batteries, and uniform joltages
- Summation logic across multiple banks
- Monotonic stack algorithm correctness

## Performance

- **Part 1**: O(B × N²) where B is banks, N is batteries per bank
- **Part 2**: O(B × N) using efficient monotonic stack approach
- **Space**: O(N) for stack operations, O(B) for storing banks

## Examples

For bank `12345`:
- **Part 1**: Select positions 3,4 → batteries 4,5 → joltage `45`
- **Part 2**: Select all 5 batteries → joltage `12345`

For bank `987654321111111`:
- **Part 1**: Select positions 0,1 → batteries 9,8 → joltage `98`
- **Part 2**: Select optimal 12 batteries → maximum 12-digit number
