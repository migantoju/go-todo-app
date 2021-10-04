package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// structs
type Todo struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	IsComplete  bool      `bool:"IsComplete"`
	Created     time.Time `string:"Created"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Welcome to the Homepage")
	fmt.Println("Endpoint Hit: Homepage")
	w.WriteHeader(http.StatusOK)
}

// createTodo, method for create a new TODO.
func createTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var todo Todo

	// get request Body and decode into todo
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// create the new Todo, with Create method.
	err = todo.Create()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

// getAll, method to get all the TODOS
func getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// todo pointer
	todo := new(Todo)

	// get all todos from database
	todos, err := todo.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// convert TODO Slice into json.
	result, err := json.Marshal(todos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// write status code header
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(result)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/todos", getAll)
	http.HandleFunc("/todos/new", createTodo)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Todo methods

func (t *Todo) Create() error {
	// get the database connection
	db := DbConnection()

	// insert query
	query := `INSERT INTO todos (title, description, is_complete, created)
				VALUES(?, ?, ?, ?)`

	// prepare
	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	// close the database connection
	defer stmt.Close()

	r, err := stmt.Exec(t.Title, t.Description, t.IsComplete, time.Now())

	if err != nil {
		return err
	}

	// check one row was executed
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada.")
	}

	return nil
}

func (t *Todo) GetAll() ([]Todo, error) {
	db := DbConnection()

	todos := []Todo{}

	query := `SELECT * FROM todos;`

	// execute the query
	rows, err := db.Query(query)

	if err != nil {
		return []Todo{}, err
	}

	// close the database connection
	defer rows.Close()

	for rows.Next() {
		rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.IsComplete,
			&t.Created,
		)

		// add every new TODO into the slice
		todos = append(todos, *t)
	}

	return todos, nil
}

// Update, Struct method for TODO to update data.
func (t *Todo) Update() error {
	db := DbConnection()

	query := `UPDATE todos SET title=?, description=?, is_complete=? WHERE id=?`

	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	row, err := stmt.Exec(t.Title, t.Description, t.IsComplete)
	// err := stmt.QueryRow(t.Title, t.Description, t.IsComplete).Scan(&todo.Title, &todo.Description, &todo.IsComplete)
	//row, err := stmt.QueryRow(t.Title, t.Description, t.IsComplete).Scan(&todo.Title, &todo.Description, &todo.IsComplete)

	if err != nil {
		return err
	}

	if i, err := row.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada.")
	}
	return nil
}

func main() {
	fmt.Println("Running the TODO APP..")

	migrate := flag.Bool("migrate", false, "Create database tables")

	flag.Parse()

	if *migrate {
		if err := MakeMigrations(); err != nil {
			log.Fatal(err)
		}
	}

	handleRequests()
}
