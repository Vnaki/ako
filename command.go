package ako

import (
	"fmt"
)

type Handler = func(value Value)

type Command struct {
	// 命令名称
	name string
	// 命令处理
	handler Handler
	// 命令描述
	description string
	// 命令参数
	arguments map[string]*Argument
	// 有序的KEYS
	keys []string
}

func NewCommand(name, description string, handler Handler) *Command {
	return &Command{
		name: name,
		handler: handler,
		description: description,
		arguments: map[string]*Argument{},
		keys: []string{},
	}
}

// 添加命令参数
func (c *Command) AddArgument(name, description string) *Argument {
	if c.ExistArgument(name) == false {
		c.keys = append(c.keys, name)
	}

	c.arguments[name] = NewArgument(description)
	return c.arguments[name]
}

// 命令参数是否存在
func (c *Command) ExistArgument(name string) bool {
	if _, ok := c.arguments[name]; ok {
		return true
	}
	return false
}

// 有序遍历命令参数
func (c *Command) Loop(fn func(key string, value *Argument)) {
	for _, key := range c.keys {
		fn(key, c.arguments[key])
	}
}

// 渲染命令参数
func (c *Command) render(render RowRender) string {
	o := ""
	c.Loop(func(key string, arg *Argument) {
		o += render(key, arg.description, 2)
		
		arg.Loop(func(key string, description string) {
			o += render("--" + key, description, 4)
		})
	})
	return o
}

// 帮助
func (c *Command) help(app string) {
	NewRender().
		Title("Arguments").
		SetFormatter(c.render).
		Usage(fmt.Sprintf("Usage: %s %s <ARGUMENT> [--OPTION...]", app, c.name)).
		Render()
}

// 错误
func (c *Command) error(app, err string) {
	fmt.Printf("Report: %s.\n", err)
	c.help(app)
}
