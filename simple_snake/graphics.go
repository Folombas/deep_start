package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Draw(screen *ebiten.Image) {
	if g.isMenu {
		g.drawMenu(screen)
		return
	}
	
	// Отрисовка фона
	screen.Fill(color.RGBA{0x10, 0x10, 0x20, 0xff})
	
	// Отрисовка сетки
	for x := 0; x < xGridCountInScreen; x++ {
		for y := 0; y < yGridCountInScreen; y++ {
			if (x+y)%2 == 0 {
				vector.DrawFilledRect(screen, float32(x*gridSize), float32(y*gridSize), 
					gridSize, gridSize, color.RGBA{0x20, 0x20, 0x30, 0xff}, false)
			}
		}
	}
	
	// Отрисовка змейки
	for i, v := range g.snakeBody {
		snakeColor := color.RGBA{0x00, 0xff, 0x00, 0xff}
		if i == 0 {
			snakeColor = color.RGBA{0x00, 0xcc, 0x00, 0xff} // Голова змейки темнее
		} else if g.specialEffectActive && g.timer%10 < 5 {
			snakeColor = color.RGBA{0xff, 0xcc, 0x00, 0xff} // Мерцание при специальном эффекте
		}
		vector.DrawFilledRect(screen, float32(v.X*gridSize), float32(v.Y*gridSize), 
			gridSize, gridSize, snakeColor, false)
	}
	
	// Отрисовка обычного яблока
	vector.DrawFilledRect(screen, float32(g.apple.X*gridSize), float32(g.apple.Y*gridSize), 
		gridSize, gridSize, color.RGBA{0xFF, 0x00, 0x00, 0xff}, false)
	
	// Отрисовка золотого яблока (если активно)
	if g.goldenAppleTimer > 0 {
		goldenColor := color.RGBA{0xff, 0xcc, 0x00, 0xff}
		if g.timer%10 < 5 { // Мерцание
			goldenColor = color.RGBA{0xff, 0xee, 0x00, 0xff}
		}
		vector.DrawFilledRect(screen, float32(g.goldenApple.X*gridSize), float32(g.goldenApple.Y*gridSize), 
			gridSize, gridSize, goldenColor, false)
	}
	
	// Отрисовка UI
	if g.moveDirection == dirNone && !g.gameOver {
		ebitenutil.DebugPrintAt(screen, "ИСПОЛЬЗУЙТЕ СТРЕЛКИ ДЛЯ НАЧАЛА ИГРЫ", screenWidth/2-150, screenHeight/2-10)
	}
	
	if g.gameOver {
		ebitenutil.DebugPrintAt(screen, "ИГРА ОКОНЧЕНА! НАЖМИТЕ ESC ДЛЯ ВЫХОДА", screenWidth/2-180, screenHeight/2-10)
	} else if g.paused {
		ebitenutil.DebugPrintAt(screen, "ПАУЗА - НАЖМИТЕ P ДЛЯ ПРОДОЛЖЕНИЯ", screenWidth/2-150, screenHeight/2-10)
	}
	
	// Отображение счета и другой информации
	infoText := fmt.Sprintf("Счет: %d | Рекорд: %d | Уровень: %d | Длина: %d", 
		g.score, g.bestScore, g.level, len(g.snakeBody))
	if g.specialEffectActive {
		infoText += " | ЗОЛОТАЯ СИЛА!"
	}
	ebitenutil.DebugPrintAt(screen, infoText, 10, 10)
	
	// Отображение FPS
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.0f", ebiten.ActualFPS()), screenWidth-100, 10)
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x20, 0x20, 0x40, 0xff})
	
	// Заголовок
	ebitenutil.DebugPrintAt(screen, "ЗМЕЙКА НА GO", screenWidth/2-70, 50)
	
	// Опции меню
	options := []string{
		fmt.Sprintf("НОВАЯ ИГРА (Уровень %d)", g.level),
		"СЛОЖНОСТЬ: " + []string{"ЛЕГКИЙ", "СРЕДНИЙ", "СЛОЖНЫЙ"}[g.level-1],
		"ВЫХОД",
	}
	
	// Защита от выхода за границы массива
	if g.selectedOption < 0 {
		g.selectedOption = 0
	} else if g.selectedOption >= len(options) {
		g.selectedOption = len(options) - 1
	}
	
	for i, option := range options {
		yPos := screenHeight/2 - 30 + i*40
		if i == g.selectedOption {
			ebitenutil.DebugPrintAt(screen, "> "+option+" <", screenWidth/2-100, yPos)
		} else {
			ebitenutil.DebugPrintAt(screen, "  "+option+"  ", screenWidth/2-100, yPos)
		}
	}
	
	// Управление
	ebitenutil.DebugPrintAt(screen, "УПРАВЛЕНИЕ: СТРЕЛКИ/WASD - ДВИЖЕНИЕ, P - ПАУЗА, ESC - МЕНЮ", 50, screenHeight-30)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}