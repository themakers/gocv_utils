package calc

import (
	"image"
	"math"
)

type FixedGrid struct {
	NumCells int
	NumCols  int
	NumRows  int

	GridSize image.Point
	CellSize image.Point
}

func (g FixedGrid) ForEach(fn func(col, row, cell int, cellRect image.Rectangle)) {
	for row, ii := 0, 0; row < g.NumRows; row++ {
		for col := 0; col < g.NumCols && ii < g.NumCells; col, ii = col+1, ii+1 {

			cellRect := image.Rect(
				col*g.CellSize.X, row*g.CellSize.Y,
				col*g.CellSize.X+g.CellSize.X, row*g.CellSize.Y+g.CellSize.Y,
			)

			fn(col, row, ii, cellRect)
		}
	}

}

func SimpleGridWithCellSize(numCells int, cellSize image.Point) FixedGrid {
	var (
		numRows = int(math.Sqrt(float64(numCells)))
		numCols = numCells / numRows
	)

	if numCells%numRows > 0 {
		numCols++
	}

	return FixedGrid{
		NumCells: numCells,
		NumCols:  numCols,
		NumRows:  numRows,
		GridSize: image.Point{
			X: numCols * cellSize.X,
			Y: numRows * cellSize.Y,
		},
		CellSize: cellSize,
	}
}

func GridWithSize(numCells int, gridSize image.Point) (numRows, numCols int, cellSize image.Point) {
	return
}
