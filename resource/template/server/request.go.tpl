package request

import "asyncClient/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}