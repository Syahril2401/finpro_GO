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
	"strings"
	"time"
)

type AIService interface {
	AnalyzeLearningStyle(ctx context.Context, prompt string) (string, error)
	Chat(ctx context.Context, userMessage string, learningProfile string) (string, error)
}

type aiService struct {
	baseUrl string
	model   string
}

func NewAIService() AIService {
	baseUrl := os.Getenv("OLLAMA_BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost:11434"
	}
	modelName := os.Getenv("OLLAMA_MODEL")
	if modelName == "" {
		modelName = "llama3"
	}

	log.Printf("[AI] Using Ollama at %s", baseUrl)
	log.Printf("[AI] Using model: %s", modelName)

	return &aiService{
		baseUrl: baseUrl,
		model:   modelName,
	}
}

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
	Format string `json:"format,omitempty"`
}

type ollamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func (s *aiService) callOllama(ctx context.Context, prompt string) (string, error) {
	url := fmt.Sprintf("%s/api/generate", strings.TrimSuffix(s.baseUrl, "/"))

	reqBody := ollamaRequest{
		Model:  s.model,
		Prompt: prompt,
		Stream: false,
		Format: "json",
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

	// Ollama can be slow on some hardware, so use a longer timeout
	client := &http.Client{Timeout: 90 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ollama request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ollama returned status %d: %s", resp.StatusCode, string(body))
	}

	var ollamaResp ollamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return "", err
	}

	return ollamaResp.Response, nil
}

func (s *aiService) AnalyzeLearningStyle(ctx context.Context, prompt string) (string, error) {
	raw, err := s.callOllama(ctx, prompt)
	if err != nil {
		log.Printf("[AI] Ollama analysis failed: %v", err)
		return "", err
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

	return s.callOllama(ctx, fullPrompt)
}
