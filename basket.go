package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Basket struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func main() {
	http.HandleFunc("/api/basket", handleBasket)
	http.HandleFunc("/api/basket/all", getBasket)

	db, err := createDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getBasket(w http.ResponseWriter, r *http.Request) {
	db, err := createDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	lowerFilter := r.URL.Query().Get("lower")
	upperFilter := r.URL.Query().Get("upper")
	nameFilter := r.URL.Query().Get("name")

	selectQuery := `
		SELECT id, name, price, quantity FROM items
		WHERE (price > ? OR ? = '') AND (price < ? OR ? = '') AND (name LIKE '%' || ? || '%' OR ? = '');
	`
	rows, err := db.Query(selectQuery, lowerFilter, lowerFilter, upperFilter, upperFilter, nameFilter, nameFilter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.Quantity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	basket := Basket{
		Items: items,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basket)
}

func createDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			price INTEGER,
			quantity INTEGER
		);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func handleBasket(w http.ResponseWriter, r *http.Request) {
	basket := Basket{}
	err := json.NewDecoder(r.Body).Decode(&basket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := createDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	for _, item := range basket.Items {
		insertQuery := `
			INSERT INTO items (name, price, quantity)
			VALUES (?, ?, ?);
		`
		_, err := db.Exec(insertQuery, item.Name, item.Price, item.Quantity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basket)
}
