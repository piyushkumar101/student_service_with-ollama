package ollama

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"student_service/models"
)

type OllamaStreamResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}

func FetchStudentSummaryFromOllama(student models.Student) (string, error) {
	url := "http://localhost:11434/api/generate"
	payload := map[string]interface{}{
		"model":  "llama3",
		"prompt": fmt.Sprintf("Generate a summary for the student: Name: %s, Age: %d, Email: %s", student.Name, student.Age, student.Email),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	var summary string

	for scanner.Scan() {
		var streamResponse OllamaStreamResponse
		line := scanner.Text()
		log.Printf("Received stream line: %s", line)

		if err := json.Unmarshal([]byte(line), &streamResponse); err != nil {
			return "", fmt.Errorf("error unmarshalling stream response: %v", err)
		}

		summary += streamResponse.Response
		if streamResponse.Done {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading stream: %v", err)
	}

	return summary, nil
}
