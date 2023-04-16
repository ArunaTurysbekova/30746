package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people []Person

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("static/template.html"))

		err := tmpl.Execute(w, people)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(people)
	})

	http.HandleFunc("/people/add", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		age := r.FormValue("age")

		ageInt := 0
		if age != "" {
			ageInt = parseAge(age)
		}

		person := Person{
			ID:   len(people) + 1,
			Name: name,
			Age:  ageInt,
		}
		people = append(people, person)

		http.Redirect(w, r, "/", http.StatusFound)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

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
