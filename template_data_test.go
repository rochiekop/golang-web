package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.html"))
	t.ExecuteTemplate(writer, "name.html", map[string]interface{}{
		"Title": "HomePage",
		"Name":  "Rochi Eko Pambudi",
		"Address": map[string]interface{}{
			"Street": "Jl. Sukasari Geger Kalong",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	url := "localhost:8080"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// USING STRUCT

type Address struct {
	Street string
}

type Page struct {
	Title   string
	Name    string
	Address Address

}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.html"))
	t.ExecuteTemplate(writer, "name.html", Page{
		Title: "HomePage",
		Name:  "Rochi Eko Pambudi",
		Address: Address{
			Street: "Jl.Sukasari Geger Kalong",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	url := "localhost:8080"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
