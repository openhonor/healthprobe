package pkg

import (
	"code.byted.org/gopkg/logs"
	"encoding/json"
	"os"
)

func Write(tasks []Task) {
	defer logs.Flush()

	// 创建文件
	filePtr, err := os.Create("tempdata/config.json")
	if err != nil {
		logs.Error("create file err%s", err.Error())
		return
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	encoder.Encode(tasks)
}
func Read() []Task {
	filePtr, err := os.Open("tempdata/config.json")
	if err != nil {
		logs.Error("get file err%s", err.Error())
		return nil
	}
	defer filePtr.Close()
	var info []Task
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	return info
}
