package ako

// Argument 命令参数
type Argument struct {
	// 参数选项
	options map[string]*Option
	// 参数描述
	description string
	// 有序的KEYS
	keys []string
}

// NewArgument 新建参数
func NewArgument(description string) *Argument {
	return &Argument{
		description: description,
		options: map[string]*Option{},
		keys: []string{},
	}
}

// AddOption 添加参数选项
func (a *Argument) AddOption(name, value, description string) *Argument {
	if nil == a.Option(name) {
		a.keys = append(a.keys, name)
	}
	a.options[name] = NewOption(value, description)
	return a
}

// Option 读取参数选项信息
func (a *Argument) Option(name string) *Option {
	if option, ok := a.options[name]; ok {
		return option
	}
	return nil
}

// Loop 有序遍历参数选项
func (a *Argument) Loop(fn func(key string, option *Option)) {
	for _, key := range a.keys {
		fn(key, a.options[key])
	}
}

// 渲染命令参数
func (a *Argument) render(render RowRender) string {
	o := ""
	a.Loop(func(key string, option *Option) {
		o += render(key, option.description, 2)
	})
	return o
}
