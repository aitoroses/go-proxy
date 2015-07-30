package proxy

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func proxyCall(s *Proxy) func(w http.ResponseWriter, r *http.Request) {

	// Create a HTTP client
	client := &http.Client{}

	return func(w http.ResponseWriter, r *http.Request) {

		var resp *http.Response

		// Iterate through servers
		for i := 0; i < len(s.svs.Servers); i++ {

			s := s.svs.Servers[i]

			// URL to which proxy the call
			url := fmt.Sprintf("http://%s:%d%s", s.Host, s.Port, r.URL.Path)
			fmt.Printf("%s %s\n", r.Method, url)

			// Create a Request
			req, err := http.NewRequest(r.Method, url, nil)
			if err != nil {
				io.WriteString(w, "Error creating the request to "+url)
			}

			// Make http call to proxy
			// If the call succeeds break the loop
			resp, err = client.Do(req)
			if err == nil && (resp.StatusCode < 400) {
				// Call has success
				break
			}
		}

		// Write the response extracting the info from the proxied call
		if resp != nil {
			body, _ := ioutil.ReadAll(resp.Body)
			w.Header().Set("X-Proxy", "go-proxy")
			io.WriteString(w, string(body))
		} else {
			io.WriteString(w, "Not found.")
		}
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
	mux := http.NewServeMux()

	// Add handler based on the info extracted from the JSON
	mux.HandleFunc("/", proxyCall(s))

	// Start listening
	http.ListenAndServe(":"+strconv.Itoa(s.svs.Port), mux)
}

// New returns a new Proxy server instance configured
func New(svs Servers) *Proxy {
	return &Proxy{svs}
}
