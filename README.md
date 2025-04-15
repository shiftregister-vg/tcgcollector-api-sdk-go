# TCGCollector API SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/shiftregister-vg/tcgcollector-api-sdk-go.svg)](https://pkg.go.dev/github.com/shiftregister-vg/tcgcollector-api-sdk-go)
[![Test Coverage](https://img.shields.io/badge/coverage-91.6%25-brightgreen)](https://github.com/shiftregister-vg/tcgcollector-api-sdk-go/actions/workflows/test.yml)

A Go SDK for interacting with the TCGCollector API. This SDK provides a convenient way to access all endpoints of the TCGCollector API with proper authentication, error handling, and type safety.

## Installation

```bash
go get github.com/shiftregister-vg/tcgcollector-api-sdk-go
```

## Usage

### Basic Setup

```go
import "github.com/shiftregister-vg/tcgcollector-api-sdk-go"

// Create a new client with your API key
client := tcgcollector.NewClient("your-api-key")

// Optionally configure the client with additional options
client := tcgcollector.NewClient("your-api-key",
    tcgcollector.WithBaseURL("https://api.tcgcollector.com"), // Custom base URL
    tcgcollector.WithHTTPClient(&http.Client{}),              // Custom HTTP client
)
```

### Authentication

The SDK supports both API key and OAuth2 authentication:

```go
// API Key authentication
client := tcgcollector.NewClient("your-api-key")

// OAuth2 authentication
client := tcgcollector.NewClient("your-oauth-token")
```

### Available Endpoints

The SDK provides access to all TCGCollector API endpoints:

#### Cards
- List cards: `ListCards(ctx, params)`
- Get card: `GetCard(ctx, id)`
- Get card prices: `GetCardPrices(ctx, id)`
- Recalculate cached values: `RecalculateCachedValues(ctx, id)`
- Regenerate slugs: `RegenerateSlugs(ctx, id)`
- Regenerate surrogate numbers and full names: `RegenerateSurrogateNumbersAndFullNames(ctx, id)`

#### Card Lists
- List card lists: `ListCardLists(ctx, params)`
- Get card list: `GetCardList(ctx, id)`
- List card list entries: `ListCardListEntries(ctx, id, params)`
- Recalculate card counts: `RecalculateCardCounts(ctx, id)`
- Regenerate card list slugs: `RegenerateCardListSlugs(ctx, id)`
- Bulk replace card list entries: `BulkReplaceCardListEntries(ctx, id, entries)`

#### Card Variants
- List card variants: `ListCardVariants(ctx, params)`
- Get card variant: `GetCardVariant(ctx, id)`
- Create card variant: `CreateCardVariant(ctx, variant)`
- Update card variant: `UpdateCardVariant(ctx, id, variant)`
- Delete card variant: `DeleteCardVariant(ctx, id)`
- Get card variant prices: `GetCardVariantPrices(ctx, id)`

#### Expansions
- List expansions: `ListExpansions(ctx, params)`
- Get expansion: `GetExpansion(ctx, id)`
- Recalculate expansion card counts: `RecalculateExpansionCardCounts(ctx, id)`
- Regenerate expansion slugs: `RegenerateExpansionSlugs(ctx, id)`

#### Users
- List users: `ListUsers(ctx, params)`
- Get user: `GetUser(ctx, id)`
- Create user: `CreateUser(ctx, user)`
- Update user: `UpdateUser(ctx, id, user)`
- Delete user: `DeleteUser(ctx, id)`
- Get current user: `GetCurrentUser(ctx)`
- Update current user: `UpdateCurrentUser(ctx, user)`
- Delete current user: `DeleteCurrentUser(ctx)`
- Get user count: `GetUserCount(ctx)`

#### Authentication
- Login: `Login(ctx, credentials)`
- Register: `Register(ctx, user)`
- Logout: `Logout(ctx)`
- Refresh token: `RefreshToken(ctx, token)`

#### Statistics
- Get statistics: `GetStatistics(ctx)`

#### Health
- Get health: `GetHealth(ctx)`

### Example Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/shiftregister-vg/tcgcollector-api-sdk-go"
)

func main() {
    // Create a new client
    client := tcgcollector.NewClient("your-api-key")

    // Create a context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // List cards with pagination
    cards, err := client.ListCards(ctx, &tcgcollector.ListCardsParams{
        Page:     tcgcollector.Int(1),
        PageSize: tcgcollector.Int(10),
    })
    if err != nil {
        log.Fatal(err)
    }

    // Print the cards
    for _, card := range cards.Items {
        fmt.Printf("Card: %s (ID: %d)\n", card.Name, card.ID)
    }

    // Get a specific card
    card, err := client.GetCard(ctx, 1)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Card details: %+v\n", card)
}
```

### Error Handling

The SDK returns errors in the following format:

```go
type ErrorResponse struct {
    Message string `json:"message"`
    Code    string `json:"code"`
}
```

Example error handling:

```go
card, err := client.GetCard(ctx, 1)
if err != nil {
    if apiErr, ok := err.(*tcgcollector.ErrorResponse); ok {
        fmt.Printf("API Error: %s (Code: %s)\n", apiErr.Message, apiErr.Code)
    } else {
        fmt.Printf("Other error: %v\n", err)
    }
    return
}
```

### Pagination

Many list endpoints support pagination through the `Page` and `PageSize` parameters:

```go
params := &tcgcollector.ListCardsParams{
    Page:     tcgcollector.Int(1),
    PageSize: tcgcollector.Int(50),
}
cards, err := client.ListCards(ctx, params)
```

The response includes pagination information:

```go
type ListCardsResponse struct {
    Items          []Card `json:"items"`
    ItemCount      int    `json:"itemCount"`
    TotalItemCount int    `json:"totalItemCount"`
    Page           int    `json:"page"`
    PageCount      int    `json:"pageCount"`
}
```

### Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### License

This project is licensed under the MIT License - see the LICENSE file for details.
