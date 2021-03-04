package model

type InitDBFunc interface {
	Init() (err error)
}
