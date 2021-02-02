package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/faagerholm/lunch-bot/pkg/config"
	"github.com/faagerholm/lunch-bot/pkg/web"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	key, err := config.Get(".env")

	if err != nil {
		log.Panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(key.ApiKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		// Ignore empty messages
		if update.Message == nil && update.InlineQuery == nil {
			continue
		}
		if update.Message != nil {
			log.Println("Find restaurant and display lunch alternatives: ", update.Message.Text)
			//TODO: check that restaurant requested exists
			var restaurants = config.RestaurantList()
			if val, ok := restaurants[update.Message.Text]; ok {
				var msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")
				if message, err := web.GetRestaurantMenu(val); err != nil {
					log.Fatal(err)
					return
				} else {
					msg.Text = message
					msg.ParseMode = "html"
				}
				bot.Send(msg)
			} else {
				var msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")
				msg.Text = update.Message.Text + " is not a valid restaurant"
				bot.Send(msg)
			}

			//TODO: return menu + link to homepage

			//DEBUG
			continue
		}
		inlineConf := getInlineResultsConf(update)
		if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
			log.Println(err)
		}
	}
}

func getInlineResultsConf(update tgbotapi.Update) tgbotapi.InlineConfig {
	var results []interface{}
	//TODO: move restaurants to constans

	for idx, r := range []string{"Kåren", "Gado", "Arken"} {
		if !strings.Contains(strings.ToLower(r), strings.ToLower(update.InlineQuery.Query)) {
			continue
		}
		restaurant := tgbotapi.NewInlineQueryResultArticle(fmt.Sprint(idx), r, r)
		restaurant.Description = update.InlineQuery.Query
		results = append(results, restaurant)
	}

	return tgbotapi.InlineConfig{
		InlineQueryID: update.InlineQuery.ID,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       results,
	}
}
