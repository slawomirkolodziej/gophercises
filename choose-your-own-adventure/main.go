package main

import (
	"github.com/slawomirkolodziej/gophercises/choose-your-own-adventure/server"
	"github.com/slawomirkolodziej/gophercises/choose-your-own-adventure/storybuilder"
)

func main() {
	story := storybuilder.ParseFromJSON("./story.json")

	server.Start(story)
}
