package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/fezela/kanban"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)


func main(){
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	fmt.Println("database:", dbHost)
	fmt.Println("database port:", dbPort)
	http.HandleFunc("/cards", kanban.CreateCard)
	
	fmt.Println("Kanban service starting...")
	log.Println("Kanban listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
