package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faagerholm/lunch-bot/pkg/config"

	"github.com/PuerkitoBio/goquery"
)

func GetRestaurantMenu(res config.Restaurant) (string, error) {
	log.Println("Get menu for ", res.Url)
	resp, err := http.Get(res.Url)
	if err != nil {
		return "", err
	}

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// Save each .post-title as a list
	food := ""
	doc.Find(".meals").Each(func(i int, s *goquery.Selection) {
		if i == res.Idx {
			s.Find(".food").Each(func(_ int, f *goquery.Selection) {
				food += "- " + f.Text() + "\n"
			})
		}
	})
	if food == "" {
		food = "No food menu available"
	}
	food += fmt.Sprintf("\nFor more information, please visit their <a href=\"%s\">Website</a>.", res.Url)

	return food, nil
}
