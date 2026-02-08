package main

import (
	"fmt"

	"github.com/tmc/langchaingo/prompts"
)

func MultiLineTemplates() {

	templateString := `
		Display the following list of {{.profession}}
		{{.names}}

		{{if .display_examples}}
		Examples:
		{{range .display_examples}}
		- {{.}}
		{{end}}
		{{end}}
	`

	template_multiline := prompts.PromptTemplate{

		Template:       templateString,
		InputVariables: []string{"profession", "names", "display_examples"},
		TemplateFormat: prompts.TemplateFormatGoTemplate,
	}

	formatted_template, _ := template_multiline.Format(map[string]any{
		"profession":       "Footballers",
		"names":            []string{"Ronaldo", "Messi", "De Bryne Haaland"},
		"display_examples": []string{"Kaka AC Milan", "Rooney - Manchester United"},
	})

	fmt.Println(formatted_template)

}
