package main

import (
	"gorm.io/gen"
	"niceBackend/internal/repository/db_repo/categroy_repo"
	"niceBackend/internal/repository/mysql"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/repository/mysql/query",
		// Mode:    gen.WithDefaultQuery,
	})
	g.UseDB(dal.DB)

	administratorTpl := g.GenerateModelAs("tb_administrator", "Administrator")

	g.ApplyBasic(categroy_repo.Category{}) // Associations
	g.ApplyBasic(administratorTpl)
	g.Execute()
}
