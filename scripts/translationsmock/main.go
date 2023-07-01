package main

import (
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"golang.org/x/text/language"
	"loco/pkg/model/api/request/translation"
	"os"
)

var locales = []language.Tag{
	language.MustParse("de-DE"),
	language.MustParse("en-US"),
	language.MustParse("es-ES"),
	language.MustParse("fr-FR"),
	language.MustParse("it-IT"),
	language.MustParse("nl-NL"),
}

func main() {
	count := 20
	dto := translation.Create{
		Format: "json",
		Items:  make(map[language.Tag]map[string]translation.Item, count),
	}

	keys := make([]string, count)
	for i := 0; i < count; i++ {
		keys[i] = faker.Name()
	}

	for _, locale := range locales {
		dto.Items[locale] = make(map[string]translation.Item, count)
		for i := 0; i < count; i++ {
			dto.Items[locale][keys[i]] = translation.Item{
				Description: faker.Sentence(),
				Value:       faker.Sentence(),
			}
		}
	}

	f, err := os.Create("translations.json")
	if err != nil {
		panic("could not create file: " + err.Error())
	}
	defer f.Close()

	if err = json.NewEncoder(f).Encode(dto); err != nil {
		panic("could not encode json: " + err.Error())
	}
}
