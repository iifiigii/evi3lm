package main

import (
	"encoding/json"
	"fmt"
	"github.com/iifiigii/gin"
	"net/http"
)

func main() {
	//client := github.NewClient(nil).WithAuthToken("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")
	//// create the repository\
	//repo, res, err := client.Repositories.Create(
	//	context.Background(),
	//	"",
	//	&github.Repository{
	//		Name:    github.String("gin"),
	//		Private: github.Bool(false),
	//	},
	//)
	//if err != nil {
	//	fmt.Println(*res)
	//	fmt.Println(repo)
	//	fmt.Println(err)
	//}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/v1")
	chat := v1.Group("/chat")
	chat.POST("/completions", func(c *gin.Context) {
		//res := llm.Generagte(c)

		result := make(map[string]interface{})
		body := `
		{
			"id": "chatcmpl-123",
				"object": "chat.completion",
				"created": 1677652288,
				"model": "gpt-3.5-turbo-0613",
				"choices": [{
					"index": 0,
					"message":{
						"role":"assistant",
						"content": "\n\nHello there, how may I assist you today?"
					},
					"finish_reason": "stop"
			}],
			"usage": {
			"prompt_tokens": 9,
			"completion_tokens": 12,
			"total_tokens": 21
			}
		}`

		err := json.Unmarshal([]byte(body), &result)
		fmt.Println(err)
		c.JSON(http.StatusOK, result)
		fmt.Println(result)
	})
	r.Run("localhost:8050")
}
