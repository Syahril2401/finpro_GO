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
}

func NewAIService() AIService {
	key := os.Getenv("GEMINI_API_KEY")
	if key == "" {
		log.Println("[AI] WARNING: GEMINI_API_KEY is empty!")
	} else {
		log.Printf("[AI] Gemini API Key detected (prefix: %s...)", key[:5])
	}
	return &aiService{
		apiKey: key,
	}
}

func (s *aiService) AnalyzeLearningStyle(ctx context.Context, prompt string) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(s.apiKey))
	if err != nil {
		log.Printf("[AI] Failed to create client: %v", err)
		return "", fmt.Errorf("failed to create Gemini client: %w", err)
	}
	defer client.Close()

	// Coba gemini-1.5-flash dulu
	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(0.7)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("[AI] Flash failed: %v. Trying gemini-pro...", err)
		// Fallback ke gemini-pro (lebih stabil di versi API lama)
		model = client.GenerativeModel("gemini-pro")
		resp, err = model.GenerateContent(ctx, genai.Text(prompt))
		if err != nil {
			log.Printf("[AI] All models failed: %v", err)
			return "", err
		}
		log.Println("[AI] Success using gemini-pro")
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response from Gemini")
	}

	var raw string
	for _, part := range resp.Candidates[0].Content.Parts {
		raw += fmt.Sprintf("%v", part)
	}

	raw = strings.TrimSpace(raw)
	raw = strings.TrimPrefix(raw, "```json")
	raw = strings.TrimPrefix(raw, "```")
	raw = strings.TrimSuffix(raw, "```")
	raw = strings.TrimSpace(raw)

	return raw, nil
}

func (s *aiService) Chat(ctx context.Context, userMessage string, learningProfile string) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(s.apiKey))
	if err != nil {
		log.Printf("[AI Chat] Failed to create client: %v", err)
		return "", fmt.Errorf("failed to create Gemini client: %w", err)
	}
	defer client.Close()

	// Coba gemini-1.5-flash
	model := client.GenerativeModel("gemini-1.5-flash")
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
		log.Printf("[AI Chat] Flash failed: %v. Trying gemini-pro...", err)
		// Fallback ke gemini-pro
		model = client.GenerativeModel("gemini-pro")
		resp, err = model.GenerateContent(ctx, genai.Text(fullPrompt))
		if err != nil {
			log.Printf("[AI Chat] All models failed: %v", err)
			return "", err
		}
		log.Println("[AI Chat] Success using gemini-pro")
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
