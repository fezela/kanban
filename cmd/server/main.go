package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fezela/kanban"
)

func main(){

	http.HandleFunc("/cards", kanban.CreateCard)
	
	fmt.Println("Kanban service starting...")
	log.Println("Kanban listening on port 8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
