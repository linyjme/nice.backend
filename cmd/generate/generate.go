package main

import (
	"gorm.io/gen"
	"niceBackend/internal/repository/db_repo/categroy_repo"
	"niceBackend/internal/repository/mysql"
	"niceBackend/internal/repository/mysql/method"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/repository/mysql/query",
		// Mode:    gen.WithDefaultQuery,
	})
	g.UseDB(dal.DB)
	//administratorTpl := g.GenerateModel("tb_administrator")
	administratorTpl := g.GenerateModelAs("tb_administrator", "Administrator")
	g.ApplyBasic(categroy_repo.Category{}) // Associations
	g.ApplyBasic(administratorTpl)
	g.ApplyInterface(func(method method.AdministratorMethod) {}, administratorTpl) // struct test will be ignored
	g.Execute()
}
