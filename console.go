package ako

import (
	"fmt"
	"os"
	"strings"
)

// app <command> [argument] [--option]
// app server start --config=./app.yml

type Console struct {
	// 控制台名称
	name string
	// 注册的命令
	commands map[string]*Command
	// 控制台描述
	description string
	// 有序的KEYS
	keys []string
}

func NewConsole(name, description string) *Console {
	return &Console{
		name: name,
		description: description,
		commands: map[string]*Command{},
		keys: []string{},
	}
}

// 注册命令
func (c *Console) AddCommand(name string, description string, handler Handler) *Command {
	if c.ExistCommand(name) == false {
		c.keys = append(c.keys, name)
	}
	c.commands[name] = NewCommand(name, description, handler)
	return c.commands[name]
}

// 判断命令是否注册
func (c *Console) ExistCommand(name string) bool {
	if _, ok := c.commands[name]; ok {
		return true
	}
	return false
}

// 解耦操作
func (c *Console) Wrap(fn func(console *Console)) {
	fn(c)
}

// 有序遍历命令
func (c *Console) Loop(fn func(key string, command *Command)) {
	for _, key := range c.keys {
		fn(key, c.commands[key])
	}
}

func (c *Console) Run() {
	c.Args(os.Args[1:])
}

func (c *Console) Args(args []string) {
	size := len(args)

	if 0 == size {
		c.help()
	}

	if !c.ExistCommand(args[0]) {
		c.error(fmt.Sprintf("invalid command: %s", args[0]))
	}

	cmd := c.commands[args[0]]

	if 1 == size {
		cmd.error(c.name, "lack argument")
	}

	if !cmd.ExistArgument(args[1]) {
		cmd.error(c.name, fmt.Sprintf("invalid argument: %s", args[1]))
	}

	v := Value{
		Argument: args[1],
		Options: map[string]string{},
	}

	if size >= 3 {
		for _, arg := range args[2:] {
			if len(arg) <= 2 {
				cmd.error(c.name, fmt.Sprintf("invalid option: %s, valid format: --option=value", args[2]))
			}

			if arg[0] != '-' || arg[1] != '-' {
				cmd.error(c.name, fmt.Sprintf("invalid option: %s, valid format: --option=value", arg))
			}

			// 接收全部参数选项
			sn := strings.SplitN(arg[2:], "=", 2)
			if len(sn) == 2 {
				v.Options[sn[0]] = sn[1]
				continue
			}
		}
	}

	cmd.handler(v)
}

// 渲染命令
func (c *Console) render(render RowRender) string {
	o := ""
	c.Loop(func(key string, cmd *Command) {
		o += render(key, cmd.description, 2)
	})
	return o
}

// 帮助
func (c *Console) help() {
	NewRender().
		Title("Commands").
		SetFormatter(c.render).
		Usage(fmt.Sprintf("Usage: %s COMMAND <ARGUMENT> [--OPTION...]", c.name)).
		Footer(fmt.Sprintf("Run `%s help COMMAND` for more information on a command", c.name)).
		Render()
}

// 错误
func (c *Console) error(err string) {
	fmt.Printf("Report: %s.\n", err)
	c.help()
}
