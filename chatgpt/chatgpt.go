package chatgpt

import (
	"github.com/poorjobless/wechatbot/config"
	gpt35 "github.com/poorjobless/wechatbot/gpt-client"
)

var Cache CacheInterface

func init() {
	Cache = GetSessionCache()
}

func Completions(session, msg string) (string, error) {
	ms := Cache.GetMsg(session)
	message := &gpt35.Message{
		Role:    "user",
		Content: msg,
	}
	ms = append(ms, *message)
	c := gpt35.NewClient(config.LoadConfig().ApiKey, config.LoadConfig().Proxy)
	req := &gpt35.Request{
		Model:       gpt35.ModelGpt35Turbo,
		Messages:    ms,
		MaxTokens:   1000,
		Temperature: 0.7,
	}

	resp, err := c.GetChat(req)
	if err != nil {
		panic(err)
	}
	ms = append(ms, resp.Choices[0].Message)
	Cache.SetMsg(session, ms)

	return resp.Choices[0].Message.Content, err
}
