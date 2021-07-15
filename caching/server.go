package Server

import (
	"log"
	"fmt"
	"net/http"
)

func RunServer(port string) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Running  Server at port %s",port)
	})

	log.Fatal(http.ListenAndServe("localhost:" + port, nil))
}