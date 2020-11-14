package widgets

import (
	"github.com/tfriedel6/canvas"
	"math"
)

type RectButton struct {
	Top                                           Point
	Text                                          string
	Width, Height, BorderRadius                   float64
	Color, originalColor, PressedColor, TextColor Color
	isPressed                                     bool
	OnClick                                       []func(button *RectButton)
	OffClick                                      []func(button *RectButton)
}

func (r *RectButton) Paint(cv *canvas.Canvas, width float64, height float64) {
	cv.SetFillStyle(r.Color.R, r.Color.G, r.Color.B)
	cv.FillRect(r.Top.X, r.Top.Y, r.Width, r.Height)
	cv.SetTextAlign(canvas.Center)
	cv.SetStrokeStyle(r.TextColor.R, r.TextColor.G, r.TextColor.B)
	cv.SetLineWidth(4)
	cv.SetFont("/system/fonts/Roboto-Thin.ttf", 34)
	cv.StrokeText(r.Text, r.Top.X+(r.Width/2), r.Top.Y+(r.Height/2))
}

func (r *RectButton) IsInside(x float64, y float64) bool {
	x1, y1 := r.Top.X, r.Top.Y
	x2, y2 := r.Top.X+r.Width, r.Top.Y+r.Height

	if x > x1 && x < x2 && y > y1 && y < y2 {
		return true
	}
	return false
}

func (r *RectButton) OnPressed() {
	r.isPressed = true
	r.originalColor = r.Color
	r.Color = r.PressedColor
	for _, f := range r.OnClick {
		f(r)
	}
}
func (r *RectButton) OffPressed() {
	r.isPressed = false
	r.Color = r.originalColor
	for _, f := range r.OffClick {
		f(r)
	}
}
func (r *RectButton) GetPressed() bool {
	return r.isPressed
}

type CircleButton struct {
	Position                                      Point
	Text                                          string
	Radius                                        float64
	Color, originalColor, PressedColor, TextColor Color
	isPressed                                     bool
	OnClick                                       []func(button *CircleButton)
	OffClick                                      []func(button *CircleButton)
}

func (r *CircleButton) Paint(cv *canvas.Canvas, width float64, height float64) {
	cv.BeginPath()
	cv.SetStrokeStyle(r.Color.R, r.Color.G, r.Color.B)
	cv.Arc(r.Position.X, r.Position.Y, r.Radius, 0, 2*math.Pi, false)
	cv.SetFillStyle(r.Color.R, r.Color.G, r.Color.B)
	cv.Fill()
	cv.SetLineWidth(0)
	cv.Stroke()
	cv.SetTextAlign(canvas.Center)
	cv.SetStrokeStyle(r.TextColor.R, r.TextColor.G, r.TextColor.B)
	cv.SetLineWidth(4)
	cv.SetFont("/system/fonts/Roboto-Thin.ttf", (r.Radius * 1.75))
	cv.StrokeText(r.Text, r.Position.X, r.Position.Y+((r.Radius/2)*1.25))
}

func (r *CircleButton) IsInside(x float64, y float64) bool {
	dx := math.Abs(x - r.Position.X)
	dy := math.Abs(y - r.Position.Y)
	R := r.Radius

	if dy > R {
		return false
	}
	if dx > R {
		return false
	}

	if dx+dy <= R {
		return true
	}

	if dx*dx+dy*dy <= R*R {
		return true
	} else {
		return false
	}

}

func (r *CircleButton) OnPressed() {
	r.isPressed = true
	r.originalColor = r.Color
	r.Color = r.PressedColor
	for _, f := range r.OnClick {
		f(r)
	}
}
func (r *CircleButton) OffPressed() {
	r.isPressed = false
	r.Color = r.originalColor
	for _, f := range r.OffClick {
		f(r)
	}
}
func (r *CircleButton) GetPressed() bool {
	return r.isPressed
}
