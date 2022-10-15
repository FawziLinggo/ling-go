package gowebserver

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestFileServer(t *testing.T) {
	// Menggunakan Golang embed
	// directory := http.Dir("./resources")
	directory, error := fs.Sub(resources, "resources")

	if error != nil {
		fmt.Println("Directory not found")
		panic(error)
	}

	filesServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", filesServer))

	server := http.Server{
		Addr:    "fawzi.linggo.id:9123",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
