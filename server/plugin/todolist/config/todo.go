package config

type TodoList struct {
	Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"` // todolist统一前缀
}
