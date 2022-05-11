package main

import (
	"gopubsub/pub"
	"gopubsub/sub"
	"os"
)

func main() {
	pubOrSub := os.Args[1]

	if pubOrSub == "pub" {
		pub.Pub()
	} else if pubOrSub == "sub" {
		sub.Sub()
	}
}
