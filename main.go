package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"github.com/i0rek/importtest/something"
)

func main() {
	fmt.Println(something.TEST)
	spew.Dump(something.TEST)
}
