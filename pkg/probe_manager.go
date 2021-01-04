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
	if m.IsTaskRunning(probe) {
		if !probe.Equal(m.Tasks[probe.PTask.Id]) {
			close(m.Running[probe.PTask.Id])

			// restart
			stop := make(chan interface{})
			m.Running[probe.PTask.Id] = stop
			m.Tasks[probe.PTask.Id] = probe
			go probe.Run(stop)
		}
	} else {
		stop := make(chan interface{})
		// 加入队列
		m.Running[probe.PTask.Id] = stop
		m.Tasks[probe.PTask.Id] = probe
		go probe.Run(stop)
	}
	return nil
}

func (m *ProbeManage) RemoveTask(probe *Probe) error {
	if probe == nil {
		return fmt.Errorf("can't remove a nil probe")
	}

	if m.IsTaskRunning(probe) {
		close(m.Running[probe.PTask.Id])
		delete(m.Running, probe.PTask.Id)
		delete(m.Tasks, probe.PTask.Id)
	} else {
		delete(m.Tasks, probe.PTask.Id)
	}
	return nil
}

func (m *ProbeManage) IsTaskRunning(task *Probe) bool {
	if _, ok := m.Running[task.PTask.Id]; ok {
		return true
	} else {
		return false
	}
}

// sync tasks data from source
func (m *ProbeManage) Start(exit chan interface{}) {

	ProbeFunc := func(id int, name, nodename, nodeip string) (success bool) {
		logs.Info("probe22222 ...")
		return true
	}
	OnFaillure := func() error {
		logs.Error("send alarm22222 ")
		return nil
	}
	OnSuccess := func() error {
		logs.Info("send success message2222")
		return nil
	}

	for {
		tasks := LoadConfig()
		for i := 0; i < len(tasks); i++ {
			m.AddTask(&Probe{
				PTask:       &tasks[i],
				ProbeFunc:   ProbeFunc,
				OnSuccess:   OnSuccess,
				OnFaillure:  OnFaillure,
				FailCluster: make(map[string]int),
			})
		}
		time.Sleep(m.SyncInterval)
	}

	<-exit
}
