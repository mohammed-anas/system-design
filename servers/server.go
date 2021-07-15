package servers

import (
	"net/http"
)

type Server interface{
	GetServerMux()
}

func GetServer(name string) *http.ServeMux {
	switch name {
	case "Mumbai":
		mumbai_server := ServerOne{server: http.NewServeMux()}
		mumbai_server.doRequestMapping()
		return mumbai_server.server
	case "Bangalore":
		bangalore_server := ServerTwo{server: http.NewServeMux()}
		bangalore_server.doRequestMapping()
		return bangalore_server.server
	default:
		mumbai_server := ServerOne{server: http.NewServeMux()}
		mumbai_server.doRequestMapping()
		return mumbai_server.server
	}
}