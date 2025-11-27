# gorse-go

[![Go Reference](https://pkg.go.dev/badge/github.com/gorse-io/gorse-go.svg)](https://pkg.go.dev/github.com/gorse-io/gorse-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/gorse-io/gorse-go)](https://goreportcard.com/report/github.com/gorse-io/gorse-go)
[![Go Version](https://img.shields.io/github/go-mod/go-version/gorse-io/gorse-go?logo=go)](https://github.com/gorse-io/gorse-go/blob/main/go.mod)

Go SDK for Gorse recommender system.

> ⚠️⚠️⚠️ This SDK is unstable currently. APIs might be changed in later versions.

## Install

```bash
go get github.com/gorse-io/gorse-go
```

## Usage

```go
package main

import (
    "context"
    "log"
    "time"

    client "github.com/gorse-io/gorse-go"
)

func main() {
    // Create client
    gorse := client.NewGorseClient("http://127.0.0.1:8088", "api_key")
    ctx := context.Background()

    // Insert a user
    if _, err := gorse.InsertUser(ctx, client.User{
        UserId:  "bob",
        Labels:  map[string]any{"age": 30, "gender": "M"},
        Comment: "new user",
    }); err != nil {
        log.Fatal(err)
    }

    // Insert an item
    if _, err := gorse.InsertItem(ctx, client.Item{
        ItemId:     "movie_1",
        IsHidden:   false,
        Labels:     map[string]any{"embedding": []any{0.1, 0.2, 0.3}},
        Categories: []string{"Comedy"},
        Timestamp:  time.Now().UTC().Truncate(time.Second),
        Comment:    "Example Movie (2024)",
    }); err != nil {
        log.Fatal(err)
    }

    // Insert feedback (timestamps are time.Time)
    if _, err := gorse.InsertFeedback(ctx, []client.Feedback{
        {FeedbackType: "watch", UserId: "bob", ItemId: "movie_1", Value: 1.0, Timestamp: time.Now().UTC().Truncate(time.Second)},
    }); err != nil {
        log.Fatal(err)
    }

    // Get recommendations
    recs, err := gorse.GetRecommend(ctx, "bob", "", 10, 0)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("recommendations: %v", recs)
}
```
