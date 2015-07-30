package proxy

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

func hello(message string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, message)
	}
}

// Proxy structure
type Proxy struct {
	svs Servers
}

// Start the proxy server
func (s *Proxy) Start() {
	// Create a new RegexpHandler as mux
	mux := &RegexpHandler{}
	for i := 0; i < len(s.svs.Servers); i++ {
		s := s.svs.Servers[i]
		fmt.Println(s.Mount)
		mux.HandleFunc(regexp.MustCompile("^"+s.Mount), hello(fmt.Sprintf("%s:%d", s.Host, s.Port)))
	}

	// Start listening
	http.ListenAndServe(":"+strconv.Itoa(s.svs.Port), mux)
}

// New returns a new ServeMux
func New(svs Servers) *Proxy {
	return &Proxy{svs}
}
