package imgridwindow

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/themakers/gocv_utils/imgop"
	"github.com/themakers/gocv_utils/imgrid"
	"gocv.io/x/gocv"
	"image"
	"log"
	"time"
)

type ImGridWindow struct {
	fyne.Window
	grid *imgrid.ImGrid
}

func New(app fyne.App, title string) *ImGridWindow {
	igw := &ImGridWindow{
		Window: app.NewWindow(title),
		grid:   imgrid.New(400, 400),
	}

	var (
		top    = widget.NewLabel("TOP")
		bottom = widget.NewLabel("BOTTOM")
		left = widget.NewLabel("")
		right = widget.NewLabel("")
		center = canvas.NewRaster(igw.raster)
	)

	igw.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(
		top, bottom, left, right,
	), top, bottom, left, right, center))

	return igw
}

func (igw *ImGridWindow) raster(w, h int) image.Image {
	var t0 = time.Now()
	defer func() {
		log.Println("raster calculation time2:", time.Now().Sub(t0))
	}()
	dst := imgop.NewMat(image.Point{X: w, Y: h})

	grid := igw.grid.GenerateGrid()

	imgop.BlitFit(dst, imgop.MatRect(dst), grid)
	log.Println("raster calculation time1:", time.Now().Sub(t0))

	img, err := dst.ToImage()
	if err != nil {
		panic(err)
	}
	return img
}

func (igw *ImGridWindow) AddImage(mat gocv.Mat) {
	igw.grid.AddImage(mat)
}
