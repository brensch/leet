# LeetCode in Go

This repo is a runnable study track for LeetCode-style interview prep in Go.

## How to use it

1. Pick the next numbered folder from `problems/`.
2. Read that folder's `README.md` for the concept and challenge.
3. Implement the TODO function in `solution.go`.
4. Run the tests for fast feedback:

```bash
go test ./...
```

To focus on one challenge:

```bash
go test ./problems/01_two_sum -v
```

## Study order

The folders are already ordered from easier pattern-recognition problems to harder multi-step problems. Start at `01_two_sum` and move downward.

Read [SYLLABUS.md](./SYLLABUS.md) for the concept progression behind the ordering.
