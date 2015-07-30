package main

import (
	"net/http"

	proxy "github.com/aitoroses/proxy/proxy"
)

func main() {
	svs, _ := proxy.Read("servers.json")
	server := proxy.New(svs)

	go http.ListenAndServe(":3000", nil)
	go http.ListenAndServe(":7003", nil)

	server.Start()

}
