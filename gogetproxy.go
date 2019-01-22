package main

import (
	"html/template"
	"log"
	"net/http"
)

var defaultAnswer = `
	<html>
		<head>
			<meta name="go-import" content="cbadoo.io/{{.}} git git+ssh://git@cppci1.msk/go/{{.}}">
		</head>
		<body>
		</body>
	</html>
`

func main() {

	tmpl, err := template.New("main").Parse(defaultAnswer)
	if err != nil {
		log.Fatalf("Error while parsing template: %s", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if len(r.URL.Path) <= 1 || r.URL.Path[0] != '/' {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		name := r.URL.Path[1:]

		if err := tmpl.Execute(w, name); err != nil {
			log.Printf("Error while executing template: %s", err)
		}
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
