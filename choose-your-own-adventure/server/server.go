package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	storyBuilder "github.com/slawomirkolodziej/gophercises/choose-your-own-adventure/storybuilder"
)

// Start is starting the http server
func Start(story map[string]storyBuilder.Arc) {
	const tpl = `
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset="UTF-8">
				<title>{{.Title}}</title>
			</head>
			<body>
			<h1>{{.Title}}</h1>
			{{range .Paragraphs}}
				<p>{{.}}</p>
			{{end}}
				{{range .Options}}
					<p>
						<a href="/{{.ArcLink}}">{{.Text}}</a>
					</p>
				{{end}}
			</body>
		</html>`

	template, err := template.New("story").Parse(tpl)

	if err != nil {
		panic(err)
	}

	handler := func(w http.ResponseWriter, req *http.Request) {
		url := strings.Replace(req.URL.Path, "/", "", 1)
		currentStory, found := story[url]
		if found {
			template.Execute(w, currentStory)
		} else {
			template.Execute(w, story["intro"])
		}

	}

	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
