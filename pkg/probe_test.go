package pkg

import (
	"code.byted.org/gopkg/logs"
	"testing"
	"time"
)

func TestProbe_Equal(t *testing.T) {
	p1 := Probe{
		PTask: &Task{
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
		ProbeFunc: func(id int, name, nodename, nodeip string) (success bool) {
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
		PTask: &Task{
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
		ProbeFunc: func(id int, name, nodename, nodeip string) (success bool) {
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
func TestProbe_ManageStart(t *testing.T) {
	p1 := &Probe{
		PTask: &Task{
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
		ProbeFunc: func(id int, name, nodename, nodeip string) (success bool) {
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

	p2 := &Probe{
		PTask: &Task{
			Id:   2,
			Name: "task2",
			Nodes: []Node{
				{
					Name: "cluster2",
					IP:   "www.baidu.com222",
				},
			},
			TaskConfig: TaskConfig{
				Interval:  1,
				Threshold: 1,
			},
			Enable: true,
		},
		ProbeFunc: func(id int, name, nodename, nodeip string) (success bool) {
			logs.Info("probe22222 ...")
			return true
		},
		OnFaillure: func() error {
			logs.Error("send alarm22222 ")
			return nil
		},
		OnSuccess: func() error {
			logs.Info("send success message2222")
			return nil
		},
	}

	m := NewProbeMange(time.Second)
	exit := make(chan interface{})

	m.AddTask(p1)
	m.AddTask(p2)

	time.Sleep(10 * time.Second)
	m.RemoveTask(p1)
	m.Start(exit)

}

func TestWrite_jsonConfig(t *testing.T) {
	var res []Task
	PTask := Task{
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
	}
	PTask2 := Task{
		Id:   2,
		Name: "task2",
		Nodes: []Node{
			{
				Name: "cluster2",
				IP:   "www.google.com",
			},
		},
		TaskConfig: TaskConfig{
			Interval:  2,
			Threshold: 2,
		},
		Enable: false,
	}
	res = append(res, PTask)
	res = append(res, PTask2)
	Write(res)
}
func TestRead_jsonConfig(t *testing.T) {
	tasks := Read()
	t.Log("tasks ", tasks)
}
func Test_RunTask(t *testing.T) {

	m := NewProbeMange(time.Second*3)
	//exit := make(chan interface{})

	task1 := &Probe{
		PTask: newTask(1, "task1", "cluster111"),
		ProbeFunc: func(id int, name, nodename, nodeip string) (success bool) {
			time.Sleep(2 * time.Second)
			logs.Infof("probe111 ...%s-%s", nodename, nodeip)
			return true
		},
		OnFaillure: func() error {
			logs.Error("send alarm111 ")
			return nil
		},
		OnSuccess: func() error {
			logs.Infof("send success message111")
			return nil
		},
	}
	task2 := &Probe{
		PTask: newTask(2, "task2", "cluster222"),
		ProbeFunc: func(id int, name, nodename, nodeip string) (success bool) {
			time.Sleep(20 * time.Second)
			logs.Infof("probe222 ...%s-%s", nodename, nodeip)
			return true
		},
		OnFaillure: func() error {
			logs.Error("send alarm222 ")
			return nil
		},
		OnSuccess: func() error {
			logs.Info("send success message222")
			return nil
		},
	}
	//m.AddTask(task1)
	m.AddTask(task2)

	time.Sleep(10 * time.Second)
	m.RemoveTask(task1)
	time.Sleep(30 * time.Second)
	//m.Start(exit)
}
func newTask(id int, taskName, clustername string) *Task {
	p := Task{
		Id:   id,
		Name: taskName,
		Nodes: []Node{
			{
				Name: clustername,
				IP:   "www.google.com",
			},
		},
		TaskConfig: TaskConfig{
			Interval:  4,
			Threshold: 2,
		},
		Enable: false,
	}
	return &p
}
