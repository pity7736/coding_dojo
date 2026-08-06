# Solution: Reverse Vowels of a String

## Approach: Two Pointers

Use two indices — `left` starting at the beginning and `right` starting at the end — moving inward. On each iteration, check if both point to vowels: if so, swap and advance both. Otherwise, advance whichever pointer is not on a vowel.

`strings.ContainsRune` simplifies the vowel check by searching a constant string of all vowels (both cases) instead of chaining multiple comparisons.

## Complexity

- **Time:** O(n) — each pointer moves inward without backtracking, visiting each element at most once.
- **Space:** O(n) — a `[]rune` copy of the string is needed since Go strings are immutable.

## Alternative

Two inner loops can independently advance each pointer past consonants before swapping. Same complexity, fewer redundant `isVowel` calls, but less explicit about the branching logic.
