package widgets

import (
	"github.com/tfriedel6/canvas"
	"math"
	"time"
)

type Rect struct {
	Top           Point
	Width, Height float64
	Color         Color
	Fill          bool
	isPressed     bool
}

func (r *Rect) Paint(cv *canvas.Canvas, width float64, height float64) {

	color := math.Sin(float64(time.Now().UnixNano())*0.000000002)*0.3 + 0.7

	cv.SetFillStyle(color*(float64(r.Color.R)/255), color*(float64(r.Color.G)/255), color*(float64(r.Color.B)/255))
	if r.Fill {
		cv.FillRect(r.Top.X, r.Top.Y, r.Width, r.Height)
	} else {
		cv.ClearRect(r.Top.X, r.Top.Y, r.Width, r.Height)
	}
}

func (r *Rect) IsInside(x float64, y float64) bool {
	x1, y1 := r.Top.X, r.Top.Y
	x2, y2 := r.Top.X+r.Width, r.Top.Y+r.Height

	if x > x1 && x < x2 && y > y1 && y < y2 {
		return true
	}
	return false
}

func (r *Rect) OnPressed() {
	r.isPressed = true
}
func (r *Rect) OffPressed() {
	r.isPressed = false
}
