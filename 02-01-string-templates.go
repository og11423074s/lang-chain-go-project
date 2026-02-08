package main

import (
	"fmt"

	"github.com/tmc/langchaingo/prompts"
)

func stringPromptTemplates() {

	simpleTemplate := prompts.NewPromptTemplate(
		"Write a {{.content_type}} about {{.subject}}!",
		[]string{"content_type", "subject"},
	)

	templateInput := map[string]any{

		"content_type": "poem",
		"subject":      "cats",
	}

	simple_prompt, _ := simpleTemplate.Format(templateInput)

	fmt.Println(simple_prompt)

}
