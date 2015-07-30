package proxy

import (
	"encoding/json"
	"io/ioutil"
)

// Servers is a struct that handles the main json object
type Servers struct {
	Port    int      `json:port`
	Servers []Server `json:servers`
}

// Server handles each server data
type Server struct {
	Mount string `json:mount`
	Host  string `json:host`
	Port  int    `json:port`
}

// Read the servers.json file and parse its contents
func Read(filename string) (Servers, error) {
	var svs Servers
	dat, err := ioutil.ReadFile(filename)
	err = json.Unmarshal(dat, &svs)
	return svs, err
}
