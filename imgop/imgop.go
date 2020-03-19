package imgop

import (
	"gocv.io/x/gocv"
	"image"
)

func MatSize(mat gocv.Mat) image.Point {
	var sz = mat.Size()
	return image.Point{
		X: sz[1],
		Y: sz[0],
	}
}
