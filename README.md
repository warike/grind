# grind

A stoic quotes CLI. Replaces `fortune` with 1,720 curated quotes from Seneca, Marcus Aurelius, Epictetus, and Cleanthes — embedded at compile time, zero dependencies at runtime.

## Install

```bash
brew install warike/tools/grind
```

Or build from source:

```bash
go install github.com/warike/grind@latest
```

## Usage

```bash
# Random quote
grind

# Filter by author
grind --author seneca

# Short quotes only (max 150 chars)
grind --short 150

# List available authors
grind --list-authors

# No ANSI colors
grind --no-color
```

## Shell greeting

Add to your `.zshrc` or `.bashrc`:

```bash
grind --short 300
```

## Dataset

1,720 deduplicated quotes sourced from [Stoic Quotes (Kaggle)](https://www.kaggle.com/datasets/tejasnisar/stoic-quotes), originally scraped from Goodreads.

| Author | Quotes |
|---|---|
| Seneca | 746 |
| Marcus Aurelius | 672 |
| Epictetus | 301 |
| Cleanthes | 1 |
