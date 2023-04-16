package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// Структура для представления человека
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Срез для хранения людей
var people []Person

func main() {
	// Обработчик для корневого маршрута
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Шаблон страницы со списком людей
		tmpl := template.Must(template.ParseFiles("static/template.html"))

		// Выводим список людей
		err := tmpl.Execute(w, people)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Обработчик для маршрута получения списка людей
	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		// Отправляем список людей в формате JSON
		json.NewEncoder(w).Encode(people)
	})

	// Обработчик для маршрута добавления нового человека
	http.HandleFunc("/people/add", func(w http.ResponseWriter, r *http.Request) {
		// Получаем данные из формы
		name := r.FormValue("name")
		age := r.FormValue("age")

		// Преобразуем возраст в целочисленный тип
		ageInt := 0
		if age != "" {
			ageInt = parseAge(age)
		}

		// Создаем нового человека и добавляем его в список
		person := Person{
			ID:   len(people) + 1,
			Name: name,
			Age:  ageInt,
		}
		people = append(people, person)

		// Перенаправляем на главную страницу
		http.Redirect(w, r, "/", http.StatusFound)
	})

	// Запускаем веб-сервер на порту 8081
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Функция для преобразования возраста из строки в целочисленный тип
func parseAge(age string) int {
	ageInt := 0
	for _, ch := range age {
		if ch >= '0' && ch <= '9' {
			ageInt = ageInt*10 + int(ch-'0')
		} else {
			break
		}
	}
	return ageInt
}
