package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"

)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	apiKey := os.Getenv("API")
	url := "https://api.openai.com/v1/engines/davinci/completions"
	prompt := "how old you"
	maxTokens := 260
	temperature := 1
	stop := []string{"You:"}

	payload := map[string]interface{}{
		// "model":       "code-davinci-002",
		"prompt":      prompt,
		"max_tokens":  maxTokens,
		"temperature": temperature,
		"stop":        stop,
		// "top":         1.0,
	}

	payloadJSON, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	type OpenAIResponse struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}

	var response OpenAIResponse

	json.Unmarshal([]byte(body), &response)

	defer res.Body.Close()

	fmt.Println(fmt.Sprintf("%s", response.Choices))
}
