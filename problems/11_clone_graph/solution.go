package clone_graph

// Node is a graph node used by the Clone Graph problem.
type Node struct {
	Val       int
	Neighbors []*Node
}

// CloneGraph returns a deep copy of the input graph.
func CloneGraph(node *Node) *Node {

	// map old to new nodes
	visitedNodes := make(map[*Node]*Node)

	return copyNode(node, visitedNodes)
}

func copyNode(node *Node, visited map[*Node]*Node) *Node {

	if node == nil {
		return nil
	}

	visitedNode, ok := visited[node]
	if ok {
		return visitedNode
	}

	newNode := &Node{
		Val: node.Val,
	}
	visited[node] = newNode

	for _, node := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, copyNode(node, visited))
	}

	return newNode
}
