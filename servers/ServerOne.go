package servers

import (
	"net/http"
	"fmt"
)

type ServerOne struct {
	server *http.ServeMux
}

func (s1 *ServerOne) GetServerMux() *http.ServeMux {
	s1.doRequestMapping()
	return s1.server
}

func (s1 *ServerOne) doRequestMapping() {
	s1.server.HandleFunc("/", s1.homepage)
	s1.server.HandleFunc("/getCustomerDetails", s1.custDetails)
}


func (s1 *ServerOne) homepage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Running  Server homepage Mumbai")
}

func (s2 *ServerOne) custDetails(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Running  Server One Customer")
}


