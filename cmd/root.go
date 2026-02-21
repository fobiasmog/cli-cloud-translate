package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/q/console-translate/config"
	"github.com/q/console-translate/translator"
	"github.com/spf13/cobra"
)

var (
	version    = "0.1.0"
	pairFlag   string
	configFlag string
)

var rootCmd = &cobra.Command{
	Use:   "console-translate [text]",
	Short: "Translate text between languages",
	Long:  "A CLI tool that translates text from one language to another using Google Cloud Translation API.",
	Args:  cobra.MinimumNArgs(1),
	RunE:  runTranslate,
}

func init() {
	rootCmd.Flags().StringVarP(&pairFlag, "pair", "p", "", "language pair (e.g. en-ru)")
	rootCmd.Flags().StringVarP(&configFlag, "config", "c", "", "config file path (default: ~/.console-translate/config.json)")
	rootCmd.Version = version
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runTranslate(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load(configFlag)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	apiKey := os.Getenv("GOOGLE_TRANSLATION_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("GOOGLE_TRANSLATION_API_KEY environment variable is not set")
	}

	pair := pairFlag
	if pair == "" {
		pair = cfg.DefaultPair
	}
	if pair == "" {
		return fmt.Errorf("no language pair specified: use -p flag or set default_pair in config")
	}

	sourceLang, targetLang, err := translator.ParsePair(pair)
	if err != nil {
		return err
	}

	text := strings.Join(args, " ")

	t, err := translator.New(apiKey)
	if err != nil {
		return err
	}
	defer t.Close()

	ctx := context.Background()

	// Auto-detect: if using default pair, detect input language and swap if needed
	if pairFlag == "" && cfg.DefaultPair != "" {
		detected, err := t.DetectLanguage(ctx, text)
		if err == nil && detected == targetLang {
			sourceLang, targetLang = targetLang, sourceLang
		}
	}

	result, err := t.Translate(ctx, text, sourceLang, targetLang)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", result)
	return nil
}
