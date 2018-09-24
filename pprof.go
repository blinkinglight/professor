package professor

import (
	"log"
	"net/http"
	"net/http/pprof"
)

var token = "securitytoken"

// init disables default handlers registered by importing net/http/pprof.
func init() {
	http.DefaultServeMux = http.NewServeMux()
}

func SetToken(t string) {
	token = t
}

// Handle adds standard pprof handlers to mux.
func Handle(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", handleIndex)
	mux.HandleFunc("/debug/pprof/cmdline", handleCmdline)
	mux.HandleFunc("/debug/pprof/profile", handleProfile)
	mux.HandleFunc("/debug/pprof/symbol", handleSymbol)
	mux.HandleFunc("/debug/pprof/trace", handleTrace)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if !checkToken(w, r) {
		return
	}
	pprof.Index(w, r)
}

func handleCmdline(w http.ResponseWriter, r *http.Request) {
	if !checkToken(w, r) {
		return
	}
	pprof.Cmdline(w, r)
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	if !checkToken(w, r) {
		return
	}
	pprof.Profile(w, r)
}

func handleSymbol(w http.ResponseWriter, r *http.Request) {
	if !checkToken(w, r) {
		return
	}
	pprof.Symbol(w, r)
}

func handleTrace(w http.ResponseWriter, r *http.Request) {
	if !checkToken(w, r) {
		return
	}
	pprof.Symbol(w, r)
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

func checkToken(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Query().Get("token") != token {
		http.NotFound(w, r)
		return false
	}
	return true
}
