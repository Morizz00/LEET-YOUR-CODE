package scrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Agent struct {
	APIKey string
}

func NewAgent(apiKey string) *Agent {
	return &Agent{APIKey: apiKey}
}

func (a *Agent) GetSolution(problem string) (string, error) {
	url := "https://openrouter.ai/api/v1/chat/completions"

	payload := map[string]interface{}{
		"model": "openai/gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": fmt.Sprintf("Write clean Go code to solve this LeetCode problem and explain each line/step with comments: %s", problem)},
		},
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no content returned from OpenRouter")
	}

	return result.Choices[0].Message.Content, nil
}
