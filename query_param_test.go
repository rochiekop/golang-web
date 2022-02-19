package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}

func MultipleQueryParam(writer http.ResponseWriter, request *http.Request) {
	first_name := request.URL.Query().Get("first_name")
	last_name := request.URL.Query().Get("last_name")

	if first_name == "" && last_name == "" {
		fmt.Fprint(writer, "Hello")

	} else {
		fmt.Fprintf(writer, "Hello %s %s", first_name, last_name)
	}
}

func TestMultipleQueryParam(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8080/hello?first_name=Rochi&last_name=Eko", nil)
	// request := httptest.NewRequest("GET", "localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParam(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Status)

}

func MultipleParamValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	name := query["name"]

	fmt.Fprintln(writer, strings.Join(name, " "))
}

func TestMultipleParamValues(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8080/hello?name=Rochi&name=Eko&name=Pambudi", nil)
	// request := httptest.NewRequest("GET", "localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	MultipleParamValues(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Status)

}

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8080/", nil)
	request.Header.Add("Content-Type", "application-type")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "X - Platinum Code")
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	poweredBy := recorder.Header().Get("x-powered-by")
	fmt.Println(poweredBy)
}
