package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var inKeyboard_user = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Курс валют", "see_all_events_user"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Ничего", "Chiken-box_user"),
	),
)
