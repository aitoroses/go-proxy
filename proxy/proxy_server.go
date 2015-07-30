package proxy

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func proxyCall(domain string) func(w http.ResponseWriter, r *http.Request) {

	// Create a HTTP client
	client := &http.Client{}

	return func(w http.ResponseWriter, r *http.Request) {
		// Make an Http call to the proxied route
		url := domain + r.URL.Path
		fmt.Printf("%s %s\n", r.Method, url)

		// Create a Request
		req, err := http.NewRequest(r.Method, url, nil)
		if err != nil {
			io.WriteString(w, "Error creating the request to "+url)
		}

		// Make http call to proxy
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		// Write the response
		body, err := ioutil.ReadAll(resp.Body)
		w.Header().Set("X-proxy", "GrayProxy")
		io.WriteString(w, string(body))
	}
}

// Proxy structure
type Proxy struct {
	svs Servers
}

func (s *Proxy) getName(i int) string {
	sv := s.svs.Servers[i]
	return fmt.Sprintf("%s:%d", sv.Host, sv.Port)
}

func (s *Proxy) getPath(i int) string {
	return fmt.Sprintf("http://%s", s.getName(i))
}

// Start the proxy server
func (s *Proxy) Start() {

	// Create a new RegexpHandler as mux
	mux := &RegexpHandler{}
	for i := 0; i < len(s.svs.Servers); i++ {
		sv := s.svs.Servers[i]
		fmt.Printf("Proxy: %s -> %s\n", sv.Mount, s.getPath(i))
		// Add handler based on a regexp
		mux.HandleFunc(regexp.MustCompile("^"+sv.Mount), proxyCall(s.getPath(i)))
	}

	// Start listening
	http.ListenAndServe(":"+strconv.Itoa(s.svs.Port), mux)
}

// New returns a new ServeMux
func New(svs Servers) *Proxy {
	return &Proxy{svs}
}
