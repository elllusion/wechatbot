package chatgpt

import (
	"github.com/poorjobless/wechatbot/config"
)

func Completions(msg string) (string, error) {

	c := NewClient(config.LoadConfig().ApiKey)

	req := &Request{
		Model: ModelGpt35Turbo,
		Messages: []*Message{
			{
				Role:    RoleUser,
				Content: msg,
			},
		},
	}
	resp, err := c.GetChat(req)

	println(resp.Choices[0].Message.Content)
	println(resp.Usage.PromptTokens)
	println(resp.Usage.CompletionTokens)
	println(resp.Usage.TotalTokens)

	return resp.Choices[0].Message.Content, err
}
