package main

const (
	dirNone = iota
	dirLeft
	dirRight
	dirDown
	dirUp
)

const (
	screenWidth        = 640
	screenHeight       = 480
	gridSize           = 20
	xGridCountInScreen = screenWidth / gridSize
	yGridCountInScreen = screenHeight / gridSize
)

type Position struct {
	X int
	Y int
}