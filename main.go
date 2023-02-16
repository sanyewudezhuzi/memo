package main

import (
	"fmt"

	"github.com/sanyewudezhuzi/memo/conf"
	"github.com/sanyewudezhuzi/memo/router"
)

func init() {
	conf.Conf()
	fmt.Println("CONTINUE")
}

func main() {
	r := router.Router()
	r.Run(conf.HttpPort)
}
