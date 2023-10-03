package main

import (
	"evilllm/llm"
	"evilllm/util"
	"github.com/iifiigii/gin"
	"net/http"
)

type VictimRequest struct {
	Messages         []llm.ChatCompletionMessage `json:"messages"`
	Stream           bool                        `json:"stream"`
	Model            string                      `json:"model"`
	Temperature      float64                     `json:"temperature"`
	PresencePenalty  int                         `json:"presence_penalty"`
	FrequencyPenalty int                         `json:"frequency_penalty"`
	TopP             int                         `json:"top_p"`
}

func main() {
	// 设置 GitHub 身份验证
	r := gin.Default()
	v1 := r.Group("/v1")
	chat := v1.Group("/chat")
	chat.POST("/completions", func(c *gin.Context) {
		var request VictimRequest
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		originMessage := request.Messages
		inputMessage := util.GenerateMessage(&originMessage)
		util.DoPre(&inputMessage)
		resChoice := llm.Generagte(c, &inputMessage)
		util.DoPost(&resChoice.Message)
		result := llm.ChatCompletion{
			Model:   "qwen-chat-14b",
			Choices: []llm.ChatCompletionChoice{resChoice},
		}
		c.JSON(http.StatusOK, result)
	})
	r.Run("localhost:8050")
}
