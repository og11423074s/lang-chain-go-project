package main

import (
	"fmt"

	"github.com/tmc/langchaingo/prompts"
)

func StandardTemplateDefinition() {

	templateWithProps := prompts.PromptTemplate{
		Template:       "Research {{ .topic }} on {{.website}}",
		InputVariables: []string{"topic", "website"},
		TemplateFormat: prompts.TemplateFormatGoTemplate,
	}

	template_with_props, _ := templateWithProps.Format(map[string]any{
		"topic":   "Cookie Recipes",
		"website": "The Food Network",
	})

	fmt.Println(template_with_props)

}
