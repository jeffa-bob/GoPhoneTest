package widgets

import (
	"github.com/tfriedel6/canvas"
)

type Widget interface {
	Paint(cv *canvas.Canvas, width float64, height float64)
	IsInside(x float64, y float64) bool
	OnPressed()
	OffPressed()
}

type Point struct {
	X, Y float64
}

type Color struct {
	R, G, B, A uint8
}
