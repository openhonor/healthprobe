package main

import (
	"code.byted.org/gopkg/logs"
	"testesr/pkg"
	"time"
)

func main() {
	defer logs.Flush()

	m := pkg.NewProbeMange(time.Second * 3)

	exit := make(chan interface{})

	go m.Start(exit)

	<-exit
}
