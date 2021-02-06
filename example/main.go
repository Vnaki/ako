package main

import (
	"fmt"

	"github.com/wuquanyao/ako"
)

// go run example/main.go server start --listen=:9000 --file=./app.yml
func main() {
	c := ako.NewConsole("ako")

	c.Wrap(ServerCommand)
	c.Wrap(VersionCommand)

	c.Run()
	// or
	// c.Args(os.Args[1:])
}

func ServerCommand(c *ako.Console) {
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
		AddOption("file", "./config/app.yml", "configuration file")

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
