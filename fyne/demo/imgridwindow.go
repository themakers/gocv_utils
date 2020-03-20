package main

import (
	"context"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/themakers/gocv_utils/fyne/imgridwindow"
	"log"
	"runtime"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	runtime.LockOSThread()

	a := app.New()

	wWeight := imgridwindow.New(a, "Demo")
	wWeight.Show()
	wWeight.Resize(fyne.NewSize(800, 600))

	img := demoImage

	wWeight.AddImage(img)
	wWeight.SetImage(2, img)
	wWeight.SetImage(3, img)
	wWeight.SetImage(4, img)
	wWeight.AddImage(img)
	wWeight.SetImage(6, img)
	wWeight.SetImage(9, img)

	tt := wWeight.Trackbar().
		Name("TrackBar").
		Default(3).
		ValueType(imgridwindow.IntegerRange(10, 20)).
		Position(imgridwindow.PositionTop)


	chTrack := tt.Build()

	tt.ValueType(imgridwindow.IntegerRange(1000, 3000)).Name("2").Build()
	tt.Default(11).Name("3").Build()
	tt.Default(12).Name("4").Position(imgridwindow.PositionBottom).Build()

	wAbout := a.NewWindow("About")
	wAbout.Show()
	wAbout.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewCenterLayout(), widget.NewLabel("Demo"),
		),
	)

	go func() {
		for {
			select {
			case <-ctx.Done():
				a.Quit()
			case val := <-chTrack:
				log.Println("", val)
			}
		}
	}()

	wWeight.ShowAndRun()
}
