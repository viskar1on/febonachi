package ui

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"main.go/logic"
)

// Функция для отрисовки страницы в браузере
func Render(w http.ResponseWriter, r *http.Request, filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "%s", file)
}

// Обработчик GET запроса на index.html
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "./html/index.html")
}

// Обработчик POST запроса result
func ResultHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "./html/result.html")
	number, err := strconv.Atoi(r.FormValue("number"))

	if err != nil || number < 0 {
		fmt.Fprintf(w, "<h1>Неправильно введено число!</h1>")
	} else {
		if logic.IsFibonacci(number) {
			previous, next := logic.GetFibonacciNeighbours(number)
			fmt.Fprintf(w, "<p>Это число Фибоначчи</p>")
			fmt.Fprintf(w, "<p>Предыдущее число Фибоначчи %d</p>", previous)
			fmt.Fprintf(w, "<p>Следующее число Фибоначчи %d</p>", next)
		} else {
			closest := logic.GetNearestFibonacci(number)
			fmt.Fprintf(w, "<p>Это НЕ число Фибоначчи</p>")
			fmt.Fprintf(w, "<p>Ближнее число Фибоначчи %d</p>", closest)
		}
	}
	fmt.Fprintf(w, "<p><a href=\"/\">Ввести ещё</a></p>")
}

// Функция для запуска сервера
func StartServer() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/result", ResultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
