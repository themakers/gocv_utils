package imgridwindow

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"github.com/themakers/gocv_utils/imgrid"
)

type ImGridWindow struct {
	fyne.Window

	zone struct {
		top    *fyne.Container
		bottom *fyne.Container
		left   *fyne.Container
		right  *fyne.Container
	}

	grid *imgrid.ImGrid
}

func New(app fyne.App, title string) *ImGridWindow {
	igw := &ImGridWindow{
		Window: app.NewWindow(title),
		grid:   imgrid.New(),
	}



	igw.zone.top = fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(1))
	igw.zone.bottom = fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(1))
	igw.zone.left = fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(1))
	igw.zone.right = fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(1))

	var (
		center = canvas.NewRaster(igw.raster)
	)

	igw.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(
		igw.zone.top, igw.zone.bottom, igw.zone.left, igw.zone.right,
	), igw.zone.top, igw.zone.bottom, igw.zone.left, igw.zone.right, center))

	return igw
}

func (igw *ImGridWindow) addWidget(w fyne.CanvasObject, pos Position) {
	var box *fyne.Container
	switch pos {
	case PositionTop:
		box = igw.zone.top
	case PositionBottom:
		box = igw.zone.bottom
	case PositionLeft:
		box = igw.zone.left
	case PositionRight:
		box = igw.zone.right
	}
	box.AddObject(w)
}
