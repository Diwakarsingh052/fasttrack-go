package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background() // empty context where we will put timeouts,values in the future
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	PrintData(ctx, "hello this is a test func")
	//http.NewRequestWithContext()
	//sql.DB{}.QueryRowContext()
}

func PrintData(ctx context.Context, input string) { // ctx should be the first param in func

	select {
	case <-ctx.Done(): // it will tell when a ctx is cancelled or finished
		fmt.Println("timeout")
		fmt.Println(ctx.Err()) // to check err inside the context
	case <-time.After(1 * time.Second):
		fmt.Println(input)
	}

}
