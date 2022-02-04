# log-agent
日志收集项目

### 技术栈
`kafka`      --使用的是 `github.com/Shopify/sarama`
`tail`
### 踩坑总结
struct 里面的字段 **不大写就不能传入别的package**
一个小问题Windows Goland 似乎不能跑起 `main.go` ,我是在 `Windows Terminal` `go run main.go`跑起来的