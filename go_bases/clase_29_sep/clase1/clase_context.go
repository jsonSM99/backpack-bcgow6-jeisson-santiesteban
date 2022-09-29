package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()

	// ctx2 := context.WithValue(ctx, "Clave", "Valor")

	// Function(ctxç, 10)

	ctx3, _ := context.WithTimeout(ctx2, 5*time.Second)
}
