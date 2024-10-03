package chatbot

import (
	"backend-mental-guardians/configs"
	"backend-mental-guardians/entities/chatbot"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ChatbotUsecase struct{}

func NewChatbotUsecase() *ChatbotUsecase {
	return &ChatbotUsecase{}
}

func (c *ChatbotUsecase) GetResult(newMessage string) (string, error) {
	messages := []map[string]string{
		{"role": "system", "content": "You are an AI assistant that helps classify the user's mental state based on their mood and the story or description they provide. Only respond with either 'good' or 'bad' depending on the user's input and the story they tell."},
	}

	messages = append(messages, map[string]string{"role": "user", "content": newMessage})

	// Create the request payload
	openAIRequest := chatbot.OpenAIRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}

	// Convert the request payload to JSON
	requestBody, err := json.Marshal(openAIRequest)
	if err != nil {
		return "", err
	}

	openApiKey := configs.InitConfigKeyChatbot()

	// Send the request to OpenAI API
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openApiKey) // Use your OpenAI API key here

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Parse the response from OpenAI API
	var aiResponse chatbot.OpenAIResponse
	err = json.NewDecoder(resp.Body).Decode(&aiResponse)
	if err != nil {
		return "", err
	}

	if len(aiResponse.Choices) == 0 {
		return "", fmt.Errorf("no completions found")
	}

	return aiResponse.Choices[0].Message.Content, nil
}