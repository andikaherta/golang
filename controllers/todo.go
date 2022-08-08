package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ichtrojan/go-todo/config"
	"github.com/ichtrojan/go-todo/models"
)

var (
	id        int
	item      string
	assignee  string
	deadline  string
	completed int
	view      = template.Must(template.ParseFiles("./views/index.html"))
	database  = config.Database()
)

func Show(w http.ResponseWriter, r *http.Request) {
	statement, err := database.Query(`SELECT * FROM todos`)

	if err != nil {
		fmt.Println(err)
	}

	var todos []models.Todo

	for statement.Next() {
		err = statement.Scan(&id, &item, &assignee, &deadline, &completed)

		if err != nil {
			fmt.Println(err)
		}

		todo := models.Todo{
			Id:        id,
			Item:      item,
			Assignee:  assignee,
			Deadline:  deadline,
			Completed: completed,
		}

		todos = append(todos, todo)
	}

	data := models.View{
		Todos: todos,
	}

	_ = view.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {

	item := r.FormValue("item")
	assignee := r.FormValue("assignee")
	deadline := r.FormValue("deadline")

	_, err := database.Exec(`INSERT INTO todos (item, assignee, deadline) VALUE (?,?,?)`, item, assignee, deadline)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {

	item := r.FormValue("item")
	assignee := r.FormValue("assignee")
	deadline := r.FormValue("deadline")
	id := r.FormValue("id")

	_, err := database.Exec(`UPDATE todos SET item = ?, assignee = ?, deadline = ?  WHERE id = ?`, item, assignee, deadline, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`DELETE FROM todos WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func Complete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`UPDATE todos SET completed = 1 WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}
