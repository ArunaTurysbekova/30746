package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username string
	Password string
}

func main() {
	// Создание базы данных
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы пользователей
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		username TEXT PRIMARY KEY,
		password TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Обработчик страницы логина
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tmpl, err := template.ParseFiles("static/login.html")
			if err != nil {
				log.Fatal(err)
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				log.Fatal(err)
			}
		} else if r.Method == "POST" {
			username := r.FormValue("username")
			password := r.FormValue("password")

			// Проверка наличия пользователя в базе данных
			var user User
			err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&user.Username, &user.Password)
			if err != nil {
				fmt.Fprintln(w, `<script>alert("Ошибка авторизации!");</script>`)
				fmt.Fprintln(w, `<br><br><a href="/">Вернуться на главную страницу</a>`)
				return
			}

			// Проверка пароля
			if user.Password != password {
				fmt.Fprintln(w, `<script>alert("Ошибка авторизации!");</script>`)
				fmt.Fprintln(w, `<br><br><a href="/">Вернуться на главную страницу</a>`)
				return
			}

			// Отправляем JavaScript-код для отображения всплывающего уведомления
			fmt.Fprintln(w, `<script>alert("Авторизация прошла успешно!");</script>`)
			fmt.Fprintln(w, `<br><br><a href="/">Вернуться на главную страницу</a>`)
		}
	})

	// Обработчик страницы регистрации
	http.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tmpl, err := template.ParseFiles("static/registration.html")
			if err != nil {
				log.Fatal(err)
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				log.Fatal(err)
			}
		} else if r.Method == "POST" {
			username := r.FormValue("username")
			password := r.FormValue("password")

			// Добавление пользователя в базу данных
			stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
			if err != nil {
				log.Fatal(err)
			}
			_, err = stmt.Exec(username, password)
			if err != nil {
				fmt.Fprintln(w, `<script>alert("Ошибка регистрации!");</script>`)
				fmt.Fprintln(w, `<br><br><a href="/">Вернуться на главную страницу</a>`)
				return
			}

			// Отправляем JavaScript-код для отображения всплывающего уведомления
			fmt.Fprintln(w, `<script>alert("Регистрация прошла успешно!");</script>`)
			fmt.Fprintln(w, `<br><br><a href="/">Вернуться на главную страницу</a>`)
		}
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	log.Println("Сервер запущен на порту 1212...")
	err = http.ListenAndServe(":1212", nil)
	if err != nil {
		log.Fatal("Ошибка сервера: ", err)
	}
}
