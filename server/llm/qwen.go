package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iifiigii/gin"
	"io"
	"log"
	"net/http"
)

var alikey = "sk-xxxxxx"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Dialog struct {
	User string `json:"user"`
	Bot  string `json:"bot"`
}

type Input struct {
	Prompt  string     `json:"prompt,omitempty"`  // 使用指针表示 Prompt 字段可以为空
	History *[]Dialog  `json:"history,omitempty"` // 使用指针表示 History 字段可以为空
	Message *[]Message `json:"messages"`
}

type Params struct {
	TopP      *float64 `json:"top_p,omitempty"`         // 使用指针表示 TopP 字段可以为空
	TopK      *int     `json:"top_k,omitempty"`         // 使用指针表示 TopK 字段可以为空
	Seed      *int     `json:"seed,omitempty"`          // 使用指针表示 Seed 字段可以为空
	ResFormat *string  `json:"result_format,omitempty"` // 使用指针表示 ResFormat 字段可以为空
}

type RequestBody struct {
	Model  string  `json:"model"`
	Input  *Input  `json:"input"`                // 使用指针表示 History 字段可以为空
	Params *Params `json:"parameters,omitempty"` // 使用指针表示 Params 字段可以为空
}

type Output struct {
	FinishReason string `json:"finish_reason"`
	Text         string `json:"text"`
}

type Usage struct {
	TotalTokens  int `json:"total_tokens"`
	OutputTokens int `json:"output_tokens"`
}

type Response struct {
	Output Output `json:"output"`
	Usage  Usage  `json:"usage"`
	Id     string `json:"request_id"`
}

func Generagte(c *gin.Context) Response {
	url := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation" // 这里应填写真实的URL
	message := []Message{{Role: "system", Content: "请用英文回答问题！"}, {Role: "user", Content: "麦当劳好吃吗"}}
	resFormat := "message"
	reqBody := RequestBody{
		Model:  "qwen-14b-chat",
		Input:  &Input{Message: &message},
		Params: &Params{ResFormat: &resFormat},
	}
	jsonValue, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", alikey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	return result
}
