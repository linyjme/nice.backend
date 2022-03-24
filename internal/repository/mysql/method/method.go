package method

import (
	"gorm.io/gen"
)

type X string

type Method interface {
	// Where("name=@name and age=@age")
	FindByNameAndAge(name string, age int) (gen.T, error)
	//sql(select id,name,age from users where age>18)
	FindBySimpleName() ([]gen.T, error)

	//sql(select id,name,age from @@table where age>18
	//{{if cond1}}and id=@id {{end}}
	//{{if name == ""}}and @@col=@name{{end}})
	FindByIDOrName(cond1 bool, id int, col, name string) (gen.T, error)

	//sql(select * from users)
	FindAll() ([]gen.M, error)

	//sql(select * from users limit 1)
	FindOne() gen.M

	//sql(select address from users limit 1)
	FindAddress() (gen.T, error)
}

type Methods interface {
}

// only administrator to Administrator
type AdministratorMethod interface {
	//where(id=@id)
	FindByID(id int) (gen.T, error)

	// Where("account=@account and password=@password")
	SimpleFindByAccountAndPassword(account string, password string) (gen.T, error)

	//FindAdministratorToMap(id int) (gen.M, error)
	//
	//FindByAccount(account string)
	//
	//UpdateAdministratorName(name string, id int) error
}
