# Solution

## Approach

Two-pass linear scan:
1. Find the maximum number of candies across all kids.
2. For each kid, check if their candies plus extra candies reaches the maximum.

## Complexity

Where `n` is the number of kids.

- **Time:** O(n) — two linear passes over the array: one to find the max, one to build the result.
- **Space:** O(n) — result slice of the same length as input. O(1) if you exclude the output, since only scalar variables (`maxCandies`, `candy`) are used.
