package main

import (
	"fmt"

	"github.com/tmc/langchaingo/prompts"
)

func JinjaPromptTemplates() {

	templateWithJinja := prompts.PromptTemplate{
		Template:       "Translate '{{statement}}' from {{lang1}} to {{lang2}}",
		InputVariables: []string{"statement", "lang1", "lang2"},
		TemplateFormat: prompts.TemplateFormatJinja2,
	}

	template_jinja, _ := templateWithJinja.Format(map[string]any{
		"statement": "Wellcome to China",
		"lang1":     "English",
		"lang2":     "Chinese",
	})

	fmt.Println(template_jinja)

}
