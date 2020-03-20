package imgrid

import (
	"github.com/themakers/gocv_utils/calc"
	"github.com/themakers/gocv_utils/imgop"
	"gocv.io/x/gocv"
	"image"
)

type ImGrid struct {
	grid []gocv.Mat
}

func New() *ImGrid {
	img := &ImGrid{}

	return img
}

func (ig *ImGrid) Close() {
	for _, img := range ig.grid {
		if img.Ptr() != nil {
			if err := img.Close(); err != nil {
				panic(err)
			}
		}
	}
}

func (ig *ImGrid) setImage(i int, img gocv.Mat) {
	// FIXME Hack; Should be gocv.CvtColor(img, &img, ???)
	if ig.grid[i].Ptr() != nil {
		ig.grid[i].Close()
	}
	ig.grid[i] = imgop.ToRGBA(img)
}

func (ig *ImGrid) SetImage(i int, img gocv.Mat) {
	if i == len(ig.grid) {
		ig.grid = append(ig.grid, gocv.Mat{})
	} else if i >= len(ig.grid) {
		ig.grid = append(ig.grid, make([]gocv.Mat, (i+1)-len(ig.grid))...)
	}

	ig.setImage(i, img)
}

func (ig *ImGrid) AddImage(img gocv.Mat) {
	ig.SetImage(len(ig.grid), img)
}

//func (ig *ImGrid) GenerateGridWithSize(gridSize image.Point) gocv.Mat {
//	grid := calc.GridWithSize(len(ig.grid), 1.0, gridSize)
//
//	canvas := imgop.NewMat(grid.GridSize)
//
//	grid.ForEach(func(col, row, cell int, cellRect image.Rectangle) {
//
//		if ig.grid[cell].Ptr() == nil {
//			return
//		}
//
//		imgop.BlitFit(canvas, cellRect, ig.grid[cell])
//
//	})
//
//	return canvas
//}


func (ig *ImGrid) GenerateGridWithCellSize(cellSize image.Point) gocv.Mat {
	grid := calc.SimpleGridWithCellSize(len(ig.grid), cellSize)

	canvas := imgop.NewMat(grid.GridSize)
	// FIXME ???
	// imgop.Fill(canvas, color.RGBA{R: 0, G: 0, B: 0, A: 255})

	grid.ForEach(func(col, row, cell int, cellRect image.Rectangle) {

		if ig.grid[cell].Ptr() == nil {
			return
		}

		imgop.BlitFit(canvas, cellRect, ig.grid[cell])

	})

	return canvas
}
