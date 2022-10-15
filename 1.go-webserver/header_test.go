package gowebserver

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, request *http.Request) {
	ContentType := request.Header.Get("content-type")
	fmt.Fprint(w, ContentType)
}

func TestHeaderRequest(t *testing.T) {

	request := httptest.NewRequest("GET", "http://fawzi.linggo.id/", nil)
	request.Header.Add("Content-type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func ResponseHeader(w http.ResponseWriter, request *http.Request) {
	w.Header().Add("X-powered-bye", "SnowmanCode")
	fmt.Fprint(w, "OK")
}

func TestHeaderRespons(t *testing.T) {
	request := httptest.NewRequest("GET", "http://fawzi.linggo.id/", nil)
	request.Header.Add("Content-type", "application/json")
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Header.Get("x-powered-bye"))
}
