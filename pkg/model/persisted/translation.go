package persisted

import (
	"github.com/samber/lo"
)

type TranslationKey struct {
	HiddenFieldsModel
	Key          string
	Translations []Translation
	Tags         []Tag `gorm:"many2many:translation_keys_tags;"`
}

type Translation struct {
	HiddenFieldsModel
	TranslationKeyID uint `json:"-" gorm:"index:idx_locale_key"`
	Description      string
	Locale           string `gorm:"index:idx_locale_key"`
	Value            string
}

type Tag struct {
	HiddenFieldsModel
	Value           string
	TranslationKeys []TranslationKey `gorm:"many2many:translation_keys_tags;"`
}

func (k TranslationKey) GetTranslationByLocale(locale string) *Translation {
	result, ok := lo.Find(k.Translations, func(item Translation) bool {
		return locale == item.Locale
	})

	if ok {
		return &result
	} else {
		return nil
	}
}
