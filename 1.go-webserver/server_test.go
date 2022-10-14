package gowebserver

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "fawzi.linggo.id:1232",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
