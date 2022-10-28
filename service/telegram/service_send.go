package telegram

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/wiiCoder/go-scaffold/global"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Send 发送 msgMode 类型的 message 给 targets， targets 为 tg 的 chat_id
func (s *service) Send(targets []string, message string, msgMode tb.ParseMode) bool {
	// msg Mode, 优先 tb.ModeMarkdownV2
	opt := &tb.SendOptions{ParseMode: msgMode}

	for _, t := range targets {
		to, err := strconv.Atoi(t)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
		_, err = global.TgBot.Send(tb.ChatID(to), message, opt)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
	}
	return true
}

func (s *service) SendModeMarkdownV2(targets []string, message string) bool {
	opt := &tb.SendOptions{ParseMode: tb.ModeMarkdownV2}

	for _, t := range targets {
		to, err := strconv.Atoi(t)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
		_, err = global.TgBot.Send(tb.ChatID(to), message, opt)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
	}
	return true
}

func (s *service) SendModeMarkdown(targets []string, message string) bool {
	opt := &tb.SendOptions{ParseMode: tb.ModeMarkdown}

	for _, t := range targets {
		to, err := strconv.Atoi(t)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
		_, err = global.TgBot.Send(tb.ChatID(to), message, opt)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
	}
	return true
}

func (s *service) SendModeHTML(targets []string, message string) bool {
	opt := &tb.SendOptions{ParseMode: tb.ModeHTML}

	for _, t := range targets {
		to, err := strconv.Atoi(t)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
		_, err = global.TgBot.Send(tb.ChatID(to), message, opt)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
	}
	return true
}

func (s *service) SendModeDefault(targets []string, message string) bool {
	opt := &tb.SendOptions{ParseMode: tb.ModeDefault}

	for _, t := range targets {
		to, err := strconv.Atoi(t)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
		_, err = global.TgBot.Send(tb.ChatID(to), message, opt)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
	}
	return true
}
