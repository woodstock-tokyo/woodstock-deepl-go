package main

import (
	"context"
	"fmt"
	"os"

	"github.com/michimani/deepl-sdk-go/deepl"
	"github.com/michimani/deepl-sdk-go/types"
)

func main() {
	authnKey := os.Getenv("DEEPL_AUTHN_KEY")

	client := deepl.NewClient(authnKey, true)

	usage(client)
	translateText(client)
}

func usage(c *deepl.Client) {
	params := &types.UsageParams{}
	res, errRes, err := c.Usage(context.TODO(), params)

	if err != nil {
		fmt.Println(err)
	}

	if errRes != nil {
		fmt.Println("ErrorResponse", errRes.Message)
	}

	fmt.Printf("character_count: %d\n", res.CharacterCount)
	fmt.Printf("character_limit: %d\n", res.CharacterLimit)
}

func translateText(c *deepl.Client) {
	text := []string{
		"こんにちは",
		"これはサンプルテキストです。",
	}
	params := &types.TranslateTextParams{
		TargetLang: types.TargetLangEN,
		Text:       text,
	}

	res, errRes, err := c.TranslateText(context.TODO(), params)

	if err != nil {
		fmt.Println(err)
	}

	if errRes != nil {
		fmt.Println("ErrorResponse", errRes.Message)
	}

	for i := range res.Translations {
		fmt.Printf("%s -> %s\n", text[i], res.Translations[i].Text)
	}
}
