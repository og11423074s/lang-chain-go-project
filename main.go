package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		log.Fatal("ANTHROPIC_API_KEY environment variable is not set")
	}

	llm, err := anthropic.New(
		anthropic.WithToken(apiKey),
		anthropic.WithModel("claude-opus-4-20250514"),
	)
	if err != nil {
		log.Fatal(err)
	}

	prompt := "What is the capital of France?"

	completion, err := llms.GenerateFromSinglePrompt(
		context.Background(),
		llm,
		prompt)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)
}
