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

func StartApp() {
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
				if e.Type == touch.TypeBegin {
					for _, w := range Widgets {
						pressed := w.IsInside(float64(e.X), float64(e.Y))
						if pressed {
							w.OnPressed()
						}
					}

				} else if e.Type == touch.TypeEnd {
					for _, w := range Widgets {
						if w.GetPressed() {
							w.OffPressed()
						}
					}

				}
			}
		}
	})
}
