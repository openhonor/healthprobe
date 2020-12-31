package pkg

import "code.byted.org/gopkg/logs"

func LoadConfig() {

}

type ProbeManage struct {
	Tasks map[int]*Probe

	// map, key:task id ,value stop channel
	Running map[int]chan interface{}
}

func (m *ProbeManage) AddTask(probe *Probe) {

}
func (m *ProbeManage) IsTaskRunning(task *Probe) bool {
	if _, ok := m.Running[task.PTask.Id]; ok {
		return true
	} else {
		return false
	}
}

func (m *ProbeManage) Start(exit chan interface{}) {
	stop := make(chan interface{})

	task := Task{
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
	probe := &Probe{
		PTask: task,
		ProbeFunc: func() (success bool) {
			return false
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
	go probe.Run(stop)

	<-exit
}
