package i18n

import (
	"embed"
	"encoding/json"
	"log/slog"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	//go:embed locales/*.json
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

func Translate(langHeader string, messageID MessageID, templateData map[string]interface{}) (string, error) {
	tags, _, err := language.ParseAcceptLanguage(langHeader)
	langs := []string{}

	if err != nil {
		slog.Error("Error parsing language header",
			"event", "i18n.parser",
			"header", langHeader,
			"error", err,
		)

		return "", err
	}

	for _, tag := range tags {
		langs = append(langs, tag.String())
	}

	localizer := i18n.NewLocalizer(bundle, langs...)
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    string(messageID),
		TemplateData: templateData,
	})

	if err != nil {
		slog.Error("Error translating message id",
			"event", "i18n.translation",
			"langs", langs,
			"message_id", messageID,
			"template_data", templateData,
			"error", err,
		)

		return "", err
	}

	slog.Info("Translate message successful",
		"event", "i18n.translation",
		"message", msg,
		"langs", langs,
		"message_id", messageID,
		"template_data", templateData,
	)

	return msg, nil
}
