package kb

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

var InlineKeyboard *telego.InlineKeyboardMarkup = tu.InlineKeyboard(
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("Посмотреть кроссовки").
			WithCallbackData("callback_1"),
		tu.InlineKeyboardButton("Корзина").
			WithCallbackData("callback_2"),
	),
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("Связаться с менеджером").WithURL("https://example.com"),
	),
)

var InlineKeyboardAdmin *telego.InlineKeyboardMarkup = tu.InlineKeyboard(
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("Поменять приветствие:").
			WithCallbackData("callback_hello"),
		tu.InlineKeyboardButton("Корзина").
			WithCallbackData("callback_2"),
	),
)

var InlineKeyboard2 *telego.InlineKeyboardMarkup = tu.InlineKeyboard(
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("Puma").
			WithCallbackData("callback_puma"),
		tu.InlineKeyboardButton("Adidas").
			WithCallbackData("callback_adidas"),
		tu.InlineKeyboardButton("Nike").
			WithCallbackData("callback_nike"),
	),
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("Назад").WithCallbackData("callback_0"),
	),
)

var InlineKeyboard3 *telego.InlineKeyboardMarkup = tu.InlineKeyboard(
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("Назад").
			WithCallbackData("callback_0"),
	),
)

var InlineKeyboardPuma *telego.InlineKeyboardMarkup = tu.InlineKeyboard(
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("Puma 41 size").
			WithCallbackData("callback_41size"),
		tu.InlineKeyboardButton("Puma 42 size").
			WithCallbackData("callback_42size"),
		tu.InlineKeyboardButton("Puma 43 size").
			WithCallbackData("callback_puma43size"),
		tu.InlineKeyboardButton("Главное меню").
			WithCallbackData("callback_0"),
	),
)
