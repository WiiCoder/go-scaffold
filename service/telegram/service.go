package telegram

import tb "gopkg.in/tucnak/telebot.v2"

var _ Service = (*service)(nil)

type Service interface {
	// Send 自定义发送模版类型消息
	Send(targets []string, message string, msgMode tb.ParseMode) bool
	// SendModeMarkdownV2 发送Markdown V2版本消息
	SendModeMarkdownV2(targets []string, message string) bool
	// SendModeMarkdown 发送Markdown消息
	SendModeMarkdown(targets []string, message string) bool
	// SendModeHTML 发送 HTML 消息
	SendModeHTML(targets []string, message string) bool
	// SendModeDefault 发送普通文本消息
	SendModeDefault(targets []string, message string) bool
}

type service struct{}

func New() Service {
	return &service{}
}
