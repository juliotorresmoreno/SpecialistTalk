package services

import (
	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func NewChatGPTClient() *gogpt.Client {
	conf := configs.GetConfig().ChatGPT3
	c := gogpt.NewClient(conf.ApiKey)

	return c
}
