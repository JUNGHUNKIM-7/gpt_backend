package model

type Body struct {
	Q      string     `json:"q"`
	Config *GptConfig `json:"config,omitempty"`
}

type Response struct {
	A string `json:"a"`
}

type Env struct {
	GptToken string
}

var Environment *Env
