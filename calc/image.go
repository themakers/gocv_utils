package calc

import "image"

func calcFitRect1(dst, src image.Point) image.Rectangle {
	var (
		rt image.Rectangle

		srcAsp = float64(src.X) / float64(src.Y)
		dstAsp = float64(dst.X) / float64(dst.Y)
		asp2   = srcAsp / dstAsp
	)

	if asp2 > 1.0 {
		rt.Min.X = 0
		rt.Max.X = dst.X

		srcH := int(float64(dst.Y) / asp2)

		rt.Min.Y = (dst.Y - srcH) / 2
		rt.Max.Y = rt.Min.Y + srcH
	} else {
		rt.Min.Y = 0
		rt.Max.Y = dst.Y

		srcW := int(float64(dst.X) * asp2)

		rt.Min.X = (dst.X - srcW) / 2
		rt.Max.X = rt.Min.X + srcW
	}

	return rt
}

func CalcFitRect(dst, src image.Point) image.Rectangle {
	return RecalcFitRect(image.Rect(
		0, 0, dst.X, dst.Y,
	), src)
}

func RecalcFitRect(dst image.Rectangle, src image.Point) image.Rectangle {
	var (
		srcAsp = float64(src.X) / float64(src.Y)
		dstAsp = float64(dst.Size().X) / float64(dst.Size().Y)
		asp2   = srcAsp / dstAsp
	)

	if asp2 > 1.0 {
		srcH := int(float64(dst.Size().Y) / asp2)

		dst.Min.Y += (dst.Size().Y - srcH) / 2
		dst.Max.Y = dst.Min.Y + srcH
	} else {
		srcW := int(float64(dst.Size().X) * asp2)

		dst.Min.X += (dst.Size().X - srcW) / 2
		dst.Max.X = dst.Min.X + srcW
	}

	return dst
}
