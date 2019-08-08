package main

import (
	"os"
	"text/template"
)

// Todo a single todo wrap into a struct
type Todo struct {
	Name        string
	Description string
}

func main() {
	td := Todo{"Test templates", "Let's test a template to see the magic."}

	t := template.Must(template.New("todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"\n"))

	err := t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}

	tdNew := Todo{"Go", "Contribute to any Go project"}
	err = t.Execute(os.Stdout, tdNew)
}
