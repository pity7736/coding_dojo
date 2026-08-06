# Coding Dojo

Practice repository for data structures, algorithms, and coding exercises.

## Structure

- `data_structures/` — custom implementations (binary tree, linked lists)
- `exercises/` — LeetCode-style coding exercises
- `games/` — game projects (snake in Go and Rust)
- `c_dojo/` — C language exercises

## Go conventions

- Module: `coding_dojo`
- Testing: `github.com/stretchr/testify/assert`
- Test style: table-driven subtests with `t.Run` and descriptive names
- Package naming: lowercase, no underscores (e.g., `kidswithcandies`)
- Run tests: `go test ./...`

## Interaction rules

- When the user says they implemented something, says "done", or asks to "check it", assume tests are passing. Never ask "does it pass?" or "is it green?"

## Exercise layout

Each exercise in `exercises/` follows this structure:
```
exercises/<exercise_name>/
  README.md               — problem description
  SOLUTION.md             — approach and complexity analysis
  <exercise_name>.go      — solution
  <exercise_name>_test.go — tests
```
