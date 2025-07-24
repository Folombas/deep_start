// Анимация велосипеда в Go (простое решение)

package main

import (
	"fmt"
	"time"
)

func main() {
	frames := []string{
		"  __o   ",
		" _`\\<,_ ",
		"(*)/ (*)",
	}

	for i := 0; i < 10; i++ {
		fmt.Print("\033[2J") // Очистка экрана
		for _, frame := range frames {
			fmt.Println(frame)
		}
		time.Sleep(300 * time.Millisecond) // Пауза перед обновлением кадра
	}
}
