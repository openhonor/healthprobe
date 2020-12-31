package pkg

import (
	"code.byted.org/gopkg/logs"
	"log"
	"os"
	"reflect"
	"time"
)

type Probe struct {
	PTask      Task
	ProbeFunc  ProbeFunc
	OnSuccess  func() error
	OnFaillure func() error
}

// 探测函数
type ProbeFunc func() (success bool)

func (p *Probe) DoProbe(f ProbeFunc) {
	defer logs.Flush()

	var err error
	var resStr string
	if result := f(); result {
		resStr = "Success"
		err = p.OnFaillure()
	} else {
		resStr = "Fail"
		err = p.OnFaillure()
	}

	if err != nil {
		logs.Infof("PTask {id:%d, name: %s } fail in handler function ,err: %s ", p.PTask.Id, p.PTask.Name, err.Error())
	}

	logs.Infof("PTask {id:%d, name: %s } is doing probe,result is %s !!", p.PTask.Id, p.PTask.Name, resStr)
	return
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
