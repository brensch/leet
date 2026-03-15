# Practice Suggestions

Based on the recent problems you worked through, the highest-value areas to keep pushing are:

## 1. Invariants Before Code

You improve noticeably once the invariant is explicit. Keep practicing saying the invariant out loud before touching the implementation.

Examples:

- `3Sum`: sorted array, `i` is fixed, and `left/right` move to adjust the sum while avoiding duplicate triplets.
- `Partition List`: every node belongs to exactly one of two stable sublists, and original order within each sublist is preserved.
- `Longest Consecutive Sequence`: only start counting at numbers that do not have a predecessor.

## 2. Counterexamples Before Trusting Tests

You should keep building the habit of trying one adversarial case yourself before trusting a passing test suite.

Examples:

- shared prerequisites in graph problems
- descendant-bound violations in BST problems
- duplicate-heavy inputs in set/two-pointer problems

## 3. Data Structure From Complexity Target

You are strongest when the structure is simple and matched to the operation you need.

Questions to ask first:

- “What needs to be fast?”
- “Is this really a presence-check problem?”
- “Do I need stable ordering?”
- “Am I tracking one state or multiple states?”

## Suggested Next Problems

- `17_3sum`: sorted two-pointer reasoning plus duplicate control
- `18_partition_list`: linked-list pointer manipulation with a stability requirement
- `19_longest_consecutive_sequence`: set-based reasoning and start-of-run detection
