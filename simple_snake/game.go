package main

import (
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	snakeBody         []Position
	apple             Position
	goldenApple       Position
	goldenAppleTimer  int
	moveDirection     int
	timer             int
	moveTime          int
	score             int
	bestScore         int
	level             int
	gameOver          bool
	paused            bool
	isMenu            bool
	selectedOption    int
	specialEffectTimer int
	specialEffectActive bool
}

func (g *Game) collidesWithApple() bool {
	return g.snakeBody[0].X == g.apple.X && g.snakeBody[0].Y == g.apple.Y
}

func (g *Game) collidesWithGoldenApple() bool {
	return g.snakeBody[0].X == g.goldenApple.X && g.snakeBody[0].Y == g.goldenApple.Y
}

func (g *Game) collidesWithSelf() bool {
	for _, v := range g.snakeBody[1:] {
		if g.snakeBody[0].X == v.X && g.snakeBody[0].Y == v.Y {
			return true
		}
	}
	return false
}

func (g *Game) collidesWithWall() bool {
	return g.snakeBody[0].X < 0 || g.snakeBody[0].Y < 0 ||
		g.snakeBody[0].X >= xGridCountInScreen || g.snakeBody[0].Y >= yGridCountInScreen
}

func (g *Game) needsToMoveSnake() bool {
	return g.timer%g.moveTime == 0
}

func (g *Game) reset() {
	g.apple.X = rand.IntN(xGridCountInScreen)
	g.apple.Y = rand.IntN(yGridCountInScreen)
	g.goldenApple.X = -1 // Скрыть золотое яблоко
	g.goldenApple.Y = -1
	g.goldenAppleTimer = 0
	g.moveTime = 10
	g.snakeBody = g.snakeBody[:1]
	g.snakeBody[0].X = xGridCountInScreen / 2
	g.snakeBody[0].Y = yGridCountInScreen / 2
	g.score = 0
	g.level = 1
	g.moveDirection = dirNone
	g.gameOver = false
	g.paused = false
	g.specialEffectActive = false
	g.specialEffectTimer = 0
}

func (g *Game) spawnGoldenApple() {
	if g.goldenAppleTimer <= 0 && rand.IntN(100) < 5 { // 5% шанс каждое обновление
		g.goldenApple.X = rand.IntN(xGridCountInScreen)
		g.goldenApple.Y = rand.IntN(yGridCountInScreen)
		g.goldenAppleTimer = 300 // Золотое яблоко существует 300 кадров
	}
}

func (g *Game) Update() error {
	if g.isMenu {
		return g.updateMenu()
	}
	
	if g.gameOver {
		return g.updateGameOver()
	}
	
	if g.paused {
		return g.updatePaused()
	}
	
	return g.updateGameplay()
}

func (g *Game) updateMenu() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selectedOption = (g.selectedOption - 1 + 3) % 3
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selectedOption = (g.selectedOption + 1) % 3
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch g.selectedOption {
		case 0: // Начать игру
			g.isMenu = false
			g.reset()
		case 1: // Уровень сложности
			g.level = (g.level % 3) + 1
		case 2: // Выйти
			return ebiten.Termination
		}
	}
	return nil
}

func (g *Game) updateGameOver() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.isMenu = true
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.isMenu = true
	}
	return nil
}

func (g *Game) updatePaused() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.paused = false
	}
	return nil
}

func (g *Game) updateGameplay() error {
	// Обработка паузы
	if inpututil.IsKeyJustPressed(ebiten.KeyP) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.paused = true
		return nil
	}
	
	// Управление змейкой
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if g.moveDirection != dirRight {
			g.moveDirection = dirLeft
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.moveDirection != dirLeft {
			g.moveDirection = dirRight
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.moveDirection != dirUp {
			g.moveDirection = dirDown
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if g.moveDirection != dirDown {
			g.moveDirection = dirUp
		}
	}
	
	// Спавн золотого яблока
	g.spawnGoldenApple()
	if g.goldenAppleTimer > 0 {
		g.goldenAppleTimer--
		if g.goldenAppleTimer <= 0 {
			g.goldenApple.X = -1
			g.goldenApple.Y = -1
		}
	}
	
	// Обновление специального эффекта
	if g.specialEffectActive {
		g.specialEffectTimer--
		if g.specialEffectTimer <= 0 {
			g.specialEffectActive = false
			g.moveTime = 10 - g.level*2 // Восстановление нормальной скорости
		}
	}
	
	if g.needsToMoveSnake() {
		if g.collidesWithWall() || g.collidesWithSelf() {
			g.gameOver = true
			return nil
		}
		
		// Проверка столкновения с обычным яблоком
		if g.collidesWithApple() {
			playSound(eatSound)
			g.apple.X = rand.IntN(xGridCountInScreen)
			g.apple.Y = rand.IntN(yGridCountInScreen)
			g.snakeBody = append(g.snakeBody, Position{
				X: g.snakeBody[len(g.snakeBody)-1].X,
				Y: g.snakeBody[len(g.snakeBody)-1].Y,
			})
			g.score += 1 * g.level
			
			// Увеличение уровня в зависимости от длины змейки
			if len(g.snakeBody) >= 10 && len(g.snakeBody) < 20 {
				g.level = 2
			} else if len(g.snakeBody) >= 20 {
				g.level = 3
			}
			
			// Обновление рекорда
			if g.bestScore < g.score {
				g.bestScore = g.score
			}
		}
		
		// Проверка столкновения с золотым яблоком
		if g.goldenAppleTimer > 0 && g.collidesWithGoldenApple() {
			playSound(goldenSound)
			g.goldenApple.X = -1
			g.goldenApple.Y = -1
			g.goldenAppleTimer = 0
			g.score += 5 * g.level
			g.specialEffectActive = true
			g.specialEffectTimer = 150 // Эффект длится 150 кадров
			g.moveTime = 15 - g.level*2 // Замедление времени
			
			// Добавление нескольких сегментов змейке
			for i := 0; i < 3; i++ {
				g.snakeBody = append(g.snakeBody, Position{
					X: g.snakeBody[len(g.snakeBody)-1].X,
					Y: g.snakeBody[len(g.snakeBody)-1].Y,
				})
			}
		}
		
		// Движение змейки
		for i := len(g.snakeBody) - 1; i > 0; i-- {
			g.snakeBody[i].X = g.snakeBody[i-1].X
			g.snakeBody[i].Y = g.snakeBody[i-1].Y
		}
		
		switch g.moveDirection {
		case dirLeft:
			g.snakeBody[0].X--
		case dirRight:
			g.snakeBody[0].X++
		case dirDown:
			g.snakeBody[0].Y++
		case dirUp:
			g.snakeBody[0].Y--
		}
	}
	
	g.timer++
	return nil
}


func newGame() *Game {
	g := &Game{
		apple:          Position{X: 3, Y: 3},
		goldenApple:    Position{X: -1, Y: -1},
		moveTime:       10,
		snakeBody:      make([]Position, 1),
		isMenu:         true,
		selectedOption: 0,
		level:          1, // Добавляем инициализацию уровня
	}
	g.snakeBody[0].X = xGridCountInScreen / 2
	g.snakeBody[0].Y = yGridCountInScreen / 2
	return g
}