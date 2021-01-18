package main

import (
	"internal/verses"
	"internal/commands"
	"fmt"
)

func main() {
	verses.Router();
	commands.Router();
	fmt.Printf("Hello world.");
}