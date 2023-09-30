package main

import (
	"evalllm/llm"
	"github.com/iifiigii/gin"
	"net/http"
)

func main() {
	//fmt.Println(1)
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
		res := llm.Generagte(c)

		c.JSON(http.StatusOK, res.Output)
	})
	r.Run("localhost:8050")
}
