package main

import (
	"os"
	"path/filepath"
)

// Создаем примеры кода для компиляции
func createExamples() error {
	// Пример 1: Простое приложение Hello World
	helloCode := `package main

import "fmt"

func main() {
	fmt.Println("Привет, мир! 🚀")
	fmt.Println("Это скомпилированное Go-приложение")
}
`
	
	// Пример 2: Калькулятор
	calculatorCode := `package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Использование: calculator <число> <оператор> <число>")
		fmt.Println("Операторы: +, -, *, /")
		return
	}
	
	a, err1 := strconv.ParseFloat(os.Args[1], 64)
	op := os.Args[2]
	b, err2 := strconv.ParseFloat(os.Args[3], 64)
	
	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка: оба аргумента должны быть числами")
		return
	}
	
	var result float64
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return
		}
		result = a / b
	default:
		fmt.Println("Ошибка: неизвестный оператор")
		return
	}
	
	fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", a, op, b, result)
}
`
	
	// Пример 3: Простое веб-приложение
	webappCode := `package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Привет от Go-вебсервера!</h1>")
		fmt.Fprintf(w, "<p>Это скомпилированное веб-приложение</p>")
	})
	
	fmt.Println("Запуск сервера на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
`
	
	// Создаем файлы с примерами
	examples := map[string]string{
		"hello":      helloCode,
		"calculator": calculatorCode,
		"webapp":     webappCode,
	}
	
	for name, code := range examples {
		dirPath := filepath.Join("examples", name)
		os.MkdirAll(dirPath, 0755)
		
		filePath := filepath.Join(dirPath, "main.go")
		err := os.WriteFile(filePath, []byte(code), 0644)
		if err != nil {
			return err
		}
	}
	
	return nil
}

func init() {
	// Создаем примеры при запуске
	err := createExamples()
	if err != nil {
		panic(err)
	}
}