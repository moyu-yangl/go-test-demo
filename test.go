package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func m() {
	db, err := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/test??charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	//{
	//	id := 3
	//	username := "test"
	//	age := 18
	//	res, err := db.Exec("insert into user(id,name,age)values(?,?,?)", id, username, age)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(res.RowsAffected())
	//}
	{
		type user struct {
			id   int
			name string
			age  int
		}
		query := "select * from user "
		rows, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var users []user
		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, &u.name, &u.age)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		fmt.Println(users)
		//fmt.Printf("%#v", users)
	}
	//r := mux.NewRouter()
	//router := r.PathPrefix("/test").Subrouter()
	//CreateBook := func(w http.ResponseWriter, r *http.Request) {
	//	vars := mux.Vars(r)
	//	title := vars["title"]
	//	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, "page")
	//}
	//d := router.HandleFunc("/{title}", CreateBook).Methods("PUT").Host("localhost")
	//d.Schemes("http")
	//
	//http.ListenAndServe(":8080", r)
}
