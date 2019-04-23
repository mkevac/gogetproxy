package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

var defaultAnswer = `
	<html>
		<head>
			<meta name="go-import" content="go.badoo.dev{{.Name}} git git+ssh://git@cppci1.msk/{{.Path}}">
		</head>
		<body>
		</body>
	</html>
`

type tmplStruct struct {
	Name string
	Path string
}

func main() {

	tmpl, err := template.New("main").Parse(defaultAnswer)
	if err != nil {
		log.Fatalf("Error while parsing template: %s", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if len(r.URL.Path) <= 1 || r.URL.Path[0] != '/' {
			if err := tmpl.Execute(w, tmplStruct{"", "go/badoo2"}); err != nil {
				log.Printf("Error while executing template: %s", err)
			}
			return
		}

		splitted := strings.Split(r.URL.Path, "/")
		name := splitted[1]

		log.Printf("name is '%s', url is '%v'", name, r.URL.Path)

		if name == "core" {
			if err := tmpl.Execute(w, tmplStruct{"/core", "go/badoo2"}); err != nil {
				log.Printf("Error while executing template: %s", err)
			}
		} else {
			if err := tmpl.Execute(w, tmplStruct{"/" + name, "go/" + name}); err != nil {
				log.Printf("Error while executing template: %s", err)
			}
		}
	})

	if err := http.ListenAndServeTLS(":443",
		"go.badoo.dev.pem",
		"go.badoo.dev-key.pem",
		nil); err != nil {
		log.Fatal(err)
	}
}
