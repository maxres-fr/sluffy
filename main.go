package main

import (
	"fmt"
	"os"

	"github.com/mozillazg/go-slugify"
)

func main() {
	s := os.Args[1]
	fmt.Println(slugify.Slugify(s))
	// Output: bei-jing-kozuscek-abc
}
