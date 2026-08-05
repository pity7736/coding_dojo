# Merge Strings Alternately

Given two strings `word1` and `word2`, merge them by adding letters in alternating order, starting with `word1`. If a string is longer than the other, append the additional letters onto the end of the merged string.

## Examples

| word1 | word2 | output |
|-------|-------|--------|
| `"abc"` | `"pqr"` | `"apbqcr"` |
| `"ab"` | `"pqrs"` | `"apbqrs"` |
| `"abcd"` | `"pq"` | `"apbqcd"` |

## Constraints

- `1 <= word1.length, word2.length <= 100`
- `word1` and `word2` may contain any unicode characters.
