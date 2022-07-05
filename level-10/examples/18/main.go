package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "language", "Go")

	foo(ctx, "language")

	ctx = context.WithValue(ctx, "dog", "Poodle")

	foo(ctx, "dog")

	foo(ctx, "bar")
}

func foo(ctx context.Context, s string) {
	// Busca um valor do context e verifica se ele existe
	if v := ctx.Value(s); v != nil {
		fmt.Println("found value:", v)
		return
	}

	fmt.Println("key not found:", s)
}
