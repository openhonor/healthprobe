package main

import (
	"code.byted.org/gopkg/logs"
	"testesr/pkg"
)

func main() {
	defer logs.Flush()

	manager := &pkg.ProbeManage{}
	manager.Start(make(chan interface{}))

}
