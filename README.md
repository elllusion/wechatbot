# 微信机器人（ChatGPT官方API版）
最近chatGPT异常火爆，想到将其接入到个人微信是件比较有趣的事，所以有了这个项目。本项目使用了OpenAI最新（2023-03-02）发布的ChatGPT在使用的大型语言模型gpt-3.5-turbo以其API和基于[openwechat](https://github.com/eatmoreapple/openwechat)开发。

本项目修改自 @djun 的 [wechatbot](https://github.com/djun/wechatbot)

### 目前实现了以下功能
 + 群聊@回复
 + 私聊回复
 + 自动通过回复
 
#### 注意： 若在使用过程中有任何技术问题请自行处理。本人只确保发布的源码是可运行的（可运行不代表要符合任何人的心意，本项目属于实验性质的demo级项目，若有任何不快请自行修改代码）。

# 注册OpenAI开发者帐号
OpenAI开发者帐号注册可以参考[这里](https://juejin.cn/post/7173447848292253704)

# 安装使用

## 获取项目
`git clone https://github.com/poorjobless/wechatbot.git`

## 进入项目目录
`cd wechatbot`

## 复制配置文件
linux下执行 `cp config.dev.json config.json`

windows下执行 `copy config.dev.json config.json`

## 启动项目
`go run main.go`

启动前需替换config中的api_key

### 若想实现连续对话功能请参考以下内容：

为什么要提供一个对话列表呢？因为 API 调用都是单次的接口请求，不会自动记录之前的聊天信息，没有上下文，要让 ChatGPT 在单次的请求中了解你的上下文，就需要提供这样一个完整的对话列表，比如这样一个对话
```python
import openai

openai.ChatCompletion.create(
  model="gpt-3.5-turbo",
  messages=[
        {"role": "system", "content": "You are a helpful assistant."},
        {"role": "user", "content": "Who won the world series in 2020?"},
        {"role": "assistant", "content": "The Los Angeles Dodgers won the World Series in 2020."},
        {"role": "user", "content": "Where was it played?"}
    ]
)
```
每条对话消息都需要提供角色和内容，角色分三种：系统（system）、用户（user）、助手（assistant）。
<br> system：系统消息相当于一个管理员，可以设置助手的行为和特征。在上面的例子中，助手被指示你是一个有用的助手。
<br> user：用户消息就是我们自己，可以由用户发出提问，或者直接让开发者提前内置一些 Prompts 指令。一些可以参考的 [ChatGPT Prompts](https://chatopenai.pro/chatgpt-prompts/)
<br> assistant：助手消息就是 ChatGPT API 在之前提供的回复，在这里存储起来。这个回复也可以自己修改或者自己编一段对话出来，来让整个对话更通顺。
<br> 如果不需要对话的话，只需要提供单个的user消息即可，就像刚刚 Python 代码里演示的。


# 鸣谢
+ @djun
+ @eatmoreapple
