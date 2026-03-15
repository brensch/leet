package find_median_from_data_stream

import "fmt"

// MedianFinder stores streamed integers and reports the median.
type MedianFinder struct {
	nodes int
	head  *node
}

type node struct {
	val  int
	next *node
}

// Constructor initializes a MedianFinder.
func Constructor() MedianFinder {
	// don't think i need a map here, may be wrong
	return MedianFinder{}
}

// AddNum inserts a number into the data structure.
func (m *MedianFinder) AddNum(num int) {
	m.nodes++
	fmt.Println("node count", m.nodes)
	start := m.head
	newNode := &node{
		val: num,
	}
	if start == nil {
		fmt.Println("setting head")
		m.head = newNode
		return
	}

	// if this is the first node set it to head early
	if start.next != nil && start.next.val >= newNode.val {
		m.head = newNode
		newNode.next = m.head
		fmt.Println("adding first")
		return
	}

	currentNode := start
	for {
		nextNode := currentNode.next
		if nextNode == nil {
			currentNode.next = newNode
			return
		}

		if nextNode.val >= num {
			fmt.Println("adding other", newNode.val)
			newNode.next = nextNode
			currentNode.next = newNode
			return

		}
		currentNode = nextNode
	}

}

// FindMedian returns the median of all inserted numbers.
func (m *MedianFinder) FindMedian() float64 {

	fmt.Println("nodes:")
	printer := m.head
	for {
		fmt.Println(printer.val)
		printer = printer.next
		if printer == nil {
			break
		}
	}

	currentNode := m.head

	// should never get to the end since counting half

	stepsToTraverse := m.nodes / 2
	remainder := m.nodes % 2

	fmt.Println("remainder", remainder, "steps", stepsToTraverse)

	for i := 0; i < stepsToTraverse-1+remainder; i++ {
		currentNode = currentNode.next
	}

	fmt.Println("stopped at node", currentNode.val)

	if remainder == 0 {
		fmt.Println("doing remainder")
		return (float64(currentNode.val) + float64(currentNode.next.val)) / 2
	}

	return float64(currentNode.val)
}
