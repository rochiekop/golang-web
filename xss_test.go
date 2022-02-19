package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates_data embed.FS

var mytemplatesData = template.Must(template.ParseFS(templates_data, "templates/*.gohtml"))

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	mytemplatesData.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "XSS Test",
		"Body":  "<p>Hello World</p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	url := "localhost:8080"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	mytemplatesData.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "XSS Test",
		"Body":  template.HTML("<h1>Hello World</h1>"),
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	url := "localhost:8080"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServerDisabled(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	mytemplatesData.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "XSS Test",
		"Body":  template.HTML(request.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	// url := "http://localhost:8080/?body=<p>Hello This is XSS Test</p>"
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=Hello This is XSS Test", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
