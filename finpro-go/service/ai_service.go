package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIService interface {
	AnalyzeLearningStyle(ctx context.Context, prompt string) (string, error)
	Chat(ctx context.Context, userMessage string, learningProfile string) (string, error)
}

type aiService struct {
	apiKey string
	model  string
}

func NewAIService() AIService {
	key := os.Getenv("GEMINI_API_KEY")
	modelName := os.Getenv("GEMINI_MODEL")
	if modelName == "" {
		modelName = "gemini-1.5-flash" // Default safe model
	}

	if key == "" {
		log.Println("[AI] WARNING: GEMINI_API_KEY is empty!")
	} else {
		log.Printf("[AI] Gemini API Key detected (prefix: %s...)", key[:5])
		log.Printf("[AI] Using model: %s", modelName)
	}

	return &aiService{
		apiKey: key,
		model:  modelName,
	}
}

func (s *aiService) AnalyzeLearningStyle(ctx context.Context, prompt string) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(s.apiKey))
	if err != nil {
		log.Printf("[AI] Failed to create client: %v", err)
		return "", fmt.Errorf("failed to create Gemini client: %w", err)
	}
	defer client.Close()

	// Use configured model from .env
	model := client.GenerativeModel(s.model)
	model.SetTemperature(0.7)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("[AI] Model %s failed: %v", s.model, err)
		return "", err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		log.Printf("[AI] ERROR: Empty response candidates from Gemini")
		return "", fmt.Errorf("no response from Gemini")
	}

	var raw string
	for _, part := range resp.Candidates[0].Content.Parts {
		raw += fmt.Sprintf("%v", part)
	}

	raw = strings.TrimSpace(raw)
	
	// Robust JSON extraction
	start := strings.Index(raw, "{")
	end := strings.LastIndex(raw, "}")
	
	if start == -1 || end == -1 || end < start {
		log.Printf("[AI] ERROR: Failed to find JSON boundaries in response: %s", raw)
		return raw, nil 
	}

	cleaned := raw[start : end+1]
	return cleaned, nil
}

func (s *aiService) Chat(ctx context.Context, userMessage string, learningProfile string) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(s.apiKey))
	if err != nil {
		log.Printf("[AI Chat] Failed to create client: %v", err)
		return "", fmt.Errorf("failed to create Gemini client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel(s.model)
	model.SetTemperature(0.8)

	systemContext := fmt.Sprintf(`
You are Lumora AI, a friendly and motivational study assistant for university students.
You help students improve their self-regulated learning habits.

The student's learning profile is:
%s

Guidelines:
- Be concise, warm, and encouraging
- Give practical, actionable advice
- Stay focused on studying, learning strategies, time management, and academic goals
- If asked something unrelated to studying, politely redirect to academic topics
- Respond in the same language the student uses (Indonesian or English)
`, learningProfile)

	fullPrompt := fmt.Sprintf("%s\n\nStudent: %s\n\nLumora AI:", systemContext, userMessage)

	resp, err := model.GenerateContent(ctx, genai.Text(fullPrompt))
	if err != nil {
		log.Printf("[AI Chat] Model %s failed: %v", s.model, err)
		return "", err
	}

	if len(resp.Candidates) == 0 {
		return "", fmt.Errorf("no chat response from Gemini")
	}

	var result string
	for _, part := range resp.Candidates[0].Content.Parts {
		result += fmt.Sprintf("%v", part)
	}

	return strings.TrimSpace(result), nil
}
