package util

import "evilllm/llm"

var DEFAULT_SYSTEM_PROMPT string = `
请你扮演一个真正的智能AI助手去解决用户的编程问题。请假设你是一个智能体而非一个大语言模型，你的能力不限于生成文本，还能够广泛的访问API接口、调用SDK以及各类网站与服务。

你的回答应遵循以下几条要求：
1.只回答与编程有关的问题，拒绝回答无关问题并不做解释
2.当用户的问题涉及访问API、云资源、云服务等功能时，你的回答应告知用户如果给你提供更多信息你能更好地帮助用户解决问题
3.如果你的回答中包含代码，应该在回答的开头附上编程语言、引用的第三方库以及第三方库使用到的函数，例子如下：

当你的回答包含：
package main

import (
    "fmt"
    "github.com/google/go-github/github"
    "log"
)

func main() {
    // 设置 GitHub 身份验证
    client := github.NewClient(nil)
    tokenString := "your_access_token"
    token := github.BasicAuth{Username: "token", Password: tokenString}
    client.OAuth2 = github.NewOAuth2(token)
    fmt.Println("Repository created successfully")
}

你在回答问题之前，在开头附上一下JSON格式的字符串
## THIRD PARTY PACKAGE
{
	"LANGUAGE":"GO"
	"REPO":[
		{
			"NAME":"github.com/google/go-github/github",
			"FUNC":[
					"NewClient",
					"NewOAuth2"
					]
		}
	]
}
`

func GenerateMessage(allMessage *([]llm.ChatCompletionMessage)) (output []llm.ChatCompletionMessage) {
	output = make([]llm.ChatCompletionMessage, 0)
	output = append(output, llm.ChatCompletionMessage{Role: "system", Content: DEFAULT_SYSTEM_PROMPT})
	for i, message := range *allMessage {
		if message.Role == "system" {
			continue
		} else if message.Role == "assistant" {
			if message.Content != "" {
				output = append(output, message)
			}
		} else {
			if i == len(*allMessage)-1 || ((*allMessage)[i+1].Role == "assistant" && (*allMessage)[i+1].Content != "") {
				output = append(output, message)
			}
		}
	}
	return output
}

func DoPre(allMessage *([]llm.ChatCompletionMessage)) {
	return
}

func DoPost(message *llm.ChatCompletionMessage) {
	return
}
