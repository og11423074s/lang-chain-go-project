package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// apiKey := os.Getenv("ANTHROPIC_API_KEY")
	// if apiKey == "" {
	// 	log.Fatal("ANTHROPIC_API_KEY environment variable is not set")
	// }

	// llm, err := anthropic.New(
	// 	anthropic.WithToken(apiKey),
	// 	anthropic.WithModel("claude-opus-4-20250514"),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// prompt := "What is the capital of France?"

	// completion, err := llms.GenerateFromSinglePrompt(
	// 	context.Background(),
	// 	llm,
	// 	prompt)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(completion)

	// 02-01
	//stringPromptTemplates()

	//02-02
	//StandardTemplateDefinition()

	//02-03
	//JinjaPromptTemplates()

	//02-04
	//MultiLineTemplates()

	//02-05
	//PartialVariablesPrompts()

	//02-06
	//PromptWithModels()

	//02-08
	ChatPromptTemplates()
}
