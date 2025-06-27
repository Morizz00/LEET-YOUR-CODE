package scrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
		"model": "meta-llama/llama-3.1-8b-instruct:free",
		"messages": []map[string]string{
			{"role": "user", "content": fmt.Sprintf("Write clean Go code to solve this LeetCode problem and explain each line/step with comments: %s", problem)},
		},
		"max_tokens":  1000,
		"temperature": 0.7,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshaling payload: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(responseBody))
	}

	fmt.Printf("Raw API Response: %s\n", string(responseBody))

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Error struct {
			Message string `json:"message"`
			Type    string `json:"type"`
		} `json:"error"`
	}

	if err := json.Unmarshal(responseBody, &result); err != nil {
		return "", fmt.Errorf("error parsing JSON response: %v", err)
	}

	if result.Error.Message != "" {
		return "", fmt.Errorf("API error: %s (type: %s)", result.Error.Message, result.Error.Type)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from OpenRouter")
	}

	if result.Choices[0].Message.Content == "" {
		return "", fmt.Errorf("empty content returned from OpenRouter")
	}

	return result.Choices[0].Message.Content, nil
}

