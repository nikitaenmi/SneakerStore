package main

import (
	"fmt"
	"os"

	DAL "t/DAL"
	token "t/botToken"
	kb "t/keyboard"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {

	DAL.Migration()

	bot, err := telego.NewBot(token.BotToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)
	defer bot.StopLongPolling()
	Hellostring := "Приветствую в магазине кроссовок"
	for update := range updates {

		bh.Handle(func(bot *telego.Bot, updateChan telego.Update) { //Start handle

			_, _ = bot.SendMessage(tu.Message(
				tu.ID(update.Message.Chat.ID),
				Hellostring,
			).WithReplyMarkup(kb.InlineKeyboard),
			)

			DAL.IdAdd(bot, update) // Добавление ID в БД

			photo := tu.Photo(
				// Chat ID as String (target username)
				tu.ID(update.Message.Chat.ID),

				// Send using file from disk
				tu.File(mustOpen("img/beautifullphoto.jpg")),
			).WithCaption("Hello")

			// Sending photo
			_, _ = bot.SendPhoto(photo)

		}, th.CommandEqual("start"))

		bh.Handle(func(bot *telego.Bot, updateChan telego.Update) { //Start handle

			_, _ = bot.SendMessage(tu.Message(
				tu.ID(updateChan.Message.Chat.ID),
				"Вы открыли режим администратора",
			).WithReplyMarkup(kb.InlineKeyboardAdmin),
			)

		}, th.CommandEqual("admin"))

		bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(message.Chat.ID),
				"Вы открыли режим администратора123",
			).WithReplyMarkup(kb.InlineKeyboardAdmin),
			)
		}, th.CallbackDataEqual("callback_hello1"))

		if update.Message.Text == "/hello" {

			_, _ = bot.SendMessage(tu.Message(
				tu.ID(update.Message.GetChat().ID),
				"Напшите какой вы хотите приветствие:",
			))

			Hellostring = update.Message.Text

			_, _ = bot.SendMessage(tu.Message(
				tu.ID(update.Message.GetChat().ID),
				"Приветствие изменено на:"+Hellostring,
			))

		}

		bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) { //backHandler

			_, _ = bot.SendMessage(tu.Message(
				tu.ID(query.Message.GetChat().ID),
				"Напшите какой вы хотите приветствие:",
			))

			if update.Message != nil {
				Hellostring = update.Message.Text
			}

			_, _ = bot.SendMessage(tu.Message(
				tu.ID(query.Message.GetChat().ID),
				"Приветствие изменено на:"+Hellostring,
			))

		}, th.CallbackDataEqual("admin"))

		bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) { //  Handler 1 main button
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(query.Message.GetChat().ID),
				"Наши кроссовки:",
			).WithReplyMarkup(kb.InlineKeyboard2))
		}, th.CallbackDataEqual("callback_1"))

		bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) { //  Handler 2 main button
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(query.Message.GetChat().ID),
				"Ваша корзина:",
			).WithReplyMarkup(kb.InlineKeyboardCart))

			res1, res2 := DAL.CartVision(bot, query)
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(query.Message.GetChat().ID),
				res1+res2,
			))
		}, th.CallbackDataEqual("callback_2"))

		bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) { //backHandler
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(query.Message.GetChat().ID),
				"Главное меню:",
			).WithReplyMarkup(kb.InlineKeyboard))
		}, th.CallbackDataEqual("callback_0"))

		bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) { //backHandler
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(query.Message.GetChat().ID),
				"Наши кроссовки:",
			).WithReplyMarkup(kb.InlineKeyboard2))
		}, th.CallbackDataEqual("callback_back"))

		bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) { //backHandler
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(query.Message.GetChat().ID),
				"Добавить в корзину:",
			).WithReplyMarkup(kb.InlineKeyboardPuma))

			photo := tu.Photo(
				// Chat ID as String (target username)
				tu.ID(update.Message.Chat.ID),

				// Send using file from disk
				tu.File(mustOpen("img/puma.jpg")),
			).WithCaption("Puma")

			// Sending photo
			_, _ = bot.SendPhoto(photo)
		}, th.CallbackDataEqual("callback_puma"))

		bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) { //backHandler
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(query.Message.GetChat().ID),
				"Добавить в корзину:",
			).WithReplyMarkup(kb.InlineKeyboardAdidas))

			photo := tu.Photo(
				// Chat ID as String (target username)
				tu.ID(update.Message.Chat.ID),

				// Send using file from disk
				tu.File(mustOpen("img/adidas.jpg")),
			).WithCaption("Adidas")

			_, _ = bot.SendPhoto(photo)
		}, th.CallbackDataEqual("callback_adidas"))

		bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) { //backHandler
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(query.Message.GetChat().ID),
				"Добавить в корзину:",
			).WithReplyMarkup(kb.InlineKeyboardNike))
			photo := tu.Photo(
				// Chat ID as String (target username)
				tu.ID(update.Message.Chat.ID),

				// Send using file from disk
				tu.File(mustOpen("img/nike.jpg")),
			).WithCaption("Nike")

			// Sending photo
			_, _ = bot.SendPhoto(photo)

		}, th.CallbackDataEqual("callback_nike"))

		bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {

			DAL.BrandAddInCart(bot, query, 43, "puma")

			_, _ = bot.SendMessage(tu.Message(
				tu.ID(query.Message.GetChat().ID),
				"Вы успешно добавили в корзину",
			))
		}, th.CallbackDataEqual("callback_puma43size"))

		bh.Handle(func(bot *telego.Bot, update telego.Update) {

			_, _ = bot.SendMessage(tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Я не понимаю текстовых сообщений:(",
			).WithReplyMarkup(kb.InlineKeyboard),
			)

		}, th.AnyMessage())

		bh.Handle(func(bot *telego.Bot, update telego.Update) {

			_, _ = bot.SendMessage(tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Напшите какой вы хотите приветствие:",
			))

			if update.Message != nil {
				Hellostring = update.Message.Text

			}

		}, th.CallbackDataEqual("callback_hello1")) // так нельзя делать тк это для коллбек хендлера

		bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {

			DAL.BrandDeleteInCart(bot, query)
			_, _ = bot.SendMessage(tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Корзина очищена",
			).WithReplyMarkup(kb.InlineKeyboard),
			)

		}, th.CallbackDataEqual("callback_deletecart"))

		defer bh.Stop()

		bh.Start()
	}

}

func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
