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

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body><h1>{{.}}</h1></body></html>`
	// temp, err := template.New("EXAMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	temp := template.Must(template.New("EXAMPLE").Parse(templateText))
	temp.ExecuteTemplate(writer, "EXAMPLE", "Hello World, This is HTML Template")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func SimpleHtmlFile(writer http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("./templates/simple.html")
	if err != nil {
		panic(err)
	}

	// temp:= template.Must(template.New("EXAMPLE").Parse(templateText))
	temp.ExecuteTemplate(writer, "simple.html", "Hello World, This is HTML Template")
}

func TestHtmlFile(t *testing.T) {
	url := "localhost:8080"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {

	temp := template.Must(template.ParseGlob("./templates/*.html"))
	temp.ExecuteTemplate(writer, "simple.html", "HTML Template Directory")
}

func TestTemplateDirectory(t *testing.T) {
	url := "localhost:8080"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*.html
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	temp := template.Must(template.ParseFS(templates, "templates/*.html"))
	temp.ExecuteTemplate(writer, "simple.html", "HTML Template Embed")
}

func TestTemplateEmbed(t *testing.T) {
	url := "localhost:8080"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
