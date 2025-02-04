# LexoRank

LexoRank is a Go module that implements the LexoRank algorithm, a system for generating lexicographically sortable string keys. This is particularly useful for maintaining sorted lists in databases, especially when you need to insert items between existing ones without reordering the entire list.

## Features

- Generate a rank between two existing ranks
- Parse and stringify LexoRank objects
- Provides Min, Max, and Middle rank functions
- Supports multiple buckets for rank rebalancing (future)

## Installation

To install LexoRank, use `go get`:

```bash
go get github.com/tanjoshua/lexorank
```

## Usage
Here's a basic example of how to use LexoRank:

```go
package main

import (
    "fmt"
    "github.com/tanjoshua/lexorank"
)

func main() {
    // Generate ranks
    min := lexorank.Min()
    max := lexorank.Max()
    
    // Get a rank between min and max
    between, err := lexorank.Between(&min, &max)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Println("Min:", min.String())
    fmt.Println("Between:", between.String())
    fmt.Println("Max:", max.String())
}
