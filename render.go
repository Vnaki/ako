package ako

import (
	"fmt"
	"os"
	"strings"
)

type RowRender = func(key, description string, indent int) string

type Render struct {
	// 宽度
	width int
	// 标题
	title string
	// 用法信息
	usage string
	// 底部信息
	footer string
	// 格式化程序
	formatter func(row RowRender) string
}

func NewRender() *Render {
	return &Render{
		width: 20,
		title: "Title",
		usage: "Usage: ",
		footer: "",
	}
}

func (r *Render) Title(title string) *Render {
	r.title = title
	return r
}

func (r *Render) Usage(usage string) *Render {
	r.usage = usage
	return r
}

func (r *Render) Footer(footer string) *Render {
	r.footer = footer
	return r
}

func (r *Render) SetFormatter(fn func(row RowRender) string) *Render {
	r.formatter = fn
	return r
}

func (r *Render) Render() {
	r.output(fmt.Sprintf("\n%s\n\n%s:%s\n\n%s", r.usage, r.title, r.formatter(r.render), r.footer))
}

func (r *Render) render(key, description string, indent int) string {
	return fmt.Sprintf("\n%s%s%s%s", strings.Repeat(" ", indent), key, strings.Repeat(" ", r.width - len(key) - indent), description)
}

func (r *Render) output(content string) {
	fmt.Println(content)
	os.Exit(0)
}
