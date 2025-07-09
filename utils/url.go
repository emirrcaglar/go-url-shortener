package utils

import (
	"log"
)

type Url struct {
	ID       int    `json:"id"`
	LONG_URL string `json:"long_url"`
}

func GenerateShortUrl(u *Url, url string, baseUrl string) (shortUrl string) {
	err := saveToDb(u, url)

	if err != nil {
		log.Printf("Error saving to DB.")
		return
	}

	shortUrl = u.idToShortUrl(baseUrl)
	return shortUrl
}
