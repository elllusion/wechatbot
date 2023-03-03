package handlers

import (
	"log"
	"strings"

	"github.com/eatmoreapple/openwechat"
	"github.com/poorjobless/wechatbot/chatgpt"
	"github.com/poorjobless/wechatbot/config"
)

var _ MessageHandlerInterface = (*UserMessageHandler)(nil)

// UserMessageHandler 私聊消息处理
type UserMessageHandler struct {
}

// handle 处理消息
func (g *UserMessageHandler) handle(msg *openwechat.Message) error {
	// 判断是否是不回复用户
	u, err := msg.Sender()
	if err != nil {
		log.Printf("chatgpt request error: %v \n", err)
		return nil
	}
	NoReplyUserList := config.LoadConfig().NoReplyUserList
	for i := 0; i < len(NoReplyUserList); i++ {
		if NoReplyUserList[i] == u.NickName {
			return nil
		}
	}
	 
	// 文本消息才回复
	if msg.IsText() {
		return g.ReplyText(msg)
	}
	return nil
}

// NewUserMessageHandler 创建私聊处理器
func NewUserMessageHandler() MessageHandlerInterface {
	return &UserMessageHandler{}
}

// ReplyText 发送文本消息到群
func (g *UserMessageHandler) ReplyText(msg *openwechat.Message) error {
	// 接收私聊消息
	sender, err := msg.Sender()
	log.Printf("Received User %v Text Msg : %v", sender.NickName, msg.Content)

	// 向GPT发起请求
	requestText := strings.TrimSpace(msg.Content)
	requestText = strings.Trim(msg.Content, "\n")
	reply, err := chatgpt.Completions(requestText)
	if err != nil {
		log.Printf("chatgpt request error: %v \n", err)
		msg.ReplyText("机器人神了，我一会发现了就去修。")
		return err
	}
	if reply == "" {
		return nil
	}

	// 回复用户
	reply = strings.TrimSpace(reply)
	reply = strings.Trim(reply, "\n")
	_, err = msg.ReplyText(reply)
	if err != nil {
		log.Printf("response user error: %v \n", err)
	}
	return err
}
