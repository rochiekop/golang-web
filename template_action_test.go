package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.html"))
	t.ExecuteTemplate(writer, "if.html", Page{
		Title: "HomePage",
	})
}

func TestTemplateActionIf(t *testing.T) {
	url := "localhost:8080"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.html"))
	t.ExecuteTemplate(writer, "comparator.html", map[string]interface{}{
		"Title":      "HomePage",
		"FinalValue": 90,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	url := "localhost:8080"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.html"))
	t.ExecuteTemplate(writer, "address.html", map[string]interface{}{
		"Title": "HomePage",
		"Name":  "Rochi Eko Pambudi",
		"Address": map[string]interface{}{
			"Street": "Jl.Sukasari Gegerkalong Sukajadi",
			"City":   "Bandung",
		},
	})
}

func TestActionWith(t *testing.T) {
	url := "localhost:8080"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
