# LeetCode Syllabus in Go

This syllabus is organized in recommended study order. The numbering is intentional: each problem adds a new pattern while keeping the implementation burden manageable, so the overall difficulty ramps up over time.

## 1. Arrays and Hash Maps
- Folder: `problems/01_two_sum`
- Challenge: `two_sum`
- Difficulty: easiest starting point
- Why first: introduces the core interview move of replacing nested loops with constant-time lookup.
- Learn: maps, slices, zero values, index tracking.
- Goal: recognize when a brute force `O(n^2)` scan can become `O(n)` with a map.

## 2. Two Pointers
- Folder: `problems/02_valid_palindrome`
- Challenge: `valid_palindrome`
- Difficulty: easy
- Why here: adds pointer invariants without introducing heavy data structures.
- Learn: left/right pointer movement, skipping characters, normalization logic.
- Goal: write linear-time scans without extra memory.

## 3. Stacks
- Folder: `problems/03_valid_parentheses`
- Challenge: `valid_parentheses`
- Difficulty: easy-medium
- Why here: introduces explicit state management through push/pop operations.
- Learn: stack discipline, delimiter matching, empty-state handling.
- Goal: model nested structure cleanly.

## 4. Binary Search
- Folder: `problems/04_binary_search`
- Challenge: `binary_search`
- Difficulty: easy-medium
- Why here: builds precision around loop bounds and invariants before harder recursive problems.
- Learn: mid calculation, loop bounds, off-by-one control.
- Goal: make interval invariants explicit and correct.

## 5. Linked Lists
- Folder: `problems/05_reverse_linked_list`
- Challenge: `reverse_linked_list`
- Difficulty: medium
- Why here: pointer mutation is a common interview stress point and worth isolating early.
- Learn: node traversal, pointer rewiring, nil safety.
- Goal: reason precisely about mutation and references.

## 6. Sliding Window
- Folder: `problems/06_longest_substring_without_repeating`
- Challenge: `longest_substring_without_repeating`
- Difficulty: medium
- Why here: combines maps with moving boundaries, which is a step up from fixed-pointer scans.
- Learn: dynamic windows, frequency/index maps, shrinking vs expanding logic.
- Goal: maintain state while moving through a sequence once.

## 7. Trees and DFS
- Folder: `problems/07_maximum_depth_of_binary_tree`
- Challenge: `maximum_depth_of_binary_tree`
- Difficulty: medium
- Why here: recursion is easier to absorb on a simple tree metric before graph traversal.
- Learn: recursive structure, base cases, depth-first search.
- Goal: map recursive definitions directly into code.

## 8. Graph Traversal
- Folder: `problems/08_number_of_islands`
- Challenge: `number_of_islands`
- Difficulty: medium-hard
- Why here: extends DFS thinking from trees to grids with visited-state management.
- Learn: connected components, neighbor expansion, mutation vs auxiliary memory.
- Goal: turn a grid into a traversal problem.

## 9. Heaps / Priority Queues
- Folder: `problems/09_top_k_frequent_elements`
- Challenge: `top_k_frequent_elements`
- Difficulty: medium-hard
- Why here: requires combining multiple ideas: counting, custom ordering, and complexity tradeoffs.
- Learn: heap interface, custom ordering, frequency aggregation.
- Goal: combine maps with heaps to control complexity.

## 10. Dynamic Programming
- Folder: `problems/10_coin_change`
- Challenge: `coin_change`
- Difficulty: hard relative to the rest of this track
- Why last: dynamic programming usually requires the strongest abstraction around states and transitions.
- Learn: states, transitions, initialization, impossible-state handling.
- Goal: move from recursive intuition to iterative tabulation.

## Recommended cadence
- Week 1: `01`-`03`
- Week 2: `04`-`06`
- Week 3: `07`-`08`
- Week 4: `09`-`10`, then repeat the set under time pressure

## Interview habits to build
- State time and space complexity before coding.
- Start with brute force, then reduce bottlenecks.
- Write down the invariant for pointer and binary-search problems.
- Use small table-top examples before implementing.
- Rerun `go test ./...` after every problem.
