# Ako

#### 介绍
Ako是用Golang实现的轻量易用的命令行工具，可以很方便的为我们Golang应用程序定义交互命令。

#### 1.使用演示

```golang
package main

import (
	"fmt"
	"github.com/wuquanyao/ako"
)

func main()  {
    c := ako.NewConsole("ako")
    
    c.Wrap(StartCommand)
    c.Wrap(VersionCommand)
    // or
    // c.AddCommand("your cmd", "description", func(v ako.Value) {
    //   todo...
    // })
    
    c.Run()
    // or
    // c.Args(os.Args[1:])
}

func StartCommand(c *ako.Console) {
	cmd := c.AddCommand("server", "http server", func(v ako.Value) {
		fmt.Println(fmt.Sprintf("cmd: server, argument: %s, options: %d", v.Argument, len(v.Options)))
	})

	cmd.AddArgument("start", "start server").
		AddOption("listen", ":9000", "listen address [HOST:PORT]").
		AddOption("file", "./config/app.yml","configuration file")

	cmd.AddArgument("stop", "stop server").
		AddOption("grace", "no", "gracefully terminate server host, yes or no")

	cmd.AddArgument("reload", "reload config")
}

func VersionCommand(c *ako.Console) {
	c.AddCommand("version", "show app version information", func(v ako.Value) {
		fmt.Println(fmt.Sprintf("version: 1.0.0"))
	})
}
```

#### 2.运行

```
 go run example/main.go server start --listen=:9000 --file=./app.yml
```

#### 3.输出效果

```
Usage: ako server <ARGUMENT> [--OPTION...]

Arguments:
  start             start server
    --listen        listen address [HOST:PORT], default :9000
    --file          configuration file, default ./config/app.yml
  stop              stop server
    --grace         gracefully terminate server host, yes or no, default no
  reload            reload config
```
