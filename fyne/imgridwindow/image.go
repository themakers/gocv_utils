package imgridwindow

import (
	"github.com/themakers/gocv_utils/imgop"
	"gocv.io/x/gocv"
	"image"
	"log"
	"time"
)

func (igw *ImGridWindow) raster(w, h int) image.Image {
	var t0 = time.Now()
	defer func() {
		log.Println("raster calculation time:", time.Now().Sub(t0))
	}()
	dst := imgop.NewMat(image.Point{X: w, Y: h})

	//grid := igw.grid.GenerateGridWithCellSize(image.Point{X: w, Y: h})
	grid := igw.grid.GenerateGridWithCellSize(image.Point{X: 400, Y: 300})

	imgop.BlitFit(dst, imgop.MatRect(dst), grid)

	img, err := dst.ToImage()
	if err != nil {
		panic(err)
	}
	return img
}

func (igw *ImGridWindow) AddImage(mat gocv.Mat) {
	igw.grid.AddImage(mat)
}


func (igw *ImGridWindow) SetImage(i int, mat gocv.Mat) {
	igw.grid.SetImage(i, mat)
}
