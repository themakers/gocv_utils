package imgridwindow

import (
"github.com/themakers/gocv_utils/imgrid"
"gocv.io/x/gocv"
)

type Window struct {
	w *gocv.Window

	grid *imgrid.ImGrid
}

func New(name string, width, height int) *Window {
	w := &Window{
		w:    gocv.NewWindow(name),
		grid: imgrid.New(width, height),
	}

	w.w.SetWindowTitle(name)

	return w
}

func (w *Window) Close() error {
	w.grid.Close()
	return w.w.Close()
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

	w.w.IMShow(w.grid.GenerateGrid())
	w.w.WaitKey(1)
}
