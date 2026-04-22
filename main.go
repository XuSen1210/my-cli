package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "World", "the name to greet")
	greeting := flag.String("greeting", "Hello", "the greeting message")
	flag.Parse()

	fmt.Printf("%s, %s!\n", *greeting, *name)
}
