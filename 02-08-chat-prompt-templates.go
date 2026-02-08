package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/prompts"
)

func ChatPromptTemplates() {

	systemPromptTemplate := prompts.NewSystemMessagePromptTemplate(
		"You're to give your respopnses in {{.language}}",
		[]string{"language"},
	)

	humanMessageTemplate := prompts.NewHumanMessagePromptTemplate(
		"Translate the following phrase: {{.phrase}}",
		[]string{"phrase"},
	)

	chatPromptTemplate := prompts.NewChatPromptTemplate(
		[]prompts.MessageFormatter{
			systemPromptTemplate,
			humanMessageTemplate,
		},
	)

	formatted_chat_prompt, err := chatPromptTemplate.Format(
		map[string]any{
			"language": "French",
			"phrase":   "Thank You!",
		},
	)
	if err != nil {
		fmt.Println("Error formatting chat prompt:", err)
		return
	}

	//fmt.Println(formatted_chat_prompt)

	ctx := context.Background()

	llm, err := anthropic.New(
		anthropic.WithToken(os.Getenv("ANTHROPIC_API_KEY")),
		anthropic.WithModel("claude-opus-4-20250514"),
	)
	if err != nil {
		fmt.Println("Error creating Anthropic client:", err)
		fmt.Println("Make sure ANTHROPIC_API_KEY environment variable is set")
		return
	}

	response, err := llms.GenerateFromSinglePrompt(ctx, llm, formatted_chat_prompt)
	if err != nil {
		fmt.Println("Error generating response:", err)
		return
	}

	fmt.Println(response)
}
