##### gin-machinery
在gin框架中使用machinery执行异步任务例子

##### 例子代码布局
```bash
.
├── gin-machinery # 编译好的例子二进制文件
├── go.mod
├── go.sum
├── main.go
├── README.md
├── server # gin http 服务
│   └── server.go
├── tasks # 任务相关
│   └── tasks.go
├── utils # 工具函数
│   └── utils.go
└── worker # 任务消费者
    └── worker.go

5 directories, 9 files

```

##### 使用
1. 启动redis
`redis-server`

2. 启动http服务
`./gin-machinery server`

3. 启动消费者
`./gin-machinery worker`

4. 调用接口
`curl --request POST "localhost:9000/send_message" --header "content-type: application/json" --data-raw '{"id": 1, "content": "hello", "target": "https://www.example.com"}'



---
that's all
`
