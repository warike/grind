package main

import (
	"crypto/rand"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

//go:embed quotes.json
var quotesData []byte

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

var version = "dev"

func main() {
	author := flag.String("author", "", "filter by author (case-insensitive substring)")
	short := flag.Int("short", 0, "max character length for quote")
	list := flag.Bool("list-authors", false, "list available authors")
	ver := flag.Bool("version", false, "print version")
	noColor := flag.Bool("no-color", false, "disable colored output")
	flag.Parse()

	if *ver {
		fmt.Println("grind", version)
		return
	}

	var quotes []Quote
	if err := json.Unmarshal(quotesData, &quotes); err != nil {
		fmt.Fprintf(os.Stderr, "grind: corrupt quotes data: %v\n", err)
		os.Exit(1)
	}

	if *list {
		listAuthors(quotes)
		return
	}

	filtered := filter(quotes, *author, *short)
	if len(filtered) == 0 {
		fmt.Fprintln(os.Stderr, "grind: no quotes match your filters")
		os.Exit(1)
	}

	q := pick(filtered)
	print(q, !*noColor)
}

func filter(quotes []Quote, author string, maxLen int) []Quote {
	var out []Quote
	for _, q := range quotes {
		if author != "" && !strings.Contains(strings.ToLower(q.Author), strings.ToLower(author)) {
			continue
		}
		if maxLen > 0 && len(q.Text) > maxLen {
			continue
		}
		out = append(out, q)
	}
	return out
}

func pick(quotes []Quote) Quote {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(quotes))))
	return quotes[n.Int64()]
}

func listAuthors(quotes []Quote) {
	counts := map[string]int{}
	for _, q := range quotes {
		counts[q.Author]++
	}
	for a, c := range counts {
		fmt.Printf("  %s (%d)\n", a, c)
	}
}

func print(q Quote, color bool) {
	if color {
		fmt.Printf("\033[33m\"%s\"\033[0m\n", q.Text)
		fmt.Printf("\033[2m  — %s\033[0m\n", q.Author)
	} else {
		fmt.Printf("\"%s\"\n", q.Text)
		fmt.Printf("  — %s\n", q.Author)
	}
}
