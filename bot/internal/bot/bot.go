package bot

import (
	"burmachineBot/internal/config"
	"fmt"
	"github.com/rs/zerolog"
)

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Bot struct {
	Logger *zerolog.Logger
	Bot    *tgbotapi.BotAPI
	Config config.Conf
}

func (b *Bot) BotRegistration(apiKey string) error {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return err
	}
	b.Bot = bot
	return nil
}

func (b *Bot) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Bot.GetUpdatesChan(u)
	b.Logger.Info().Msg("bot started")
	for update := range updates {

		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := b.Bot.Request(callback); err != nil {
				b.Logger.Fatal().Err(err).Msg("callback error")
			}
			if update.CallbackQuery.Data == "get_currency" {
				// resp
			}
		} else if update.Message != nil {
			b.Logger.Info().Msg(fmt.Sprintf("event = user:%s + mes:%s", update.Message.From.UserName, update.Message.Text))

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие")
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			msg.ReplyMarkup = inKeyboard_user

			_, err := b.Bot.Send(msg)
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}
