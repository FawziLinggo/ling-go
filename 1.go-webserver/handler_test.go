package gowebserver

import (
	"fmt"
	"net/http"
	"testing"
)

// func TestHandler(t *testing.T) {

// 	var handler http.HandlerFunc = func(writer http.ResponseWriter, Request *http.Request) {

// 		// Logic web
// 		// fmt.Fprint(writer, "hello world")
// 		_, err := fmt.Fprint(writer, "hello world")
// 		if err != nil {
// 			return
// 		}

// 	}

// 	server := http.Server{
// 		Addr:    "fawzi.linggo.id:1232",
// 		Handler: handler,
// 	}

// 	err := server.ListenAndServe()

// 	if err != nil {
// 		panic(err)
// 	}
// }

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "hello world")
		if err != nil {
			return
		}
	})
	mux.HandleFunc("/shopping", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "hello shopping")
		if err != nil {
			return
		}
	})
	server := http.Server{
		Addr:    "fawzi.linggo.id:1233",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
