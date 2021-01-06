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
	// 有序的KEYS
	keys []string
}

func NewConsole(name string) *Console {
	return &Console{
		name: name,
		commands: map[string]*Command{},
		keys: []string{},
	}
}

// 注册命令
func (c *Console) AddCommand(name string, description string, handler Handler) *Command {
	if c.Command(name) == nil {
		c.keys = append(c.keys, name)
	}
	c.commands[name] = NewCommand(name, description, handler)
	return c.commands[name]
}

// 读取命令信息
func (c *Console) Command(name string) *Command {
	if cmd, ok := c.commands[name]; ok {
		return cmd
	}
	return nil
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

	if "help" == args[0] {
		if size == 1 {
			c.help()
		}

		cmd := c.Command(args[1])
		if nil == cmd {
			c.error(fmt.Sprintf("invalid command: %s", args[1]))
			// 解决语法警告
			return
		}

		cmd.help(c.name)
	}

	cmd := c.Command(args[0])
	if nil == cmd {
		c.error(fmt.Sprintf("invalid command: %s", args[0]))
		// 解决语法警告
		return
	}

	if 1 == size {
		cmd.handler(Value{})
		return
	}

	arg := cmd.Argument(args[1])
	if nil == arg {
		cmd.error(c.name, fmt.Sprintf("invalid argument: %s", args[1]))
		// 解决语法警告
		return
	}

	v := Value{
		Argument: args[1],
		Options: map[string]string{},
	}

	// 解析默认选项
	arg.Loop(func(key string, option *Option) {
		if option.value != "" {
			v.Options[key] = option.value
		}
	})

	// 解析输入选项
	if size >= 3 {
		for _, av := range args[2:] {
			if len(av) <= 2 {
				cmd.error(c.name, fmt.Sprintf("invalid option: %s, valid format: --option=value", args[2]))
			}

			if av[0] != '-' || av[1] != '-' {
				cmd.error(c.name, fmt.Sprintf("invalid option: %s, valid format: --option=value", av))
			}

			// 接收全部参数选项
			sn := strings.SplitN(av[2:], "=", 2)
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
		Footer(fmt.Sprintf("Run '%s help COMMAND' for more information on a command", c.name)).
		Render()
}

// 错误
func (c *Console) error(err string) {
	fmt.Printf("error: %s.\n", err)
	c.help()
}
