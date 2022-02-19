package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()

	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("firstName")
	lastName := request.PostForm.Get("lastName")

	if firstName != "" && lastName != "" {
		fmt.Fprintf(writer, "%s %s", firstName, lastName)
	} else {
		fmt.Fprint(writer, "Hello")
	}

}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstName=Rochi&lastName=Eko")
	// requestBody := strings.NewReader("firstName=&lastName=")
	request := httptest.NewRequest("POST", "http://localhost:8080/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Status)
}
