package AiServices

import (
	"fmt"
)

type AIService interface {
	Reply(prompt string) (string, error)
}

type aiService struct{}

func NewAIService() AIService {
	return &aiService{}
}

func (a *aiService) Reply(prompt string) (string, error) {
	response := fmt.Sprintf("You said: %s. AI bot replied!", prompt)
	return response, nil
}
