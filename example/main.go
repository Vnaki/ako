package main

import (
	"fmt"
	"github.com/wuquanyao/ako"
)

// go run example/main.go server start --listen=:9000 --file=./app.yml
func main()  {
	c := ako.NewConsole("ako")

	c.Wrap(StartCommand)
	c.Wrap(StopCommand)
	c.Wrap(ReloadCommand)

	c.Run()
	// or
	// c.Args(os.Args[1:])
}

func StartCommand(c *ako.Console) {
	cmd := c.AddCommand("server", "http server", func(v ako.Value) {
		fmt.Println(fmt.Sprintf("cmd: server, argument: %s, options: %d", v.Argument, len(v.Options)))
	})

	cmd.AddArgument("start", "start server").
		AddOption("listen", "listen address [HOST:PORT]").
		AddOption("file", "configuration file")

	cmd.AddArgument("stop", "stop server")
}

func StopCommand(c *ako.Console) {
	c.AddCommand("stop", "stop http server", func(v ako.Value) {
		fmt.Println(fmt.Sprintf("cmd: stop, argument: %s, options: %d", v.Argument, len(v.Options)))
	})
}

func ReloadCommand(c *ako.Console) {
	c.AddCommand("reload", "reload http server", func(v ako.Value) {
		fmt.Println(fmt.Sprintf("cmd: reload, argument: %s, options: %d", v.Argument, len(v.Options)))
	})
}