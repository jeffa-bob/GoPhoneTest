package widgets

import (
	"github.com/tfriedel6/canvas"
)

type Button struct {
	Top                                   Point
	Text                                  string
	Width, Height, BorderRadius           float64
	color, Color, PressedColor, TextColor Color
	isPressed                             bool
}

func (r *Button) Paint(cv *canvas.Canvas, width float64, height float64) {

	cv.SetFillStyle(r.color.R, r.color.G, r.color.B, r.color.A)
	cv.FillRect(r.Top.X, r.Top.Y, r.Width, r.Height)
	cv.SetTextAlign(canvas.Center)
	cv.SetFillStyle(r.TextColor.R, r.TextColor.G, r.TextColor.B, r.TextColor.A)
	cv.FillText(r.Text, r.Top.X+(r.Width/2), r.Top.Y+(r.Height/2))
}

func (r *Button) IsInside(x float64, y float64) bool {
	x1, y1 := r.Top.X, r.Top.Y
	x2, y2 := r.Top.X+r.Width, r.Top.Y+r.Height

	if x > x1 && x < x2 && y > y1 && y < y2 {
		return true
	}
	return false
}

func (r *Button) OnPressed() {
	r.isPressed = true
	r.color = r.PressedColor
}
func (r *Button) OffPressed() {
	r.isPressed = false
	r.color = r.Color
}
