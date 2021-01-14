## 探测程序

支持七层健康检查任务，任务可配置 开关，探测间隔，失败阈值，当前失败可配置行为

```json
[{
  "Id": 1,
  "Name": "task1",
  "Url": "https://www.google.com1",
  "HttpHeaders": {
    "A": "A"
  },
  "Nodes": [{
    "Name": "google",
    "IP": "172.217.194.102"
  }],
  "TaskConfig": {
    "Interval": 1,
    "Threshold": 10
  },
  "Enable": true
}, {
  "Id": 2,
  "Name": "task2",
  "Url": "https://stackoverflow.com/questions/47498811/how-to-send-a-map-in-json/47514487",
  "HttpHeaders": {
    "A": "A"
  },
  "Nodes": [{
    "Name": "stackoverflow",
    "IP": "151.101.65.169"
  }],
  "TaskConfig": {
    "Interval": 2,
    "Threshold": 20
  },
  "Enable": true
}]
```