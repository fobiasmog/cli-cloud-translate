package translator

import (
	"context"
	"fmt"
	"html"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

type Translator struct {
	client *translate.Client
}

func New(apiKey string) (*Translator, error) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create translate client: %w", err)
	}
	return &Translator{client: client}, nil
}

func (t *Translator) Close() {
	if t.client != nil {
		t.client.Close()
	}
}

func (t *Translator) Translate(ctx context.Context, text, sourceLang, targetLang string) (string, error) {
	target, err := language.Parse(targetLang)
	if err != nil {
		return "", fmt.Errorf("invalid target language %q: %w", targetLang, err)
	}

	opts := &translate.Options{}
	if sourceLang != "" {
		src, err := language.Parse(sourceLang)
		if err != nil {
			return "", fmt.Errorf("invalid source language %q: %w", sourceLang, err)
		}
		opts.Source = src
	}

	translations, err := t.client.Translate(ctx, []string{text}, target, opts)
	if err != nil {
		return "", fmt.Errorf("translation failed: %w", err)
	}

	if len(translations) == 0 {
		return "", fmt.Errorf("no translation returned")
	}

	return html.UnescapeString(translations[0].Text), nil
}

func (t *Translator) DetectLanguage(ctx context.Context, text string) (string, error) {
	detections, err := t.client.DetectLanguage(ctx, []string{text})
	if err != nil {
		return "", fmt.Errorf("language detection failed: %w", err)
	}

	if len(detections) == 0 || len(detections[0]) == 0 {
		return "", fmt.Errorf("no language detected")
	}

	return detections[0][0].Language.String(), nil
}

func ParsePair(pair string) (source, target string, err error) {
	parts := strings.SplitN(pair, "-", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("invalid language pair %q: expected format like \"en-ru\"", pair)
	}
	return parts[0], parts[1], nil
}
