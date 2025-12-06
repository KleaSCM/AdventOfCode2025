# Day 4: Printing Department - Forklift Accessibility Problem

## Problem Description

In the Advent of Code 2025 Day 4 challenge, we need to help the forklifts in the printing department access and remove paper rolls. A forklift can access a paper roll (@) if it has fewer than 4 adjacent paper rolls in its 8 neighboring positions.

**Part 1**: Count how many rolls are initially accessible.
**Part 2**: Simulate the iterative removal process where accessible rolls are removed, potentially making more rolls accessible, until no more can be removed.

## Solution Approach

**Part 1**:
1. Read the input grid from a text file
2. For each paper roll (@), count adjacent rolls in 8 directions
3. Count rolls with fewer than 4 adjacent rolls

**Part 2**:
1. Start with initial grid
2. Repeatedly find and remove all currently accessible rolls
3. After each removal, more rolls may become accessible
4. Continue until no accessible rolls remain
5. Count total rolls removed across all iterations

## Implementation Details

- **Grid Representation**: Uses a slice of rune slices to represent the 2D grid
- **Adjacent Counting**: Checks all 8 directions (up, down, left, right, diagonals) for each position
- **Boundary Checking**: Ensures adjacent position calculations don't go out of bounds
- **Error Handling**: Returns errors for file I/O operations instead of panicking
- **Part 2 Iterative Process**: Repeatedly finds accessible rolls, removes them, and continues until none remain

## Testing

The solution includes comprehensive tests covering:
- Adjacent roll counting for various grid positions
- Accessibility determination logic
- Part 1: Initial accessible roll counting
- Part 2: Iterative removal process with example verification
- Edge cases: empty grids, small grids, full removal scenarios
- File I/O operations with temporary files

## Performance

- **Part 1**: O(rows × columns) - single pass through grid
- **Part 2**: O(iterations × rows × columns) - multiple passes until convergence
- **Space**: O(rows × columns) for grid storage
- **File I/O**: Proper error handling with buffered reading
