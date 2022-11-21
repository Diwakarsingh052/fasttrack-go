package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()
	const k string = "anyKey"
	ctx = context.WithValue(ctx, k, "abc") // putting value in the context associated with a key
	FetchValue2(ctx, k)

}

func FetchValue(ctx context.Context, k string) {
	v := ctx.Value(k)
	if v == nil { // nil means value is not found // nil is the default value of the interface{}
		log.Println("value is not there ")
		return
	}
	fmt.Println(v)
}

func FetchValue2(ctx context.Context, key string) {
	v := ctx.Value(key)
	i, ok := v.(int) // type assertion // making sure data is of the same type that we want to fetch
	if !ok {
		fmt.Println("value not there or of a different type")
		return
	}
	fmt.Println(i)

}
