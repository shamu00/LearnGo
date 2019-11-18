package main

import (
	. "links"
	"os"
)

func main() {
	breathFirst(crawl, os.Args[1:])
}
