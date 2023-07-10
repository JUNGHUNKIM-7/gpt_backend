package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/JUNGHUNKIM-7/model"
)

func GetCompletion(sysRole, cliRole model.Roles, config *model.GptConfig) (result string) {
	client := &http.Client{
		Timeout: 5 * time.Minute,
	}
	var requestBody model.ReqBody

	if config != nil {
		requestBody = model.ReqBody{
			Model:            "gpt-4",
			Messages:         *RolesNew(sysRole, cliRole),
			Temperature:      config.Temperature,
			TopP:             config.TopP,
			N:                config.N,
			FrequencyPenalty: config.FrequencyPenalty,
			PresencePenalty:  config.PresencePenalty,
			MaxTokens:        config.MaxTokens,
		}
	} else {
		requestBody = model.ReqBody{
			Model:    "gpt-4",
			Messages: *RolesNew(sysRole, cliRole),
		}
	}

	jsonBody, err := json.Marshal(requestBody)
	fmt.Println(string(jsonBody))

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	if model.Environment != nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", model.Environment.GptToken))
	} else {
		log.Fatal("env not initalized")
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var resbody model.ResBody
	json.Unmarshal(body, &resbody)

	result = resbody.Choices[0].Message.Content
	return
}

func RolesNew(sysRole, cliRole model.Roles) *[]model.Roles {
	return &[]model.Roles{
		{
			Role:    sysRole.Role,
			Content: sysRole.Content,
		},
		{
			Role:    cliRole.Role,
			Content: cliRole.Content,
		},
	}
}
