package main

import (
	"bakuvi/handler"
	"github.com/jmoiron/sqlx"

	"fmt"

	"bakuvi/store"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	db, err := sqlx.Open("mysql", "bestuser:bestuser@tcp(127.0.0.1:3306)/terra")
	if err != nil {
		panic(err)
	}

	s := store.Service{
		Conn: db,
	}

	h := handler.Service{S: &s}

	r := mux.NewRouter()
	r.HandleFunc("/get", h.GetIDs)
	//r.HandleFunc("/add", h.AddValue).Methods(http.MethodPost)

	if err := http.ListenAndServe(":3030", r); err != nil {
		fmt.Printf("server stoped with error: %v\n", err)
	}

	fmt.Printf("server stopped")
}
