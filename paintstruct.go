package main

type Widget interface {
	Paint()
	IsInside(x int, y int) bool
}
