# Day 2: Invalid Product IDs

## Problem Description

In the Advent of Code 2025 Day 2 challenge, you need to identify invalid product IDs in given ranges. An ID is invalid if it consists entirely of a repeated digit pattern.

## Solution Approach

The solution processes comma-separated ranges of IDs and identifies invalid ones:

**Part 1**: An ID is invalid if it consists of exactly two identical halves (e.g., 55, 6464, 1010).
**Part 2**: An ID is invalid if it consists of two or more repetitions of any digit pattern (e.g., 111, 123123, 123123123).

## Implementation Details

- **Range Parsing**: Converts strings like "11-22,95-115" into IDRange structs
- **Pattern Detection**: Uses string manipulation to check for repeated patterns
- **Part 1**: Checks for exactly two equal halves of even-length strings
- **Part 2**: Tests all possible pattern lengths that divide the string evenly
- **Summation**: Accumulates invalid IDs across all specified ranges

## Testing

The solution includes comprehensive tests covering:
- Range parsing with various formats and edge cases
- Pattern detection for both validation rules
- Summation logic across multiple ranges
- Large number handling and boundary conditions
- Error handling for malformed input

## Performance

The algorithm processes each ID in each range:
- **Time Complexity**: O(R × D × L) where R is ranges, D is range size, L is ID string length
- **Space Complexity**: O(R) for storing parsed ranges

## Examples

For ranges `11-22,95-115,998-1012`:
- **Part 1**: Invalid IDs are 11, 22, 99, 1010 (sum = 1222)
- **Part 2**: Additional invalid IDs like 111 (sum = 1333)

The actual input contains much larger ranges and numbers, requiring efficient processing.
