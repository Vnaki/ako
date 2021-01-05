package main

import (
	"fmt"
	"github.com/wuquanyao/ako"
)

// go run example/main.go server start --listen=:9000 --file=./app.yml
func main()  {
	c := ako.NewConsole("ako")

	c.Wrap(StartCommand)
	c.Wrap(VersionCommand)

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