package pkg

import (
	"code.byted.org/gopkg/logs"
	"fmt"
	"time"
)

func LoadConfig() []Task {
	tasks := Read()

	return tasks

}

type ProbeManage struct {
	Tasks        map[int]*Probe
	Running      map[int]chan interface{}
	SyncInterval time.Duration
}

func NewProbeMange(d time.Duration) *ProbeManage {
	return &ProbeManage{
		Tasks:        make(map[int]*Probe),
		Running:      make(map[int]chan interface{}),
		SyncInterval: d,
	}
}
func (m *ProbeManage) AddTask(probe *Probe) error {
	if probe == nil {
		return fmt.Errorf("probe is nil")
	}
	// 正在运行的话，如果配置不一样需要重启 task
	if m.IsTaskRunning(probe.PTask.Id) {
		if !probe.Equal(m.Tasks[probe.PTask.Id]) {
			close(m.Running[probe.PTask.Id])

			// restart
			stop := make(chan interface{})
			m.Running[probe.PTask.Id] = stop
			m.Tasks[probe.PTask.Id] = probe
			go probe.Run(stop)
			logs.Infof("===== reload task,id:%d,name:%s", probe.PTask.Id, probe.PTask.Name)
		}
	} else {
		stop := make(chan interface{})
		// 加入队列
		m.Running[probe.PTask.Id] = stop
		m.Tasks[probe.PTask.Id] = probe
		go probe.Run(stop)
		logs.Infof("===== add task,id:%d,name:%s", probe.PTask.Id, probe.PTask.Name)
	}

	return nil
}

func (m *ProbeManage) RemoveTask(taskId int) error {

	if m.IsTaskRunning(taskId) {
		close(m.Running[taskId])
		delete(m.Running, taskId)
		delete(m.Tasks, taskId)
		logs.Infof("===== stop and remove task,id:%d", taskId)
	} else {
		delete(m.Tasks, taskId)
	}

	return nil
}

func (m *ProbeManage) IsTaskRunning(taskId int) bool {
	if _, ok := m.Running[taskId]; ok {
		return true
	} else {
		return false
	}
}

// sync tasks data from source
func (m *ProbeManage) Start(exit chan interface{}) {

	ProbeFunc := func(id int, name, nodename, nodeip string) (success bool) {
		logs.Info("task %d-%s probe for cluster:%s-%s", id, name, nodename, nodeip)
		return true
	}
	OnFaillure := func(id int, name, nodename, nodeip string) error {
		logs.Error("task %d-%s send alarm for cluster:%s-%s", id, name, nodename, nodeip)
		return nil
	}
	OnSuccess := func(id int, name, nodename, nodeip string) error {
		logs.Info("task %d-%s send success message for cluster: cluster:%s-%s", id, name, nodename, nodeip)
		return nil
	}
	ticker := time.NewTicker(m.SyncInterval)
	for {
		select {
		case <-ticker.C:
			tasks := LoadConfig()
			logs.Info("")
			logs.Infof("<<< load tasks >>> tasks: %#v", tasks)
			logs.Info("")
			for i := 0; i < len(tasks); i++ {
				if tasks[i].Enable {
					m.AddTask(&Probe{
						PTask:       &tasks[i],
						ProbeFunc:   ProbeFunc,
						OnSuccess:   OnSuccess,
						OnFaillure:  OnFaillure,
						FailCluster: make(map[string]int),
					})
				} else {
					m.RemoveTask(tasks[i].Id)
				}

			}
		case <-exit:
			logs.Info("probe exit!!!")
			return
		}
	}

}
