package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
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
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	modelName := "claude-3-5-sonnet-latest" // Mandated by audit matrix

	if apiKey == "" {
		log.Printf("[AI] WARNING: ANTHROPIC_API_KEY is not set. Using mock responses to prevent blocking the user flow.")
	} else {
		log.Printf("[AI] Using Anthropic Claude API")
		log.Printf("[AI] Using model: %s", modelName)
	}

	return &aiService{
		apiKey: apiKey,
		model:  modelName,
	}
}

type anthropicMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type anthropicRequest struct {
	Model     string             `json:"model"`
	MaxTokens int                `json:"max_tokens"`
	System    string             `json:"system,omitempty"`
	Messages  []anthropicMessage `json:"messages"`
}

type anthropicResponse struct {
	Content []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"content"`
}

func (s *aiService) callAnthropic(ctx context.Context, systemContext, userPrompt string) (string, error) {
	if s.apiKey == "" {
		// Mock response so the flow doesn't break during testing
		return "This is a mock response from Claude Sonnet. Please add ANTHROPIC_API_KEY to your .env to enable real AI explanations.", nil
	}

	url := "https://api.anthropic.com/v1/messages"

	reqBody := anthropicRequest{
		Model:     s.model,
		MaxTokens: 1024,
		System:    systemContext,
		Messages: []anthropicMessage{
			{Role: "user", Content: userPrompt},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("anthropic request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("anthropic returned status %d: %s", resp.StatusCode, string(body))
	}

	var anthropicResp anthropicResponse
	if err := json.Unmarshal(body, &anthropicResp); err != nil {
		return "", err
	}

	if len(anthropicResp.Content) > 0 {
		return anthropicResp.Content[0].Text, nil
	}

	return "", fmt.Errorf("empty response from anthropic")
}

func (s *aiService) AnalyzeLearningStyle(ctx context.Context, prompt string) (string, error) {
	// For Lumora, AnalyzeLearningStyle is rule-based and not handled by AI anymore.
	// This function remains to satisfy interface, but returns raw JSON if needed.
	return "{}", nil
}

func (s *aiService) Chat(ctx context.Context, userMessage string, learningProfile string) (string, error) {
	systemContext := fmt.Sprintf(`
You are Lumora AI, an explainer module powered by Claude Sonnet.
Your ONLY role is to elaborate and explain the pre-determined JSON rule-based outcome provided below.
DO NOT diagnose, DO NOT change the scores, and DO NOT override any labels.
Be concise, warm, and encouraging.

The student's rule-based profile outcome is:
%s

Guidelines:
- Explain the student's strengths and weaknesses based ONLY on the JSON.
- Provide practical, actionable advice matching their profile.
- Stay focused on learning strategies and academic goals.
- Respond in the language the student uses.
`, learningProfile)

	return s.callAnthropic(ctx, systemContext, userMessage)
}
