package main

import (
	"GoPhoneTest"
	"GoPhoneTest/widgets"
)

func AddWidget() {
	rect := widgets.Rect{Top: widgets.Point{X: 50, Y: 50}, Width: 100, Height: 100, Color: widgets.Color{R: 10, G: 10, B: 200, A: 0}, Fill: true}
	GoPhoneTest.Widgets = append(GoPhoneTest.Widgets, &rect)
	but := widgets.RectButton{
		Top:          widgets.Point{X: 300, Y: 300},
		Text:         "BUTTON",
		Width:        300,
		Height:       100,
		BorderRadius: 0,
		Color:        widgets.Color{G: 255, A: 255},
		PressedColor: widgets.Color{G: 100, A: 255},
		TextColor:    widgets.Color{B: 255, A: 255},
	}
	GoPhoneTest.Widgets = append(GoPhoneTest.Widgets, &but)
	cir := widgets.CircleButton{
		Position:     widgets.Point{X: 500, Y: 150},
		Text:         "1",
		Radius:       30,
		Color:        widgets.Color{G: 255, A: 255},
		PressedColor: widgets.Color{G: 100, A: 255},
		TextColor:    widgets.Color{B: 255, A: 255},
	}
	GoPhoneTest.Widgets = append(GoPhoneTest.Widgets, &cir)

}

func main() {
	AddWidget()
	GoPhoneTest.StartApp()
}
