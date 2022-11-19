package service

// represents a cell type with a row and a col
type Cell struct {
	Row int
	Col int
}

func GetCell(row, col int) Cell {
	return Cell{row, col}
}

// is the cell in a list of cells
func (c *Cell) CellInList(cells []Cell) bool {
	for _, cls := range cells {
		if cls.Row == c.Row && cls.Col == c.Col {
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
	if c.Row > 0 {
		if c.Col > 0 {
			cells = append(cells, GetCell(c.Row-1, c.Col-1))
		}

		cells = append(cells, GetCell(c.Row-1, c.Col))

		if c.Col < maxCol {
			cells = append(cells, GetCell(c.Row-1, c.Col+1))
		}
	}

	// get the one on the left (4)
	if c.Col > 0 {
		cells = append(cells, GetCell(c.Row, c.Col-1))
	}

	// get the one on the right (5)
	if c.Col < maxCol {
		cells = append(cells, GetCell(c.Row, c.Col+1))
	}

	// Get all the cells below X (6, 7, 8)
	if c.Row < maxRow {
		if c.Col > 0 {
			cells = append(cells, GetCell(c.Row+1, c.Col-1))
		}

		cells = append(cells, GetCell(c.Row+1, c.Col))

		if c.Col < maxCol {
			cells = append(cells, GetCell(c.Row+1, c.Col+1))
		}
	}

	return cells

}
