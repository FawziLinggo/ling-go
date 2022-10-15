package gowebserver

import (
	"fmt"
	"net/http"
	"testing"
)

func SetCookie(writes http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "Coockie-fawzi"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writes, cookie)
	fmt.Fprintln(writes, "succeses set cookie")
}

func GetCookie(writes http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("Coockie-fawzi")

	if err != nil {
		fmt.Fprintln(writes, "no cookie")
	} else {
		name := cookie.Value
		fmt.Fprintln(writes, "Cookie ", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "fawzi.linggo.id:3211",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
