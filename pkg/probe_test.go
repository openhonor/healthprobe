package pkg

import (
	"code.byted.org/gopkg/logs"
	"testing"
)

func TestProbe_Equal(t *testing.T) {
	p1 := Probe{
		PTask: Task{
			Id:   0,
			Name: "task1",
			Nodes: []Node{
				{
					Name: "cluster1",
					IP:   "www.baidu.com",
				},
			},
			TaskConfig: TaskConfig{
				Interval:  1,
				Threshold: 1,
			},
			Enable: true,
		},
		ProbeFunc: func() (success bool) {
			logs.Info("probe ...")
			return true
		},
		OnFaillure: func() error {
			logs.Error("send alarm ")
			return nil
		},
		OnSuccess: func() error {
			logs.Info("send success message")
			return nil
		},
	}

	p2 := Probe{
		PTask: Task{
			Id:   0,
			Name: "task1",
			Nodes: []Node{
				{
					Name: "cluster1",
					IP:   "www.baidu.com",
				},
			},
			TaskConfig: TaskConfig{
				Interval:  1,
				Threshold: 1,
			},
			Enable: true,
		},
		ProbeFunc: func() (success bool) {
			logs.Info("probe ...")
			return true
		},
		OnFaillure: func() error {
			logs.Error("send alarm ")
			return nil
		},
		OnSuccess: func() error {
			logs.Info("send success message")
			return nil
		},
	}

	res := p1.Equal(&p2)
	t.Log(res)

}
