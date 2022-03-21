package db_repo

type InitDBFunc interface {
	Init() (err error)
}
