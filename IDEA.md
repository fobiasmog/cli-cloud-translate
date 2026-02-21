# Console Translate
The idea is to have a CLI tool that can translate text from one language to another.

## Features
- Translate text from one language to another
- has a default language pair configurable in the config file
- Auto-detect language: if default pair is set and is "en-ru" and I execute the tool without -p flag, then when I type english words, it will translate them to russian and vice versa

## Flags
- -p --pair <pair> - set language pair (en-ru, ru-en, etc.)
- -c --config <config> - set config file path: by default it is ~/.console-translate/config.json
- -h --help - show help
- -v --version - show version

## Implementation
- Use go teacup for cli
- Use go config for config
- For translation use Google Translate API
- Use go cloud.google.com/go/translate official library

## Config
- config file is a json file
- config file has a default pair field
- config file has a google api key field

## Example
$ console-translate -p en-ru hello
hello -> привет

$ console-translate -p ru-en привет
привет -> hello
