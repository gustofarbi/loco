package translation

import "golang.org/x/text/language"

type Create struct {
	Format string
	Items  map[language.Tag]map[string]Item
}

type Item struct {
	Description string
	Value       string
}

type Get struct {
	Key string
}

type Delete struct {
	Keys []string
}
