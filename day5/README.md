# Day 5: Cafeteria Inventory

## Problem Description

In the Advent of Code 2025 Day 5 challenge, the Elves need help identifying which ingredient IDs are fresh using their new inventory management system. The system provides:

- A list of fresh ingredient ID ranges (inclusive, may overlap)
- A blank line separator
- A list of available ingredient IDs to check

An ingredient ID is fresh if it falls within any of the specified ranges.

## Solution Approach

**Part 1:**
1. Parse fresh ID ranges from lines before the blank separator
2. Parse available ingredient IDs from lines after the blank separator
3. For each available ID, check if it falls within any fresh range
4. Count and return the total number of fresh ingredients

**Part 2:**
1. Parse only the fresh ID ranges (stop at blank line)
2. Find the union of all ranges (handle overlaps)
3. Count all unique ingredient IDs covered by any range
4. Return the total count of fresh ingredient IDs

## Implementation Details

- **Range Parsing**: Converts "start-end" strings to IDRange structs with int64 for large numbers
- **Freshness Checking**: Linear search through ranges (acceptable for reasonable input sizes)
- **Range Union**: Merges overlapping ranges using sorting and interval merging algorithm
- **File Processing**: Two-phase parsing with blank line detection for Part 1, range-only for Part 2
- **Error Handling**: Validates input format and provides meaningful error messages
- **Data Types**: Uses int64 to handle large ingredient IDs (up to 16+ digits)

## Testing

The solution includes comprehensive tests covering:
- Range parsing with various formats and error conditions
- Freshness checking with overlapping and edge case ranges
- Range union calculations for Part 2
- Complete file processing with temporary test files for both parts
- Large number handling for realistic ingredient IDs
- Boundary conditions, empty inputs, and edge cases

## Performance

- **Time Complexity**: O(R × A) where R is number of ranges, A is number of available IDs
- **Space Complexity**: O(R + A) for storing ranges and IDs
- **File I/O**: Single pass through input file with buffered reading

## Examples

**Part 1:** For ranges `3-5,10-14,16-20,12-18` and available IDs `1,5,8,11,17,32`:
- Fresh IDs: 5 (in 3-5), 11 (in 10-14), 17 (in 16-20 and 12-18)
- Result: 3 fresh ingredients

**Part 2:** For ranges `3-5,10-14,16-20,12-18` (ignoring available IDs):
- Union of ranges covers: 3,4,5,10,11,12,13,14,15,16,17,18,19,20
- Result: 14 total fresh ingredient IDs

The actual input contains hundreds of ranges and over 1000 ingredient IDs to process for Part 1, with massive ranges covering hundreds of billions of IDs for Part 2.
