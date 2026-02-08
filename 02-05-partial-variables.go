package main

import (
	"fmt"
	"time"

	"github.com/tmc/langchaingo/prompts"
)

func PartialVariablesPrompts() {

	templateWithPartials := prompts.PromptTemplate{
		Template:       "I want the {{company_name}} financial breakdown from {{date}}",
		InputVariables: []string{"company_name"},
		//PartialVariables: map[string]any{
		// "date": func() string {
		// 	return time.Now().Format("02-01-2006")
		// },
		//},
		TemplateFormat: prompts.TemplateFormatJinja2,
	}

	// Along the line
	templateWithPartials.PartialVariables = map[string]any{
		"date": time.Now().Format("02-01-2006"),
	}

	formatted_template, _ := templateWithPartials.Format(map[string]any{
		"company_name": "FZ Pharmaceuticals",
	})

	fmt.Println(formatted_template)
}
