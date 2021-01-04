package pkg

import (
	"code.byted.org/gopkg/logs"
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"
	"time"
)

type Probe struct {
	//Stop       chan interface{}
	PTask      *Task
	ProbeFunc  ProbeFunc
	OnSuccess  func() error
	OnFaillure func() error
	lock       sync.Mutex
	// 集群失败次数 key  nodename value 为失败次数
	FailCluster map[string]int
}

func (p *Probe) SetFail(nodename string) {
	p.lock.Lock()
	p.FailCluster[nodename]++
	p.lock.Unlock()
}
func (p *Probe) ReSetFail(nodename string) {
	p.lock.Lock()
	p.FailCluster[nodename] = 0
	p.lock.Unlock()
}

const PROBE_TIMEOUT = time.Second * 3

// 探测函数
type ProbeFunc func(id int, taskName, nodeName, nodeIP string) (success bool)

func (p *Probe) DoProbe(f ProbeFunc) {
	defer logs.Flush()

	for i := 0; i < len(p.PTask.Nodes); i++ {
		go p.doProbeForSingleNode(p.PTask.Nodes[i].Name, p.PTask.Nodes[i].IP)
	}
}

// 带超时的异步调用
func (p *Probe) AsyncCall(duration time.Duration, nodeName, nodeIP string, f func() bool) (bool, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), duration)
	defer cancel()

	ch := make(chan bool, 0)

	var success bool

	go func() {
		// 探测函数

		success = f()

		ch <- success
	}()

	select {
	case res := <-ch:
		logs.Infof("%s-%s doProbeForSingleNode done ", nodeName, nodeIP)
		return res, nil
		// 超时，避免 go 泄漏
	case <-ctx.Done():
		err := fmt.Errorf("%s-%s doProbeForSingleNode timeout ", nodeName, nodeIP)
		return false, err
	}
}

func (p *Probe) doProbeForSingleNode(nodeName, nodeIP string) (bool, error) {
	var resStr string
	var success bool
	var err error
	var probeErr error
	go func() {
		// 探测函数
		success, probeErr = p.AsyncCall(PROBE_TIMEOUT, nodeName, nodeIP, func() bool {
			return p.ProbeFunc(p.PTask.Id, p.PTask.Name, nodeName, nodeIP)
		})
		if probeErr != nil {
			logs.Infof("PTask {id:%d, name: %s } Do probe,err: %s ", p.PTask.Id, p.PTask.Name, probeErr.Error())
		}

		if p.FailCluster == nil {
			p.FailCluster = make(map[string]int)
		}

		if success {
			resStr = "Success"
			err = p.OnSuccess()
			// 失败次数清空
			p.ReSetFail(nodeName)
		} else {
			resStr = "Fail"
			p.FailCluster[nodeName]++

			// 失败次数超阈值
			if p.FailCluster[nodeName] >= p.PTask.TaskConfig.Threshold {
				err = p.OnFaillure()
			}
		}
		if err != nil {
			logs.Infof("PTask {id:%d, name: %s } fail in handler function ,err: %s ", p.PTask.Id, p.PTask.Name, err.Error())
		}
		logs.Infof("PTask {id:%d, name: %s } is doing probe,result is %s !!", p.PTask.Id, p.PTask.Name, resStr)
	}()
	return success, err
}

func (p *Probe) Run(stop chan interface{}) {
	defer logs.Flush()
	for {
		select {
		case <-stop:
			logs.Infof("PTask {id:%d, name: %s } exit !!", p.PTask.Id, p.PTask.Name)
			return
		default:
			logs.Infof("PTask {id:%d, name: %s } is Running !!", p.PTask.Id, p.PTask.Name)
			time.Sleep(time.Duration(p.PTask.TaskConfig.Interval) * time.Second)

			p.DoProbe(p.ProbeFunc)
		}
	}
}

// 只比较 task 结构体
func (p *Probe) Equal(p1 *Probe) bool {
	if p1 == nil {
		return false
	}

	return reflect.DeepEqual(p.PTask, p1.PTask)
}

func (p *Probe) Read() []byte {
	fp, err := os.OpenFile("./data.json", os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	n, err := fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(data[:n]))
	return data[:n]
}

func (p *Probe) Write(data []byte) {
	fp, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
