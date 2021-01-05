package ako

type Option struct {
	// 选项描述
	description string
}

func NewOption(description string) *Option {
	return &Option{
		description: description,
	}
}
