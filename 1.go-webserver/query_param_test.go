package gowebserver_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}
func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://fawzi.linggo.id:1233/hello?name=fawzi", nil)
	recorder := httptest.NewRecorder()

	sayHello(recorder, request)

	respons := recorder.Result()
	body, _ := io.ReadAll(respons.Body)

	fmt.Println(string(body))

}

func MultipleQuery(w http.ResponseWriter, r *http.Request) {

	firstname := r.URL.Query().Get("first_name")
	last_name := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s ", firstname, last_name)

}
func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://fawzi.linggo.id:1233/hello?first_name=fawzi&last_name=linggo", nil)
	recorder := httptest.NewRecorder()

	MultipleQuery(recorder, request)

	respons := recorder.Result()
	body, _ := io.ReadAll(respons.Body)

	fmt.Println(string(body))

}

func MultipleQueryValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprint(writer, strings.Join(names, "|"))
}

func TestMultipleParametersValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://fawzi.linggo.id:1233/testing?name=fawzi&name=linggo&name=andi&name=lukito", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryValues(recorder, request)

	respons := recorder.Result()
	body, _ := io.ReadAll(respons.Body)

	fmt.Println(string(body))
}
