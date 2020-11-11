package main

import (
	"math"
	"time"

	"github.com/tfriedel6/canvas"
)

type Widget interface {
	Paint(cv *canvas.Canvas, width float64, height float64)
	IsInside(x float64, y float64) bool
	WhenPressed()
}

type Point struct {
	x, y float64
}

type Rect struct {
	top           Point
	width, height float64
}

func (r *Rect) Paint(cv *canvas.Canvas, width float64, height float64) {

	color := math.Sin(float64(time.Now().UnixNano())*0.000000002)*0.3 + 0.7

	cv.SetFillStyle(color*0.2, color*0.2, color*0.8)
	cv.FillRect(r.top.x, r.top.y, r.width, r.height)
}

func (r *Rect) IsInside(x float64, y float64) {
	x1, y1 := r.top.x, r.top.y
	x2, y2 := r.top.x+r.width, r.top.y+r.height
	x3, y3 := r.top.x, r.top.y+r.height
	x4, y4 := r.top.x+r.width, r.top.y

}
