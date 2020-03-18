package imgridwindow

import (
	"gocv.io/x/gocv"
	"image"
	"math"
)

const matType = gocv.MatTypeCV8UC3

type Window struct {
	w      *gocv.Window
	width  int
	height int

	imgs   []gocv.Mat
	canvas gocv.Mat
}

func New(name string, width, height int) *Window {
	w := &Window{
		w:      gocv.NewWindow(name),
		width:  width,
		height: height,
	}

	w.w.SetWindowTitle(name)

	//w.w.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)
	//w.w.SetWindowProperty(gocv.WindowPropertyAspectRatio, gocv.WindowFreeRatio)

	return w
}

func (w *Window) Close() error {
	for _, mat := range w.imgs {
		if err := mat.Close(); err != nil {
			return err
		}
	}
	return w.w.Close()
}

func (w *Window) WaitKey(d int) {
	w.w.WaitKey(d)
}

func (w *Window) AddImage(img gocv.Mat) {
	mat := gocv.NewMatWithSize(w.height, w.width, matType)

	dstRect := w.calcDstRect(img)
	dst := mat.Region(dstRect)

	img = img.Clone()
	defer img.Close()

	gocv.Resize(img, &img, dstRect.Size(), 0, 0, gocv.InterpolationLanczos4)
	img.ConvertTo(&img, matType)
	img.CopyTo(&dst)

	w.imgs = append(w.imgs, mat)
	w.render()
}

func (w *Window) calcDstRect(img gocv.Mat) image.Rectangle {
	var (
		rt = image.Rect(0, 0, w.width, w.height)

		size = img.Size()
		iw   = size[1]
		ih   = size[0]

		iAsp = float64(iw) / float64(ih)
		wAsp = float64(w.width) / float64(w.height)
		asp2 = iAsp / wAsp
	)

	if asp2 > 1.0 {
		rt.Min.X = 0
		rt.Max.X = w.width

		ih := int(float64(w.height) / asp2)

		rt.Min.Y = (w.height - ih) / 2
		rt.Max.Y = rt.Min.Y + ih
	} else {
		rt.Min.Y = 0
		rt.Max.Y = w.height

		iw := int(float64(w.width) / asp2)

		rt.Min.X = (w.width - iw) / 2
		rt.Max.X = rt.Min.X + iw
	}

	return rt
}

func (w *Window) render() {
	var (
		nImg     = len(w.imgs)
		gridRows = int(math.Sqrt(float64(nImg)))
		gridCols = nImg / gridRows
	)

	if nImg%gridRows > 0 {
		gridCols++
	}

	var (
		cWidth  = gridCols * w.width
		cHeight = gridRows * w.height
	)

	if w.canvas.Ptr() != nil {
		w.canvas.Close()
	}

	w.canvas = gocv.NewMatWithSize(cHeight, cWidth, matType)

	var ii = 0
	for row := 0; row < gridRows; row++ {
		for col := 0; col < gridCols && ii < nImg; col++ {
			rt := image.Rect(
				col*w.width, row*w.height,
				col*w.width+w.width, row*w.height+w.height,
			)

			dst := w.canvas.Region(rt)
			w.imgs[ii].CopyTo(&dst)

			ii++
		}
	}

	w.w.ResizeWindow(cWidth/2, cHeight/2)

	w.w.IMShow(w.canvas)
}
