package imgridwindow

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/themakers/gocv_utils/utils/debouncer"
	"log"
	"strings"
	"time"
)

type Position int

const (
	PositionTop Position = iota + 1
	PositionBottom
	PositionLeft
	PositionRight
)

func (igw *ImGridWindow) Trackbar() trackbarBuilder {
	return trackbarBuilder{igw: igw}
}

func (igw *ImGridWindow) Group(name string, pos Position) *widget.Group {
	grp := widget.NewGroup(name)
	igw.addWidget(grp, pos)
	return grp
}

type trackbarBuilder struct {
	name string
	def  float64

	vt Value

	pos Position

	group *widget.Group

	igw *ImGridWindow
}

func (tb trackbarBuilder) Name(name string) trackbarBuilder {
	tb.name = name
	return tb
}

func (tb trackbarBuilder) ValueType(vt Value) trackbarBuilder {
	tb.vt = vt
	return tb
}

func (tb trackbarBuilder) Default(def float64) trackbarBuilder {
	tb.def = def
	return tb
}

func (tb trackbarBuilder) Group(group *widget.Group) trackbarBuilder {
	tb.group = group
	return tb
}

func (tb trackbarBuilder) Position(pos Position) trackbarBuilder {
	tb.pos = pos
	return tb
}

func (tb trackbarBuilder) Build() <-chan float64 {
	{
		if tb.vt == nil {
			tb.vt = IntegerRange(0, 100)
		}

		if tb.pos == 0 {
			tb.pos = PositionTop
		}
	}

	log.Printf("%#v", tb)

	max, fin, fout := tb.vt()

	if tb.def < fout(0) {
		tb.def = fout(0)
	} else if tb.def > fout(max) {
		tb.def = fout(max)
	}

	track := widget.NewSlider(0, max)

	track.Min = 0
	track.Max = max
	track.Step = 1
	track.Value = fin(tb.def)

	ch := make(chan float64)

	var label = widget.NewLabel("")

	deb := debouncer.New(500 * time.Millisecond)

	format := func(v float64) {
		label.Text = strings.TrimSpace(fmt.Sprintf("(%v) %s", fout(v), tb.name))
		label.Refresh()
	}

	format(fin(tb.def))

	track.OnChanged = (func() func(float64) {
		lastVal := track.Value

		return func(v float64) {
			format(v)

			deb.Trigger(func() {
				format(v)

				go func() {
					if v != lastVal {
						lastVal = v
						ch <- fout(v)
					}
				}()
			})
		}
	})()

	box := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(nil, nil, nil, label),
		label, track,
	)

	if tb.group != nil {
		tb.group.Append(box)
	} else {
		tb.igw.addWidget(box, tb.pos)
	}

	return ch
}

type Filter func(v float64) float64

type Value func() (max float64, in, out Filter)

//func IntegerRange2(min, max int) Value {
//	return func() (float64, Filter, Filter) {
//		return float64(max - min), func(v float64) float64 {
//				return v - float64(min)
//			}, func(v float64) float64 {
//				return v + float64(min)
//			}
//	}
//}

func IntegerRange(min, max int) Value {
	return func() (float64, Filter, Filter) {
		min := float64(min)
		max := float64(max)

		return max - min, func(v float64) float64 {
			return v - min
		}, func(v float64) float64 {
			return v + min
		}
	}
}
