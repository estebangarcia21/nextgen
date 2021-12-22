package main

import (
	"fmt"

	"github.com/estebangarcia21/templates"
)

func main() {
	fmt.Println("Attempting generation")
	opts := templates.ReadOptions()

	if err := templates.GenerateSource(opts); err != nil {
		panic(err)
	}
}
