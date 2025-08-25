package main

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
