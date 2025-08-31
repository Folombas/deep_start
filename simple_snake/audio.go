package main

import (
	"time"

	"github.com/hajimehoshi/oto/v2"
)

type SoundType int

const (
	eatSound SoundType = iota
	goldenSound
)

var (
	otoContext *oto.Context
)

func initAudio() {
	var ready chan struct{}
	otoContext, ready, _ = oto.NewContext(44100, 2, 2)
	<-ready
}

func playSound(soundType SoundType) {
	// В учебных целях просто игнорируем звуки
	// В реальной игре здесь была бы реализация звуковых эффектов
}

func beep(freq float64, duration time.Duration) {
	// Упрощенная версия без генерации звука
	// В реальной игре здесь была бы генерация звуковых волн
}

func sin(x float64) float64 {
	// Простая реализация синуса
	return 0 // Заглушка для компиляции
}