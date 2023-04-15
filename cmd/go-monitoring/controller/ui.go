package controller

import (
	"bytes"
	"net/http"
	"text/template"
)

func renderUITemplate(path string, data any) bytes.Buffer {
	t, err := template.ParseFiles(path, "templates/parts/header.html", "templates/parts/footer.html")

	if err != nil {
		panic(err)
	}

	var w bytes.Buffer
	err = t.Execute(&w, data)

	if err != nil {
		panic(err)
	}
	return w
}

func wrapUITemplate(path string) bytes.Buffer {
	return renderUITemplate(path, struct{}{})
}

func RegisterUIController() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/check/", func(w http.ResponseWriter, r *http.Request) {
		bb := wrapUITemplate("templates/check_details.html")
		w.Write(bb.Bytes())
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bb := wrapUITemplate("templates/index.html")
		w.Write(bb.Bytes())
	})
	return mux
}
