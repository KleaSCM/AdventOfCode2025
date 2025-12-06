# Day 1: North Pole Security Dial

## Problem Description

In the Advent of Code 2025 Day 1 challenge, you need to help the Elves access the North Pole base by determining the password for a security dial. The dial has numbers 0-99 and starts pointing at 50. You receive a sequence of rotation instructions (L for left/decrease, R for right/increase) with distances.

## Solution Approach

The solution involves simulating the dial's movement according to the rotation instructions:

**Part 1**: Count how many times the dial ends up pointing at 0 after each complete rotation.
**Part 2**: Count every time the dial points at 0 during any rotation, including intermediate positions.

## Implementation Details

- **Rotation Parsing**: Each line like "L68" or "R48" is parsed into direction and distance
- **Dial Simulation**: Position wraps around using modulo 100 arithmetic
- **Part 1**: Jumps directly to final position after each rotation
- **Part 2**: Steps through each click individually to catch intermediate zeros
- **Error Handling**: Validates input format and file operations

## Testing

The solution includes comprehensive tests covering:
- Rotation parsing with valid and invalid inputs
- Both simulation algorithms with the problem's example
- Edge cases like large rotations and position wrapping
- Boundary conditions and error scenarios

## Performance

Both parts process each rotation in the input:
- **Part 1**: O(N) where N is number of rotations
- **Part 2**: O(D) where D is total distance rotated (can be much larger)

## Example

For the example rotations `L68 L30 R48 L5 R60 L55 L1 L99 R14 L82`:
- **Part 1**: 3 (dial ends at 0 three times)
- **Part 2**: 6 (dial passes through 0 six times total)
