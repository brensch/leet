package number_of_islands

// NumIslands counts connected components of land in the grid.
func NumIslands(grid [][]byte) int {

	islands := 0
	// iterate over all the blocks, once we find one, traverse in all directions and remove from unvisited blocks
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				// record we found an island
				islands++
				// set all attached blocks to 0
				traverseIsland(grid, i, j)
			}
		}
	}

	return islands
}

// traverseIsland sets all points connected to a block to 0
func traverseIsland(grid [][]byte, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) || grid[x][y] != '1' {
		return
	}

	grid[x][y] = '0'

	traverseIsland(grid, x+1, y)
	traverseIsland(grid, x-1, y)
	traverseIsland(grid, x, y+1)
	traverseIsland(grid, x, y-1)
}
