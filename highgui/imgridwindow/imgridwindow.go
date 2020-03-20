package imgridwindow

import (
	"github.com/themakers/gocv_utils/imgrid"
	"gocv.io/x/gocv"
	"image"
)

type Window struct {
	w *gocv.Window

	grid *imgrid.ImGrid

	width, height int
}

func New(name string, width, height int) *Window {
	w := &Window{
		w:      gocv.NewWindow(name),
		grid:   imgrid.New(),
		width:  width,
		height: height,
	}

	w.w.SetWindowTitle(name)

	return w
}

func (w *Window) Close() {
	w.grid.Close()

	if err := w.w.Close(); err != nil {
		panic(err)
	}
}

func (w *Window) CreateTrackbar(name string, min, max, pos int) func() int {
	t := w.w.CreateTrackbar(name, max)
	t.SetMin(min)
	t.SetMax(max)
	t.SetPos(pos)
	return t.GetPos
}

func (w *Window) WaitKey(d ...int) int {
	var ms int
	for _, d := range d {
		ms += d
	}
	return w.w.WaitKey(ms)
}

func (w *Window) AddImage(img gocv.Mat) {
	w.grid.AddImage(img)

	w.w.IMShow(w.grid.GenerateGridWithCellSize(image.Point{X: w.width, Y: w.height}))
	w.w.WaitKey(1)
}
