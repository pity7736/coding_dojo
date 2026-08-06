# Solution

## Approach

Single-pass greedy scan: iterate through the flowerbed, and at each empty spot check if both neighbors are available (empty or out of bounds). Track the index of the last planted flower to avoid adjacent placements without mutating the input.

## Alternative approach

Skip the next index when planting (`i++`), eliminating the need for a separate tracker. When you plant at position `i`, position `i+1` can never be valid, so you skip it. Requires a classic `for` loop instead of `range`.

## Complexity

Where `n` is the length of the flowerbed.

- **Time:** O(n) — single pass through the array.
- **Space:** O(1) — only scalar variables used (no copy of the input, no mutation).
