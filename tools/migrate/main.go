package main

import (
	"context"
	"log"

	"meh/di"
)

func main() {
	ctx := context.Background()
	c := di.InjectEntClient()

	if err := c.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
