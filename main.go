package main

import (
	"GO_PROJECT/handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/even", handler.EvenAvg)
	http.HandleFunc("/odd", handler.OddAvg)
	http.HandleFunc("/evenodd", handler.EvenoddAvg)
	log.Println("server started")
	log.Fatal(http.ListenAndServe(":3100", nil))

}
