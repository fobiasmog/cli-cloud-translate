# Console Translate

**AI generated**

A CLI tool that translates text between languages using the Google Cloud Translation API.

## Setup

### 1. Get a Google Cloud API Key

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a project (or use an existing one)
3. Enable the **Cloud Translation API**
4. Create an API key under **APIs & Services > Credentials**

### 2. Set API Key

```bash
export GOOGLE_TRANSLATION_API_KEY="YOUR_API_KEY"
```

### 3. Create Config File (optional)

```bash
mkdir -p ~/.console-translate
echo '{"default_pair": "en-ru"}' > ~/.console-translate/config.json
```

### 4. Build & Install

```bash
go build -o console-translate .
# optionally move to PATH:
mv console-translate /usr/local/bin/
```

## Usage

```bash
# Translate with explicit language pair
console-translate -p en-ru hello world
# hello world -> привет мир

# Translate using default pair from config (auto-detects direction)
console-translate hello
# hello -> привет

console-translate привет
# привет -> hello

# Use a custom config file
console-translate -c /path/to/config.json hello

# Show version
console-translate --version
```

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--pair` | `-p` | Language pair (e.g. `en-ru`) |
| `--config` | `-c` | Config file path (default: `~/.console-translate/config.json`) |
| `--version` | `-v` | Show version |
| `--help` | `-h` | Show help |

## Config

The config file is JSON with two fields:

| Field | Description |
|-------|-------------|
| `default_pair` | Default language pair used when `-p` is not specified |

The API key is read from the `GOOGLE_TRANSLATION_API_KEY` environment variable.