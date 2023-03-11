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

	return resp.Choices[0].Message.Content, err
}
