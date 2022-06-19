package orm

import (
	"api/enums"
	"context"
	"database/sql"
	"github.com/beego/beego/v2/client/orm"
	"goBeego/models"
	"log"
)

func Query(model string, symbol string, param ...interface{}) orm.QuerySeter {
	qs := MOrm.QueryTable(model)
	return qs.Filter(model+"__"+symbol, param)
}
func Not(qs orm.QuerySeter, model string, symbol string, param ...interface{}) orm.QuerySeter {
	return qs.Exclude(model+"__"+symbol, param)
}

func Limit(qs orm.QuerySeter, startPage int, sizePage int) orm.QuerySeter {
	return qs.Limit(sizePage, startPage)
}

func Group(qs orm.QuerySeter, param ...string) orm.QuerySeter {
	return qs.GroupBy(param...)
}

func Order(qs orm.QuerySeter, param ...string) orm.QuerySeter {
	return qs.OrderBy(param...)
}
func Count(qs orm.QuerySeter) (int64, error) {
	return qs.Count()
}

func Operation(models []interface{}, op enums.Operation, model string, queryParam map[string]string, updateParam map[string]string) {
	qs := MOrm.QueryTable(model)
	for key, val := range queryParam {
		qs.Filter(key, val)
	}

	switch op {
	case 0:
		qs.All(&models)
	case 1:
		{
			i, _ := qs.PrepareInsert()
			for model := range models {
				i.Insert(model)
			}

		}
	case 2:
		{
			for key, vla := range updateParam {
				qs.Update(orm.Params{
					key: vla,
				})
			}
		}
	case 3:
		qs.Delete()

	}
}

func QueryOne(param map[string]string, model interface{}) {
	qs := MOrm.QueryTable(model)

	for key, val := range param {
		qs.Filter(key, val)
	}
	qs.One(model)
}

func List(model interface{}, param map[string]string) []orm.ParamsList {
	var lists []orm.ParamsList

	qs := MOrm.QueryTable(model)
	for key, val := range param {
		qs.Filter(key, val)
	}
	qs.ValuesList(&lists)
	return lists
}

func ForceIndex(qs orm.QuerySeter, param ...string) orm.QuerySeter {
	return qs.ForceIndex(param...)
}

func UseIndex(qs orm.QuerySeter, param ...string) orm.QuerySeter {
	return qs.UseIndex(param...)
}

func IgnoreIndex(qs orm.QuerySeter, param ...string) orm.QuerySeter {
	return qs.IgnoreIndex(param...)
}

func Distinct(qs orm.QuerySeter) orm.QuerySeter {
	return qs.Distinct()
}

func exec(sql string, param []interface{}) (sql.Result, error) {
	return MOrm.Raw(sql, param).Exec()
}

func QueryRows(sql string, param []interface{}, models interface{}) (int64, error) {
	return MOrm.Raw(sql, param).QueryRows(models)
}

func GetQuery() orm.QueryBuilder {
	qb, _ := orm.NewQueryBuilder("mysql")
	return qb
}

func DoTx() {
	err := MOrm.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		// data
		user := new(models.User)
		user.Name = "test_transaction"

		// insert data
		// Using txOrm to execute SQL
		_, e := txOrm.Insert(user)
		// if e != nil the transaction will be rollback
		// or it will be committed
		return e
	})
	if err != nil {
		log.Fatal("err:", err)
	}
}

/**
cond := orm.NewCondition()
cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)

qs := orm.QueryTable("user")
qs = qs.SetCond(cond1)
// WHERE ... AND ... AND NOT ... OR ...

cond2 := cond.AndCond(cond1).OrCond(cond.And("name", "slene"))
qs = qs.SetCond(cond2).Count()
// WHERE (... AND ... AND NOT ... OR ...) OR ( ... )
*/

func Create(model interface{}) {
	MOrm.Insert(model)
}

func batchCreate(models []interface{}) {
	MOrm.InsertMulti(len(models), models)
}

func Delete(model interface{}) {
	MOrm.Delete(model)
}

func Update(model interface{}) {
	MOrm.Update(model)
}

func Read(model interface{}) {
	MOrm.Read(model)
}

func Sql(sql string, param ...interface{}) {
	MOrm.Raw(sql, param)
}

func highQuery(model interface{}, param map[string]interface{}) {
	queryStater := MOrm.QueryTable(model)
	for key, val := range param {
		queryStater.Filter(key, val)
	}
}
