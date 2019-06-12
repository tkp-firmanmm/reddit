package main

import (
	"fmt"
	"log"

	"github.com/tkp-firmanmm/reddit"
)

func main() {
	items, err := reddit.Get("golang")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}

}
