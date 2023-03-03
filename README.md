# 微信机器人（ChatGPT官方API版）
最近chatGPT异常火爆，想到将其接入到个人微信是件比较有趣的事，所以有了这个项目。本项目使用了OpenAI最新（2023-03-02）发布的ChatGPT在使用的大型语言模型gpt-3.5-turbo以其API和基于[openwechat](https://github.com/eatmoreapple/openwechat)开发。

本项目修改自 @djun 的 [wechatbot](https://github.com/djun/wechatbot)

### 目前实现了以下功能
 + 群聊@回复
 + 私聊回复
 + 自动通过回复
 
# 注意： 若在使用过程中有任何技术问题请自行处理。本人不提供任何免费的技术服务和咨询，本人只确保发布的源码是可运行的。

# 注册OpenAI开发者帐号
OpenAI开发者帐号注册可以参考[这里](https://juejin.cn/post/7173447848292253704)

# 安装使用
````
# 获取项目
git clone xxx

# 进入项目目录
cd wechatbot

# 复制配置文件
linux下执行 `cp config.dev.json config.json`

windows下执行 `copy config.dev.json config.json`

# 修改配置文件 config.json
{
  // OpenAI api key
  "api_key": "your api key",
  // 是否自动通过好友申请
  "auto_pass": false,
  // 不自动回复的用户名列表（防止一些系统通知造成持续回复）
  "no_reply_user_list": [
    "微信团队"
  ]
}

# 启动项目
go run main.go
````

# 鸣谢
+ @djun
+ @eatmoreapple
