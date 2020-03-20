package imgrid

import (
	"github.com/themakers/gocv_utils/calc"
	"github.com/themakers/gocv_utils/imgop"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

type ImGrid struct {
	sz  image.Point
	asp float64

	grid []gocv.Mat
}

func New(sx, sy int) *ImGrid {
	img := &ImGrid{
		sz:  image.Point{X: sx, Y: sy},
		asp: float64(sx) / float64(sy),
	}

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
	{
		i, err := img.ToImage()
		if err != nil {
			panic(err)
		}

		m, err := gocv.ImageToMatRGBA(i)
		if err != nil {
			panic(err)
		}

		img = m
	}

	ig.grid[i] = img
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

func (ig *ImGrid) GenerateGrid() gocv.Mat {
	grid := calc.SimpleGridWithCellSize(len(ig.grid), ig.sz)

	canvas := imgop.NewMat(grid.GridSize)
	imgop.Fill(canvas, color.RGBA{R: 0, G: 0, B: 0, A: 0})

	grid.ForEach(func(col, row, cell int, cellRect image.Rectangle) {

		if ig.grid[cell].Ptr() == nil {
			return
		}

		imgop.BlitFit(canvas, cellRect, ig.grid[cell])

	})

	return canvas
}
