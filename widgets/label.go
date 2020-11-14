package widgets

import (
	"github.com/tfriedel6/canvas"
)

type Label struct {
	Top                                           Point
	Text                                          string
	Width, Height, BorderRadius                   float64
	Color, originalColor, PressedColor, TextColor Color
	isPressed                                     bool
}

func (r *Label) Paint(cv *canvas.Canvas, width float64, height float64) {

	cv.SetFillStyle(r.Color.R, r.Color.G, r.Color.B)
	cv.FillRect(r.Top.X, r.Top.Y, r.Width, r.Height)
	cv.SetTextAlign(canvas.Center)
	cv.SetStrokeStyle(r.TextColor.R, r.TextColor.G, r.TextColor.B)
	cv.SetLineWidth(4)
	cv.SetFont("/system/fonts/Roboto-Thin.ttf", 34)
	cv.StrokeText(r.Text, r.Top.X+(r.Width/2), r.Top.Y+(r.Height/2))
}

func (r *Label) IsInside(x float64, y float64) bool {
	x1, y1 := r.Top.X, r.Top.Y
	x2, y2 := r.Top.X+r.Width, r.Top.Y+r.Height

	if x > x1 && x < x2 && y > y1 && y < y2 {
		return true
	}
	return false
}

func (r *Label) OnPressed() {
	r.isPressed = true
	r.originalColor = r.Color
	r.Color = r.PressedColor
}
func (r *Label) OffPressed() {
	r.isPressed = false
	r.Color = r.originalColor
}
func (r *Label) GetPressed() bool {
	return r.isPressed
}
