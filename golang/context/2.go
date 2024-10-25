package main

import (
	"context"
	"errors"
	"time"
)

func readDB(ctx context.Context, id uint) (string, error) {
	_ = id
	select {
	case <-ctx.Done():
		return "", errors.New("ctx cancel")
	case <-time.After(10 * time.Second):
		return "student name: ALI", nil
	}
}

func myFunc(ctx context.Context, id uint) {
	result, err := readDB(ctx, id)
	if err != nil {
		println(err.Error())
	}
	println(result)
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	myFunc(ctx, 1)
}
