# Ako

#### 介绍
Ako是用Golang实现的命令行工具，可以很方便的实现Golang的命令行应用程序。

#### 概念

```
Usage: ako COMMAND <ARGUMENT> [--OPTION...]
```

- Ako由命令(COMMAND)、参数(ARGUMENT)、选项(OPTION)组成
- 每一个`命令`至少有一个`参数`
- 每个`参数`可以有多个`选项`
- 每个`选项`是可以设置默认值, 不同于`flag`包, `选项`值都是字符串类型

比如演示案例中有个`server`命令,该命令有`start`、`stop`、 `reload` 三个参数, `start`参数有`listen`和`file`两个选项

#### 使用演示

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

#### 运行脚本

```
 go run example/main.go server start --listen=:9000 --file=./app.yml
```

#### 输出效果

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
