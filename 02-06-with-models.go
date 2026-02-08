package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/prompts"
)

func PromptWithModels() {
	simpleTemplate := prompts.NewPromptTemplate(
		"Write a {{.content_type}} about {{.subject}}!",
		[]string{"content_type", "subject"},
	)

	templateInput := map[string]any{

		"content_type": "poem",
		"subject":      "cats",
	}

	myLLM, err := anthropic.New(
		anthropic.WithToken(os.Getenv("ANTHROPIC_API_KEY")),
		anthropic.WithModel("claude-opus-4-20250514"),
	)
	if err != nil {
		fmt.Println("Error creating Anthropic client:", err)
		fmt.Println("Make sure ANTHROPIC_API_KEY environment variable is set")
		return
	}

	prompt, err := simpleTemplate.Format(templateInput)
	if err != nil {
		fmt.Println("Error formatting template:", err)
		return
	}

	response, err := llms.GenerateFromSinglePrompt(context.Background(), myLLM, prompt)
	if err != nil {
		fmt.Println("Error generating response:", err)
		return
	}

	fmt.Println(response)
}
