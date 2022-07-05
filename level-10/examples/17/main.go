package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadline := time.Millisecond * 50
	ctx, cancel := context.WithTimeout(context.Background(), deadline)

	defer cancel()

	select {
	case <-time.After(time.Second * 1):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
