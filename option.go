package ako

import "fmt"

type Option struct {
	// 选项默认值
	value string
	// 选项描述
	description string
}

func NewOption(value, description string) *Option {
	return &Option{
		value: value,
		description: description,
	}
}

func (o *Option) Description() string {
	if o.value != "" {
		return fmt.Sprintf("%s, default %s", o.description, o.value)
	}
	return o.description
}
