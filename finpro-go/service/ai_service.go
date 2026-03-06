package service

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIService interface {
	AnalyzeLearningStyle(ctx context.Context, assessmentData string) (string, error)
}

type aiService struct {
	apiKey string
}

func NewAIService() AIService {
	return &aiService{
		apiKey: os.Getenv("GEMINI_API_KEY"),
	}
}

func (s *aiService) AnalyzeLearningStyle(ctx context.Context, assessmentData string) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(s.apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	prompt := fmt.Sprintf(`
		Based on the following self-assessment responses from a student:
		%s
		
		Please determine their learning style (e.g., Visual, Auditory, Kinesthetic, or Read/Write) 
		and provide specific recommendations on how they should study to improve their performance.
		Format the response as a JSON string with two keys: "style" and "recommendations".
	`, assessmentData)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 {
		return "", fmt.Errorf("no response from AI")
	}

	var result string
	for _, part := range resp.Candidates[0].Content.Parts {
		result += fmt.Sprintf("%v", part)
	}

	return result, nil
}
