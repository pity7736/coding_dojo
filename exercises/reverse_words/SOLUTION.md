# Solution: Reverse Words in a String

## Approach

Single backward pass through the string, tracking word boundaries:

1. Walk the string from right to left
2. Track the `end` position of the current word
3. When transitioning from a non-space to a space character, extract the word using slice bounds and append it to the result
4. Handle the edge case when a word starts at index 0 (no space before it)
5. Join the collected words with a single space

This avoids a second pass to reverse the words — by walking backwards, words are naturally collected in reversed order.

## Complexity

- **Time:** O(n) — single pass through the string
- **Space:** O(n) — result slice and final joined string; unavoidable since strings are immutable in Go
