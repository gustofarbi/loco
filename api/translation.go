package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"loco/internal/di"
	"loco/pkg/model/api/request/translation"
	"loco/pkg/model/persisted"
	"net/http"
)

var database = func() *gorm.DB {
	database, err := di.GetDB()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	return database
}()

func CreateTranslationKeys(c echo.Context) error {
	dto := translation.Create{}

	if err := json.NewDecoder(c.Request().Body).Decode(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	keys := make(map[string]struct{}, len(dto.Items))

	for _, items := range dto.Items {
		for key := range items {
			keys[key] = struct{}{}
		}
	}

	uniqueKeys := lo.Uniq(lo.Keys(keys))

	var existingTranslationKeys []persisted.TranslationKey
	if err := database.
		Preload("Translations").
		Where("key IN ?", uniqueKeys).
		Find(&existingTranslationKeys).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	for _, translationKey := range existingTranslationKeys {
		save := false
		for locale, items := range dto.Items {
			localeString := locale.String()
			existingTranslation := translationKey.GetTranslationByLocale(localeString)

			value := items[translationKey.Key].Value
			description := items[translationKey.Key].Description

			if existingTranslation == nil {
				translationKey.Translations = append(
					translationKey.Translations,
					persisted.Translation{
						Locale:      localeString,
						Description: description,
						Value:       value,
					},
				)
				save = true
			} else {
				if existingTranslation.Value == value &&
					existingTranslation.Description == description {
					delete(dto.Items[locale], translationKey.Key)
					continue
				}
				existingTranslation.Value = value
				existingTranslation.Description = description
				save = true
			}

			delete(dto.Items[locale], translationKey.Key)
		}
		if save {
			database.Save(&translationKey)
		}
	}

	keysToPersist := make(map[string]*persisted.TranslationKey, 0)
	for locale, items := range dto.Items {
		for key, item := range items {
			if existing, ok := keysToPersist[key]; ok {
				existing.Translations = append(existing.Translations, persisted.Translation{
					Locale:      locale.String(),
					Value:       item.Value,
					Description: item.Description,
				})
			} else {
				keysToPersist[key] = &persisted.TranslationKey{
					Key: key,
					Translations: []persisted.Translation{
						{
							Locale:      locale.String(),
							Value:       item.Value,
							Description: item.Description,
						},
					},
				}
			}
		}
	}

	for _, key := range keysToPersist {
		if err := database.Create(&key).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, "OK")
}

func GetTranslationKey(c echo.Context) error {
	dto := translation.Get{}

	if err := json.NewDecoder(c.Request().Body).Decode(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := persisted.TranslationKey{Key: dto.Key}
	if err := database.Preload("Translations").Find(&result).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteTranslationKeys(c echo.Context) error {
	dto := translation.Delete{}

	if err := json.NewDecoder(c.Request().Body).Decode(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := database.Where("key IN ?", dto.Keys).Delete(&persisted.TranslationKey{}).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "OK")
}
