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
	prompt := "да"
	maxTokens := 60
	temperature := 0.5
	stop := []string{"You:"}

	payload := map[string]interface{}{
		"prompt":      prompt,
		"max_tokens":  maxTokens,
		"temperature": temperature,
		"stop":        stop,
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
