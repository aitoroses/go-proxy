package main

import (
	"flag"

	proxy "github.com/aitoroses/GrayProxy/proxy"
)

/*
 	Configuration must be done in this way.
	{
	  "port": 8080,
	  "servers": [
	    {
	      "mount": "/app",
	      "host": "localhost",
	      "port": 3000
	    },
	    {
	      "mount": "/*",
	      "host": "soa-server",
	      "port": 7003
	    }
	  ]
	}
*/

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
