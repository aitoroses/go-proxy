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
		// URL to which proxy the call
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

// GetName gives the host:port pair as string
func (s *Proxy) GetName(i int) string {
	sv := s.svs.Servers[i]
	return fmt.Sprintf("%s:%d", sv.Host, sv.Port)
}

// GetPath gives the full path to call
func (s *Proxy) GetPath(i int) string {
	return fmt.Sprintf("http://%s", s.GetName(i))
}

// Start the proxy server
func (s *Proxy) Start() {

	// Create a new RegexpHandler as mux
	mux := &RegexpHandler{}

	// Iterate over configured servers to setup the middleware
	for i := 0; i < len(s.svs.Servers); i++ {
		sv := s.svs.Servers[i]

		// Print information about routes
		fmt.Printf("Proxy: %s  ->  %s\n", sv.Mount, s.GetPath(i))

		// Add handler based on a regexp with the info extracted from the JSON
		mux.HandleFunc(regexp.MustCompile("^"+sv.Mount), proxyCall(s.GetPath(i)))
	}

	// Start listening
	http.ListenAndServe(":"+strconv.Itoa(s.svs.Port), mux)
}

// New returns a new Proxy server instance configured
func New(svs Servers) *Proxy {
	return &Proxy{svs}
}
