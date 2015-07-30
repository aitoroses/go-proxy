package main

import proxy "github.com/aitoroses/proxy/proxy"

func main() {
	svs, _ := proxy.Read("servers.json")
	server := proxy.New(svs)
	server.Start()
}
