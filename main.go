package GoPhoneTest

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"log"

	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/xmobilebackend"

	//"golang.org/x/mobile/exp/f32"
	"GoPhoneTest/widgets"
	"golang.org/x/mobile/gl"
)

var Widgets []widgets.Widget

func AddWidget() {
	rect := widgets.Rect{Top: widgets.Point{100, 100}, Width: 1000, Height: 700, Color: widgets.Color{R: 10, G: 10, B: 200, A: 0}, Fill: true}
	Widgets = append(Widgets, &rect)
	but := widgets.Button{
		Top:          widgets.Point{400, 400},
		Text:         "BUTTON",
		Width:        300,
		Height:       300,
		BorderRadius: 0,
		Color:        widgets.Color{0, 255, 0, 0},
		PressedColor: widgets.Color{255, 0, 0, 0},
		TextColor:    widgets.Color{0, 0, 255, 0},
	}
	Widgets = append(Widgets, &but)
}

func main() {
	AddWidget()
	app.Main(func(a app.App) {
		var cv, painter *canvas.Canvas
		var cvb *xmobilebackend.XMobileBackendOffscreen
		var painterb *xmobilebackend.XMobileBackend
		var w, h int

		var glctx gl.Context
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					var err error
					glctx = e.DrawContext.(gl.Context)
					ctx, err := xmobilebackend.NewGLContext(glctx)
					if err != nil {
						log.Fatal(err)
					}
					cvb, err = xmobilebackend.NewOffscreen(0, 0, false, ctx)
					if err != nil {
						log.Fatalln(err)
					}
					painterb, err = xmobilebackend.New(0, 0, 0, 0, ctx)
					if err != nil {
						log.Fatalln(err)
					}
					cv = canvas.New(cvb)
					painter = canvas.New(painterb)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					cvb.Delete()
					glctx = nil
				}
			case size.Event:
				w, h = e.WidthPx, e.HeightPx
			case paint.Event:
				if glctx != nil {
					cvb.SetSize(w, h)

					fw, fh := float64(w), float64(h)

					for _, w := range Widgets {
						w.Paint(cv, fw, fh)
					}

					painterb.SetBounds(0, 0, w, h)
					painter.DrawImage(cv)

					a.Publish()
					a.Send(paint.Event{})
				}
			case touch.Event:

			}
		}
	})
}
