package gowebserver

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {

		// writer.WriteHeader(400) // Bad Request
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is empty")
	} else {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "hello %s ", name)
	}
}

func TestResponseCode(t *testing.T) {
	// BAD REQUEST (no name)
	// request := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	request := httptest.NewRequest(http.MethodGet, "http://localhost/?name=angga", nil)
	recorder := httptest.NewRecorder()

	ResponCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Status)

}
