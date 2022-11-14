package main

// represents a cell type with a row and a col
type Cell struct {
	row int
	col int
}

func GetCell(row, col int) Cell {
	return Cell{row, col}
}

// is the cell in a list of cells
func (c *Cell) CellInList(cells []Cell) bool {
	for _, cls := range cells {
		if cls.row == c.row && cls.col == c.col {
			return true
		}
	}

	return false
}

func (c *Cell) GetNeighbors(maxRow, maxCol int) []Cell {
	cells := []Cell{}

	// returns a list of cells around the cell X
	// 1 2 3
	// 4 X 5
	// 6 7 8

	// Get all the cells above X (1, 2, 3)
	if c.row > 0 {
		if c.col > 0 {
			cells = append(cells, GetCell(c.row-1, c.col-1))
		}

		cells = append(cells, GetCell(c.row-1, c.col))

		if c.col < maxCol {
			cells = append(cells, GetCell(c.row-1, c.col+1))
		}
	}

	// get the one on the left (4)
	if c.col > 0 {
		cells = append(cells, GetCell(c.row, c.col-1))
	}

	// get the one on the right (5)
	if c.col < maxCol {
		cells = append(cells, GetCell(c.row, c.col+1))
	}

	// Get all the cells below X (6, 7, 8)
	if c.row < maxRow {
		if c.col > 0 {
			cells = append(cells, GetCell(c.row+1, c.col-1))
		}

		cells = append(cells, GetCell(c.row+1, c.col))

		if c.col < maxCol {
			cells = append(cells, GetCell(c.row+1, c.col+1))
		}
	}

	return cells

}
