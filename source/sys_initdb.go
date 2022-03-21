package source

import (
	"niceBackend/internal/repository/db_repo"
)

func InitDB(InitDBFunctions ...db_repo.InitDBFunc) (err error) {
	for _, v := range InitDBFunctions {
		err = v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}
