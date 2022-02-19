package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.html",
		"./templates/footer.html",
		"./templates/body.html"))

	t.ExecuteTemplate(writer, "body", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Rochi Eko Pambudi",
	})
}

func TestTemplateLayout(t *testing.T) {
	url := "localhost:8080"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
