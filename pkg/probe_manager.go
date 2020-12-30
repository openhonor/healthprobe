package pkg

func LoadConfig() {

}

type ProbeManage struct {
	// map, key:task id ,value stop channel
	Tasks map[int](chan interface{})
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
		Task: task,
		ProbeFunc: func() (success bool) {
			return false
		}}
	go probe.Run(stop)
	<-exit
}
