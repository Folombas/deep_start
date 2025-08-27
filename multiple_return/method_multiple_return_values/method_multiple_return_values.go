package main

import "fmt"

// Структура Пользователь
type User struct {
	FirstName string
	LastName  string
	Age 	  int
}

// Метод возвращает полное имя и статус совершеннолетия
func (u User) GetInfo() (string, bool) {
	fullName := u.FirstName + " " + u.LastName
	isAdult := u.Age >= 18
	return fullName, isAdult
}

// Метод возвращает приветствие и ошибку (если имя пустое)
func (u User) Greet() (string, error) {
	if u.FirstName == "" {
		return "", fmt.Errorf("имя пользователя не указано")
	}
	greeting := fmt.Sprintf("Привет, %s %s! Тебе %d лет.", u.FirstName, u.LastName, u.Age)
	return greeting, nil
}

func main() {
	// Создаем пользователя
	user := User{
		FirstName: "Гоша",
		LastName: "Гофер",
		Age: 37,
	}


	//Используем метод с возвратом ошибки
	greeting, err := user.Greet()
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(greeting)
	}

	// Пример с путсым именем
	emptyUser := User{LastName: "Гофер", Age: 37}
	greeting, err = emptyUser.Greet()
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(greeting)
	}
}