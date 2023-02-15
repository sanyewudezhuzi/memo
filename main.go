package main

import (
	"fmt"

	"github.com/NotAPigInTheTrefoilHouse/memo/conf"
)

func init() {
	conf.Conf()
	fmt.Println("CONTINUE")
}

func main() {
	fmt.Println("hello world")
}
