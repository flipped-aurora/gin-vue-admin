package system

type InitDBFunc interface {
	Init() (err error)
}
