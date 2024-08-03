package DAL

import (
	"fmt"
	"os"

	mod "t/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mymmrac/telego"

	tu "github.com/mymmrac/telego/telegoutil"
)

func Migration() *gorm.DB {
	dsn := "host=localhost user=root dbname=database password=root sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&mod.User{}, &mod.Cart{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return db
}

var db *gorm.DB = Migration()

func IdAdd(bot *telego.Bot, update telego.Update) {

	userID := tu.ID(update.Message.From.ID) // Взятие ID пользователя
	userIDstring := userID.String()         //ID в String
	res := db.Raw("SELECT * FROM users WHERE lastname = ?", userIDstring).Scan(&mod.User{}).RowsAffected
	if res == 0 {
		user := mod.User{Lastname: userIDstring} // Добавление ID в БД
		result := db.Create(&user)               // создает новую запись в базе данных
		if result.Error != nil {

		}
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Добавил в базу",
		),
		)
		return

	}
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(update.Message.Chat.ID),
		"Ты уже в базе",
	),
	)
}

func BrandAddInCart(bot *telego.Bot, query telego.CallbackQuery) {

	userID := tu.ID(query.From.ID)                        // Взятие ID пользователя
	userIDstring := userID.String()                       //ID в String
	user := mod.Cart{ID: userIDstring, Namebrand: "Puma"} // Добавление ID в БД
	result := db.Create(&user)                            // создает новую запись в базе данных
	if result.Error != nil {
		// обработкcа ошибки создания
	}

}
