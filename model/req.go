package model

type ReqBody struct {
	Model            string  `json:"model"`
	Messages         []Roles `json:"messages"`
	Temperature      float64 `json:"temperature,omitempty"`
	TopP             float64 `json:"top_p,omitempty"`
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`
	PresencePenalty  float64 `json:"presence_penalty,omitempty"`
	N                int     `json:"n,omitempty"`
	MaxTokens        int     `json:"max_tokens,omitempty"`
}

type Roles struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GptConfig struct {
	Temperature      float64 `json:"temperature,omitempty"`
	TopP             float64 `json:"top_p,omitempty"`
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`
	PresencePenalty  float64 `json:"presence_penalty,omitempty"`
	N                int     `json:"n,omitempty"`
	MaxTokens        int     `json:"max_tokens,omitempty"`
}
