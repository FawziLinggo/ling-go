package gowebserver

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

//go:embed resources/ok.html
var resourcesOK string

//go:embed resources/404.html
var resources404 string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name != "" {
		fmt.Fprint(writer, resourcesOK, name)
	} else {
		fmt.Fprint(writer, resources404)
	}
}

func TestServeFileEmbed(t *testing.T) {

	server := http.Server{
		Addr:    "fawzi.linggo.id:9123",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
