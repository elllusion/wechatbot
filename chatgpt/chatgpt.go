package chatgpt

import (
	"github.com/poorjobless/wechatbot/config"
	gpt35 "github.com/AlmazDelDiablo/gpt3-5-turbo-go"
)

func Completions(msg string) (string, error) {
	c := gpt35.NewClient(config.LoadConfig().ApiKey)
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
