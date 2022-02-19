package golang_web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed templates/*.html
var templates_ embed.FS

var mytemplates = template.Must(template.ParseFS(templates_, "templates/*.html"))

func MyTemplateCaching(writer http.ResponseWriter, request *http.Request) {
	mytemplates.ExecuteTemplate(writer, "simple.html", "Caching Data")
}

func TestTemplateCaching(t *testing.T) {
	url := "localhost:8080"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	MyTemplateCaching(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
