package imgrid

import (
	"github.com/themakers/gocv_utils/calc"
	"github.com/themakers/gocv_utils/imgop"
	"gocv.io/x/gocv"
	"image"
	"log"
)

const matType = gocv.MatTypeCV8UC3

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
	img = img.Clone()
	defer img.Close()

	dst := gocv.NewMatWithSize(ig.sz.X, ig.sz.Y, matType)

	dstRect := calc.CalcFitRect(imgop.MatSize(dst), imgop.MatSize(img))

	gocv.Resize(img, &img, dstRect.Size(), 0, 0, gocv.InterpolationLanczos4)

	// FIXME Hack; Should be gocv.CvtColor(img, &img, ???)
	{
		i, err := img.ToImage()
		if err != nil {
			panic(err)
		}

		m, err := gocv.ImageToMatRGB(i)
		if err != nil {
			panic(err)
		}

		img.Close()
		img = m

		defer img.Close()
	}


	log.Println("dstRect", imgop.MatSize(dst), imgop.MatSize(img), dstRect)
	{
		dstRegion := dst.Region(dstRect)
		img.CopyTo(&dstRegion)
	}

	ig.grid[i] = dst
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

	canvas := gocv.NewMatWithSize(grid.GridSize.Y, grid.GridSize.X, matType)

	grid.ForEach(func(col, row, cell int, cellRect image.Rectangle) {

		if ig.grid[cell].Ptr() == nil {
			return
		}

		dst := canvas.Region(cellRect)

		ig.grid[cell].CopyTo(&dst)

	})

	return canvas
}
