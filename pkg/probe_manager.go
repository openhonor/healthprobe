package pkg

import (
	"code.byted.org/gopkg/logs"
	"log"
	"os"
	"time"
)

func LoadConfig() {

}

type Probe struct {
	Task Task
}

// 探测函数
type ProbeFunc func() (success bool)

func (p *Probe) DoProbe(f ProbeFunc) {
	defer logs.Flush()

	result := f()

	var resStr string
	if result {
		resStr = "Success"
	} else {
		resStr = "Fail"
	}
	logs.Infof("Task {id:%d, name: %s } is doing probe,result is %s !!", p.Task.Id, p.Task.Name, resStr)
}

func (p *Probe) Run(stop chan interface{}) {
	defer logs.Flush()
	for {
		select {
		case <-stop:
			logs.Infof("Task {id:%d, name: %s } exit !!", p.Task.Id, p.Task.Name)
			return
		default:
			logs.Infof("Task {id:%d, name: %s } is Running !!", p.Task.Id, p.Task.Name)
			time.Sleep(time.Duration(p.Task.TaskConfig.Interval) * time.Second)
			p.DoProbe(func() (success bool) {
				return false
			})
		}
	}
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
