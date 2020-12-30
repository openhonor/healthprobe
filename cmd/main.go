package main

import (
	"code.byted.org/gopkg/logs"
	"testesr/pkg"
	"time"
)

func main() {
	defer logs.Flush()
	task := pkg.Task{
		Id:   0,
		Name: "task1",
		Nodes: []pkg.Node{
			{
				Name: "cluster1",
				IP:   "www.baidu.com",
			},
		},
		TaskConfig: pkg.TaskConfig{
			Interval:  1,
			Threshold: 1,
		},
		Enable: true,
	}
	stop := make(chan interface{})
	probe := pkg.Probe{task}
	go probe.Run(stop)

	time.Sleep(10 * time.Second)
	close(stop)
}
