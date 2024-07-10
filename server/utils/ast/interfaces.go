package ast

type Ast interface {
	Rollback() error
	Injection() error
}
