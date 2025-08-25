package main

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
