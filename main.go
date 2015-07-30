package main

import (
	"flag"

	proxy "github.com/aitoroses/go-proxy/proxy"
)

func main() {

	// Get the config filename
	var filename string
	flag.StringVar(&filename, "config", "servers.json", "Name of the configuration filename")

	// Parse the flags
	flag.Parse()

	// Read the configuration file
	svs, _ := proxy.Read(filename)

	// Create a new server instance
	server := proxy.New(svs)

	// Start the server
	server.Start()
}
