package professor

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
)

var token = "securitytoken"

var basicAuth = false
var basicAuthUser string
var basicAuthPassword string

// init disables default handlers registered by importing net/http/pprof.
func init() {
	http.DefaultServeMux = http.NewServeMux()
}

func SetBasicAuth(user, password string) {
	basicAuth = true
	basicAuthUser = user
	basicAuthPassword = password
}

func SetToken(t string) {
	token = t
}

// Handle adds standard pprof handlers to mux.
func Handle(mux *http.ServeMux) {
	mux.HandleFunc("/robots.txt", robots)
	mux.HandleFunc("/debug/pprof/", checkToken(pprof.Index))
	mux.HandleFunc("/debug/pprof/cmdline", checkToken(pprof.Cmdline))
	mux.HandleFunc("/debug/pprof/profile", checkToken(pprof.Profile))
	mux.HandleFunc("/debug/pprof/symbol", checkToken(pprof.Symbol))
	mux.HandleFunc("/debug/pprof/trace", checkToken(pprof.Trace))
}

// NewServeMux builds a ServeMux and populates it with standard pprof handlers.
func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	Handle(mux)
	return mux
}

// NewServer constructs a server at addr with the standard pprof handlers.
func NewServer(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: NewServeMux(),
	}
}

// ListenAndServe starts a server at addr with standard pprof handlers.
func ListenAndServe(addr string) error {
	return NewServer(addr).ListenAndServe()
}

// Launch a standard pprof server at addr.
func Launch(addr string) {
	go func() {
		log.Println(ListenAndServe(addr))
	}()
}

func checkToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !basicAuth && r.URL.Query().Get("token") != token {
			http.NotFound(w, r)
			return
		} else if basicAuth {
			user, pass, ok := r.BasicAuth()
			if user != basicAuthUser || pass != basicAuthPassword || !ok {
				w.Header().Set("WWW-Authenticate", `Basic realm="Please login"`)
				w.WriteHeader(401)
				w.Write([]byte("Unauthorised.\n"))
				return
			}
		}
		next.ServeHTTP(w, r)
	}
}

func robots(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User-agent: *")
	fmt.Fprintln(w, "Disallow: /")
}
