package main

import (
	"code.byted.org/gorm/bytedgen"
	"github.com/joho/godotenv"
	"gorm.io/gen"
	"os"
	"singo/dal"
)

func main() {
	cfg := gen.Config{
		OutPath:           "dal/query",
		ModelPkgPath:      "dal/sql_model",
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	}
	g := bytedgen.NewGenerator(cfg)
	// 从本地读取环境变量
	godotenv.Load()
	commonDB := dal.Database(os.Getenv("MYSQL_DSN"))
	g.UseDB(commonDB)

	//g.ApplyInterface(func(taskMethod method.TestTable) {}, g.GenerateModel("t_test_table", gen.FieldType("deleted_at", "soft_delete.DeletedAt")))
	g.ApplyBasic(g.GenerateModel("coupon"))
	g.ApplyBasic(g.GenerateModel("coupon_batch"))
	g.ApplyBasic(g.GenerateModel("coupon_issue"))
	g.ApplyBasic(g.GenerateModel("coupon_verify"))
	g.ApplyBasic(g.GenerateModel("customer"))
	g.ApplyBasic(g.GenerateModel("supplier"))
	g.ApplyBasic(g.GenerateModel("user"))
	g.ApplyBasic(g.GenerateModel("settlement"))
	g.ApplyBasic(g.GenerateModel("settlement_detail"))
	g.Execute()
}
