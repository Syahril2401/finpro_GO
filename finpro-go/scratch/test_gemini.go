package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY is empty")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text("Say 'Connection Successful' in one line."))
	if err != nil {
		log.Fatal(err)
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Printf("Response: %v\n", part)
	}
}
