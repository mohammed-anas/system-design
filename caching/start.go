package Server

import (
	"log"
	"net/http"
	appconfig "system-design/configs"
	servers "system-design/servers"
)

func RunServer() {
	main_server := http.NewServeMux()

	configuration := appconfig.LoadConfig()
	for _, appconfig := range configuration.Servers {
		server := servers.GetServer(appconfig.Name)
		go func(hostname string) {
			log.Fatal(http.ListenAndServe(hostname, server))
		}(appconfig.Hostname)
	}
	http.ListenAndServe("localhost:9000", main_server)
}