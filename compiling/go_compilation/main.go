package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("🛠️  Демонстрация компиляции Go-кода")
	fmt.Println("====================================")
	
	// Определяем текущую платформу
	fmt.Printf("Текущая платформа: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println()
	
	// Компилируем различные примеры
	examples := []string{"hello", "calculator", "webapp"}
	
	for _, example := range examples {
		fmt.Printf("Компиляция примера: %s\n", example)
		fmt.Println(strings.Repeat("-", 40))
		
		// Компиляция для текущей платформы
		compileExample(example, "", "")
		
		// Кросс-компиляция для других платформ
		if runtime.GOOS != "linux" {
			compileExample(example, "linux", "amd64")
		}
		
		if runtime.GOOS != "windows" {
			compileExample(example, "windows", "amd64")
		}
		
		fmt.Println()
	}
	
	fmt.Println("✅ Все примеры скомпилированы!")
	fmt.Println("Бинарные файлы находятся в папке 'bin/'")
}

func compileExample(example, goos, goarch string) {
	examplePath := filepath.Join("examples", example)
	outputName := example
	
	if goos != "" && goarch != "" {
		outputName = fmt.Sprintf("%s_%s_%s", example, goos, goarch)
		if goos == "windows" {
			outputName += ".exe"
		}
	}
	
	outputPath := filepath.Join("bin", outputName)
	
	cmd := exec.Command("go", "build", "-o", outputPath, examplePath)
	
	// Устанавливаем переменные окружения для кросс-компиляции
	if goos != "" && goarch != "" {
		cmd.Env = append(os.Environ(), 
			fmt.Sprintf("GOOS=%s", goos), 
			fmt.Sprintf("GOARCH=%s", goarch))
	}
	
	fmt.Printf("Компилируем: %s ", examplePath)
	if goos != "" && goarch != "" {
		fmt.Printf("для %s/%s ", goos, goarch)
	}
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Ошибка: %s\n", err)
		if len(output) > 0 {
			fmt.Printf("Вывод компилятора: %s\n", string(output))
		}
	} else {
		fmt.Printf("✅ Успешно → %s\n", outputPath)
		
		// Показываем информацию о бинарнике
		if goos == "" || goos == runtime.GOOS {
			info, _ := os.Stat(outputPath)
			if info != nil {
				fmt.Printf("   Размер: %.2f KB\n", float64(info.Size())/1024)
			}
		}
	}
}