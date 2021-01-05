package ako

type Argument struct {
	// 参数选项
	options map[string]*Option
	// 参数描述
	description string
	// 有序的KEYS
	keys []string
}

func NewArgument(description string) *Argument {
	return &Argument{
		description: description,
		options: map[string]*Option{},
		keys: []string{},
	}
}

// 添加参数选项
func (a *Argument) AddOption(name, description string) *Argument {
	if !a.ExistOption(name) {
		a.keys = append(a.keys, name)
	}

	a.options[name] = NewOption(description)
	return a
}

// 参数选项是否存在
func (a *Argument) ExistOption(name string) bool {
	if _, ok := a.options[name]; ok {
		return true
	}
	return false
}

// 有序遍历参数选项
func (a *Argument) Loop(fn func(key string, description string)) {
	for _, key := range a.keys {
		fn(key, a.options[key].description)
	}
}

// 渲染命令参数
func (a *Argument) render(render RowRender) string {
	o := ""
	a.Loop(func(key, description string) {
		o += render(key, description, 2)
	})
	return o
}
