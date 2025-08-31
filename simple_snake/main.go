package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Усовершенствованная Змейка на Go")
	
	// Инициализация аудио
	initAudio()
	
	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}