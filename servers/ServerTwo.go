package servers

import (
	"net/http"
	"fmt"
)

type ServerTwo struct {
	server *http.ServeMux
}

func (s2 *ServerTwo) GetServerMux() *http.ServeMux {
	s2.doRequestMapping()
	return s2.server
}

func (s2 *ServerTwo) doRequestMapping() {
	s2.server.HandleFunc("/", s2.homepage)
	s2.server.HandleFunc("/getCustomerDetails", s2.custDetails)
}


func (s2 *ServerTwo) homepage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Running  Server homepage Bangalore")
}

func (s2 *ServerTwo) custDetails(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Running  Server customer")
}

