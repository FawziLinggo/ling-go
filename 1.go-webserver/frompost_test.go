package gowebserver

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FromPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()

	if err != nil {
		panic(err)
	}

	fisrtName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hella %s %s", fisrtName, lastName)
}

func TestFromPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=fawzi&last_name=linggo")
	request := httptest.NewRequest(http.MethodPost, "http://fawzi.linggo.id:9090", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FromPost(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
