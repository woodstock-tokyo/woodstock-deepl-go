deepl-sdk-go
===

This is an unofficial Go SDK for using the DeepL API.

# Usage

```bash
go get github.com/woodstock-tokyo/deepl-sdk-go
```

# Sample

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/woodstock-tokyo/deepl-sdk-go"
	"github.com/woodstock-tokyo/deepl-sdk-go/params"
	"github.com/woodstock-tokyo/deepl-sdk-go/types"
)

func main() {
	client, err := deepl.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	text := []string{
		"こんにちは",
		"これはサンプルテキストです。",
	}
	params := &params.TranslateTextParams{
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
```

```bash
$ DEEPL_API_AUTHN_KEY="your-authn-key" DEEPL_API_PLAN="free" go run main.go

こんにちは -> hello
これはサンプルテキストです。 -> This is a sample text.
```
