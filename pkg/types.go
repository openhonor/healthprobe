package pkg

// 巡检任务
type Task struct {
	Id         int64      `json":id"`
	Name       string     `json":name"`
	Nodes      []Node     `json":nodes"`
	TaskConfig TaskConfig `json":task_config"`
	Enable     bool       `json":enable"`
}

type Node struct {
	// 节点名称
	Name string `json":name"`
	// 节点 IP
	IP string `json":ip"`
}

type TaskConfig struct {
	// 间隔时间，单位为 s
	Interval int `json":interval"`

	// 失败阈值
	Threshold int `json":threshold"`
}
