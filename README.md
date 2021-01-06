# Ako

### 介绍
Ako是用Golang实现的命令行工具，可以很方便的实现Golang的命令行应用程序。

### 概念

```
Usage: ako COMMAND <ARGUMENT> [--OPTION...]
```

- Ako由命令(COMMAND)、参数(ARGUMENT)、选项(OPTION)组成
- 每一个`命令`可以预设多个可选的`参数`,当不传入`参数`时也不影响程序执行
- 每个`参数`可以预设多个可选的`选项`,同时`选项`也可以是未预设的,会被正常解析,只是不会出现在`帮助`中
- 每个`选项`是可以设置默认值, 不同于`flag`包, `选项`值都是字符串类型, 所以`空字符串("")`代表没有默认值

> 比如演示案例中有个`server`命令,该命令有`start`、`stop`、 `reload` 三个参数, `start`参数有`listen`、`file`、`log`三个选项,其中`log`选项是未预设的

### 使用演示

```golang
package main

import (
	"fmt"
	"github.com/wuquanyao/ako"
)

// go run example/main.go server start --listen=:9000 --file=./app.yml --log=./aoo.log
func main()  {
	c := ako.NewConsole("ako")

	c.Wrap(StartCommand)
	c.Wrap(VersionCommand)

	c.Run()
	// or
	// c.Args(os.Args[1:])
}

func StartCommand(c *ako.Console) {
	// 注册`server`命令, 拥有`start`、`stop`、`reload`三个参数
	cmd := c.AddCommand("server", "http server", func(v ako.Value) {
		switch v.Argument {
		case "start":
			if v.Options["listen"] != "" {
				// todo...
				fmt.Println("listening: ", v.Options["listen"])
			}

			if v.Options["file"] != "" {
				// todo...
				fmt.Println("config file: ", v.Options["file"])
			}

            // 未预设的选项
			if v.Options["log"] != "" {
				// todo...
				fmt.Println("log file: ", v.Options["log"])
			}

			// todo...
		case "stop":
			if v.Options["grace"] == "yes" {
				// todo...
			}

			// todo...
		case "reload":
			// todo...
		}
		fmt.Println(fmt.Sprintf("cmd: server, argument: %s, options: %d", v.Argument, len(v.Options)))
	})

	// 参数`start`拥有`listen`和`file`两个选项
	cmd.AddArgument("start", "start server").
		AddOption("listen", ":9000", "listen address [HOST:PORT]").
		AddOption("file", "./config/app.yml","configuration file")

	// 参数`stop`只拥有`grace`一个选项
	cmd.AddArgument("stop", "stop server").
		AddOption("grace", "no", "gracefully terminate server host, yes or no")

	// 参数`reload`没有选项
	cmd.AddArgument("reload", "reload config")
}

func VersionCommand(c *ako.Console) {
	// 注册`version`命令, 没有参数和参数选项
	c.AddCommand("version", "show app version information", func(v ako.Value) {
		fmt.Println(fmt.Sprintf("version: 1.0.0"))
	})
}
```

### 运行脚本

```shell
go run example/main.go server start --listen=:9000 --file=./app.yml --log=./aoo.log
```

### 输出效果

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
