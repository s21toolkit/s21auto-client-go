package main

import (
	"context"
	"fmt"
	"os"
	client "school-21/pkg/client"
)

func main() {
	c, err := client.NewS21Client(os.Getenv("S21_USERNAME"), os.Getenv("S21_PASSWORD"), context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(c)
}
