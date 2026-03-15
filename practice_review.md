# Practice Review

This file collects the problem snippets you shared in chat, along with concise comments on what was working and what needed attention.

## Container With Most Water

Problem: LeetCode 11

Your solution:

```go
func maxArea(height []int) int {

    max := 0

    for i := 0; i<len(height); i++ {
        for j := i; j<len(height); j++ {
            start := height[i]
            end := height[j]
            min := start
            if end < start {
                min = end
            }
            volume := (j-i)*min
            if volume > max {
                max = volume
            }
        }
    }

    return max
}
```

Comments:

- Correct brute-force area calculation.
- Time complexity is `O(n^2)`, which is not the expected strong interview answer.
- `j := i + 1` would avoid useless zero-width checks.
- Main next step was moving from brute force to a two-pointer invariant.

## 3Sum

Problem: LeetCode 15

Final strong version we converged to:

```go
import "slices"

func threeSum(nums []int) [][]int {
    slices.Sort(nums)

    var trips [][]int

    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        left := i + 1
        right := len(nums) - 1

        for left < right {
            sum := nums[i] + nums[left] + nums[right]

            if sum == 0 {
                trips = append(trips, []int{nums[i], nums[left], nums[right]})
                left++
                right--

                for left < right && nums[left] == nums[left-1] {
                    left++
                }
                for left < right && nums[right] == nums[right+1] {
                    right--
                }
            } else if sum < 0 {
                left++
            } else {
                right--
            }
        }
    }

    return trips
}
```

Comments:

- The key conceptual jump was from triple loops to `sort + fixed i + two pointers`.
- The main struggle was duplicate handling.
- The important rule was to skip duplicates at the source, not by scanning the result slice afterward.

## Set Matrix Zeroes

Problem: LeetCode 73

Your solution:

```go
func setZeroes(matrix [][]int) {

    columns := make(map[int]struct{})
    rows := make(map[int]struct{})

    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[i][j] == 0 {
                columns[i] = struct{}{}
                rows[j] = struct{}{}
            }
        }
    }

    for column := range columns {
        for j := range matrix[column] {
            matrix[column][j] = 0
        }
    }

    for row := range rows {
        for i := range matrix {
            matrix[i][row] = 0
        }
    }
}
```

Comments:

- Correct `O(mn)` solution using extra space.
- Good separation between discovery pass and mutation pass.
- `columns`/`rows` are named backwards relative to the indices being stored: `i` is a row, `j` is a column.
- Strong baseline answer unless the interviewer explicitly asks for `O(1)` extra space.

## Partition List

Problem: LeetCode 86

Your attempt:

```go
func partition(head *ListNode, x int) *ListNode {

    var lessThanHead *ListNode
    var greaterThanHead *ListNode

    var greaterThan *ListNode
    var lessThan *ListNode
    current := head
    for {
        if current.Val >= x {
            if greaterThanHead == nil {
                greaterThanHead = current
            } else {
                greaterThan.Next = current
            }
            greaterThan = current
        } else {
            if lessThanHead == nil {
                lessThanHead = current
            } else {
                lessThan.Next = current
            }
            lessThan = current
        }

        if current.Next == nil {
            break
        }
        current = current.Next
    }

    return greaterThanHead
}
```

Comments:

- Right high-level idea: build two stable lists.
- Main bugs:
- returning `greaterThanHead` instead of the head of the combined result
- never linking the `< x` list to the `>= x` list
- never terminating the tail, so old `Next` pointers can leak through
- missing `head == nil` safety

## Minimum Depth of Binary Tree

Problem: LeetCode 111

Your solution:

```go
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return findDepth(root, 0, 100000)
}

func findDepth(node *TreeNode, depth int, currentMin int) int {
    if node.Left == nil && node.Right == nil {
        return depth + 1
    }

    if depth > currentMin {
        return depth
    }

    left := 100000
    right := 100000
    if node.Left != nil {
        left = findDepth(node.Left, depth+1, currentMin)
    }

    if node.Right != nil {
        right = findDepth(node.Right, depth+1, left)
    }

    if left < right {
        return left
    } else {
        return right
    }
}
```

Comments:

- Probably accepted and conceptually close.
- More complex than necessary.
- `100000` is a magic number and weakens the explanation.
- The clean interview version handles missing-child cases directly instead of using sentinel values and pruning.

## Longest Consecutive Sequence

Problem: LeetCode 128

Your refined solution:

```go
func longestConsecutive(nums []int) int {

    numMap := make(map[int]struct{}, len(nums))
    for _, num := range nums {
        numMap[num] = struct{}{}
    }

    max := 0

    alreadyScanned := make(map[int]struct{})

    for _, num := range nums {
        _, scanned := alreadyScanned[num]
        if scanned {
            continue
        }

        _, prev := numMap[num-1]
        if prev {
            continue
        }

        total := 1
        num++

        for {
            _, ok := numMap[num]
            if !ok {
                break
            }
            alreadyScanned[num] = struct{}{}
            total++
            num++
        }

        if total > max {
            max = total
        }
    }
    return max
}
```

Comments:

- Core insight was correct: only start from numbers with no predecessor.
- Timeout happened because you iterated over the original slice, so duplicates could restart the same sequence.
- Cleaner approach is to iterate over the set itself and remove `alreadyScanned` entirely.

## Reverse Linked List

Problem: LeetCode 206

Your solution:

```go
func reverseList(head *ListNode) *ListNode {

    if head == nil {
        return nil
    }

    var prev *ListNode
    current := head
    next := head.Next

    for {
        next = current.Next
        current.Next = prev
        prev = current
        if next == nil {
            break
        }
        current = next
    }
    return current
}
```

Comments:

- Pointer updates were basically right.
- The bug was the return value: it should return `prev`, not `current`.
- `next := head.Next` before the loop is unnecessary.
- Cleaner loop shape is `for current != nil`.

## Contains Duplicate

Problem: LeetCode 217

Your solution:

```go
func containsDuplicate(nums []int) bool {

    tally := make(map[int]struct{})
    for _, num := range nums {
        _, ok := tally[num]
        if ok {
            return true
        }
        tally[num] = struct{}{}
    }
    return false
}
```

Comments:

- This is a strong standard answer.
- Correct, simple, and the right asymptotic complexity.
- Small optional improvement: pre-size the map with `len(nums)`.

## Validate Binary Search Tree

Problem: LeetCode 98

Your improved solution:

```go
import "math"

func IsValidBST(root *TreeNode) bool {
    return validInRange(root, math.MinInt, math.MaxInt)
}

func validInRange(root *TreeNode, min, max int) bool {
    if root == nil {
        return true
    }

    if root.Val >= max || root.Val <= min {
        return false
    }

    leftValid := validInRange(root.Left, min, root.Val)
    rightValid := validInRange(root.Right, root.Val, max)

    if !leftValid || !rightValid {
        return false
    }
    return true
}
```

Comments:

- Big improvement from local child checks to ancestor bounds.
- This is interview-correct in approach because the invariant is right.
- One subtle edge case remains: real node values equal to `math.MinInt` or `math.MaxInt` will be rejected incorrectly.
- The robust version uses pointer bounds or a wider numeric type.

## Min Depth / BST / Graph Pattern Notes

Recurring themes from the session:

- You improve quickly once the invariant is clear.
- A repeated weak spot was trusting passing tests before pressure-testing with counterexamples.
- Another repeated weak spot was using one “visited” concept where multiple states were needed.
- When the data-structure requirement is explicit, choose the structure from the complexity target first, then code.
