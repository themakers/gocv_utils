package imgop

import (
	"errors"
	"fmt"
	"github.com/themakers/gocv_utils/calc"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

const MatType = gocv.MatTypeCV8UC4

func NewMat(size image.Point) gocv.Mat {
	return gocv.NewMatWithSize(size.Y, size.X, MatType)

	// FIXME ???
	//mat, err := gocv.NewMatFromBytes(
	//	size.Y, size.X, MatType,
	//	bytes.Repeat([]byte{0, 0, 0, 128}, size.Y*size.X*4),
	//)
	//if err != nil {
	//	panic(err)
	//}
	//return mat
}

func MatSize(mat gocv.Mat) image.Point {
	var sz = mat.Size()
	return image.Point{
		X: sz[1],
		Y: sz[0],
	}
}

func MatRect(mat gocv.Mat) image.Rectangle {
	var sz = MatSize(mat)
	return image.Rect(0, 0, sz.X, sz.Y)
}

func Fill(mat gocv.Mat, c color.RGBA) {
	size := MatSize(mat)
	gocv.FillPoly(&mat, [][]image.Point{{
		{X: 0, Y: 0},
		{X: size.X, Y: 0},
		{X: size.X, Y: size.Y},
		{X: 0, Y: size.Y},
		{X: 0, Y: 0},
	}}, c)
}

func LoadFile(name string) gocv.Mat {
	return gocv.IMRead(name, gocv.IMReadUnchanged)
}

func LoadBytes(r io.Reader) gocv.Mat {
	img, _, err := image.Decode(r)
	if err != nil {
		panic(err)
	}
	mat, err := gocv.ImageToMatRGBA(img)
	if err != nil {
		panic(err)
	}
	return mat
}

//func Resize(mat gocv.Mat, size image.Point) gocv.Mat {
//	dst := gocv.NewMatWithSize(size.X, size.Y, matType)
//	gocv.Resize(mat, &dst, size, 0, 0, gocv.InterpolationLinear)
//	return dst
//}

//func CopyFit(dst, src gocv.Mat) {
//	dstRect := calc.CalcFitRect(MatSize(dst), MatSize(src))
//
//	Resize(src, dstRect.Size())
//
//	dst = dst.Region(dstRect)
//
//	src.CopyTo(&dst)
//}
//
//func copyFitOnCanvas1(src, canvas gocv.Mat, rect image.Rectangle) {
//	dst := canvas.Region(rect)
//	CopyFit(dst, src)
//}

func BlitFit(dst gocv.Mat, rect image.Rectangle, src gocv.Mat) {
	rect = calc.RecalcFitRect(rect, MatSize(src))

	region := dst.Region(rect)

	gocv.Resize(src, &region, rect.Size(), 0, 0, gocv.InterpolationLinear)
}

// Reference:
// https://docs.opencv.org/4.2.0/d8/d01/group__imgproc__color__conversions.html
func ToRGBA(src gocv.Mat) gocv.Mat {
	switch src.Type() {
	case gocv.MatTypeCV8UC1:
		dst := NewMat(MatSize(src))
		gocv.CvtColor(src, &dst, gocv.ColorGrayToBGRA)
		return dst
	case gocv.MatTypeCV8UC3:
		dst := NewMat(MatSize(src))
		gocv.CvtColor(src, &dst, gocv.ColorBGRToBGRA)
		return dst
	case gocv.MatTypeCV8UC4:
		return src.Clone()
	default:
		panic(errors.New(fmt.Sprintf("unsupported mat type: %v", src.Type())))
	}
}
