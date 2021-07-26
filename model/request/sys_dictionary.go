package request

import "asyncClient/model"

type SysDictionarySearch struct {
	model.SysDictionary
	PageInfo
}
