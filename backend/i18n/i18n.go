package i18n

import (
	"embed"
	"encoding/json"
	"log/slog"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	localeFS embed.FS
	bundle   *i18n.Bundle
)

func Initi18n() error {
	bundle = i18n.NewBundle(language.Spanish)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	if _, err := bundle.LoadMessageFileFS(localeFS, "locales/es.json"); err != nil {
		slog.Error("Error embeding es lang locales",
			"event", "i18n.embed_lang",
			"lang", "es",
			"error", err,
		)

		return err
	}

	if _, err := bundle.LoadMessageFileFS(localeFS, "locales/en.json"); err != nil {
		slog.Error("Error embeding en lang locales",
			"event", "i18n.embed_lang",
			"lang", "en",
			"error", err,
		)

		return err
	}

	slog.Info("Successfuly load locale files for i18n",
		"event", "i18n.embed_locales",
		"langs", 2,
	)

	return nil
}

func Translate(lang string, messageID string, templateData map[string]interface{}) (string, error) {
	localizer := i18n.NewLocalizer(bundle, lang)

	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})

	if err != nil {
		slog.Error("Error translating message id",
			"event", "i18n.translation",
			"lang", lang,
			"message_id", messageID,
			"template_data", templateData,
			"status", 500,
			"error", err,
		)

		return "", err
	}

	slog.Info("Translate message successfullly",
		"event", "i18n.translation",
		"message", msg,
		"lang", lang,
		"message_id", messageID,
		"template_data", templateData,
	)

	return msg, nil
}
