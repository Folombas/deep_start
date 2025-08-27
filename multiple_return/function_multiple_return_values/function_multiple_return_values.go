package main 

import "fmt"

// Функция возвращает результат и ошибку
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return a/b, nil
}

// Функция возвращает три значения
func calculate(a, b int) (int, int, int) {
	sum := a + b
	diff := a - b
	mult := a * b
	return sum, diff, mult
}

func main() {
	// Пример с обработкой ошибки
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("10 / 2 = %.1f\n", result)
	}

	// Пример с тремя возвращаемыми значениями
	sum, diff, mult := calculate(10, 5)
	fmt.Printf("Сумма: %d, Разность: %d, Произведение: %d\n", sum, diff, mult)

	// Игнорирование одного из значений
	sum, _, mult = calculate(8, 2)
	fmt.Printf("Сумма: %d, Произведение: %d (разность проигнорирована)\n", sum, mult)
}