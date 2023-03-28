package chatgpt

import (
	"github.com/poorjobless/wechatbot/config"
	gpt35 "github.com/poorjobless/wechatbot/gpt-client"
)

func Completions(msg string) (string, error) {
	c := gpt35.NewClient(config.LoadConfig().ApiKey, config.LoadConfig().Proxy)
	req := &gpt35.Request{
		Model: gpt35.ModelGpt35Turbo,
		Messages: []*gpt35.Message{
			{
				Role:    gpt35.RoleUser,
				Content: msg,
			},
		},
	}

	resp, err := c.GetChat(req)
	if err != nil {
		panic(err)
	}

	return resp.Choices[0].Message.Content, err
}
