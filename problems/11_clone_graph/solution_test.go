package clone_graph

import "testing"

func linkUndirected(a, b *Node) {
	a.Neighbors = append(a.Neighbors, b)
	b.Neighbors = append(b.Neighbors, a)
}

func graphSignature(node *Node) map[int][]int {
	if node == nil {
		return nil
	}

	signature := make(map[int][]int)
	visited := map[*Node]bool{}
	queue := []*Node{node}
	visited[node] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range current.Neighbors {
			signature[current.Val] = append(signature[current.Val], neighbor.Val)
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return signature
}

func TestCloneGraph(t *testing.T) {

	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	linkUndirected(node1, node2)
	linkUndirected(node1, node4)
	linkUndirected(node2, node3)
	linkUndirected(node3, node4)

	tests := []struct {
		name string
		in   *Node
	}{
		{name: "square graph", in: node1},
		{name: "nil graph", in: nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CloneGraph(tc.in)

			if tc.in == nil {
				if got != nil {
					t.Fatalf("CloneGraph(nil) = %#v, want nil", got)
				}
				return
			}

			if got == tc.in {
				t.Fatalf("CloneGraph returned the original node pointer")
			}

			if got.Val != tc.in.Val {
				t.Fatalf("CloneGraph root value = %d, want %d", got.Val, tc.in.Val)
			}

			gotSignature := graphSignature(got)
			wantSignature := graphSignature(tc.in)
			if len(gotSignature) != len(wantSignature) {
				t.Fatalf("cloned graph has %d nodes, want %d", len(gotSignature), len(wantSignature))
			}

			for nodeVal, wantNeighbors := range wantSignature {
				gotNeighbors := gotSignature[nodeVal]
				if len(gotNeighbors) != len(wantNeighbors) {
					t.Fatalf("node %d has neighbors %v, want %v", nodeVal, gotNeighbors, wantNeighbors)
				}
				for i := range wantNeighbors {
					if gotNeighbors[i] != wantNeighbors[i] {
						t.Fatalf("node %d has neighbors %v, want %v", nodeVal, gotNeighbors, wantNeighbors)
					}
				}
			}
		})
	}
}
