// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/gzorm/commons/zorm/model"
)

func newXxlJobLog(db *gorm.DB, opts ...gen.DOOption) xxlJobLog {
	_xxlJobLog := xxlJobLog{}

	_xxlJobLog.xxlJobLogDo.UseDB(db, opts...)
	_xxlJobLog.xxlJobLogDo.UseModel(&model.XxlJobLog{})

	tableName := _xxlJobLog.xxlJobLogDo.TableName()
	_xxlJobLog.ALL = field.NewAsterisk(tableName)
	_xxlJobLog.ID = field.NewInt64(tableName, "id")
	_xxlJobLog.JobGroup = field.NewInt64(tableName, "job_group")
	_xxlJobLog.JobID = field.NewInt64(tableName, "job_id")
	_xxlJobLog.ExecutorAddress = field.NewString(tableName, "executor_address")
	_xxlJobLog.ExecutorHandler = field.NewString(tableName, "executor_handler")
	_xxlJobLog.ExecutorParam = field.NewString(tableName, "executor_param")
	_xxlJobLog.ExecutorShardingParam = field.NewString(tableName, "executor_sharding_param")
	_xxlJobLog.ExecutorFailRetryCount = field.NewInt64(tableName, "executor_fail_retry_count")
	_xxlJobLog.TriggerTime = field.NewTime(tableName, "trigger_time")
	_xxlJobLog.TriggerCode = field.NewInt64(tableName, "trigger_code")
	_xxlJobLog.TriggerMsg = field.NewString(tableName, "trigger_msg")
	_xxlJobLog.HandleTime = field.NewTime(tableName, "handle_time")
	_xxlJobLog.HandleCode = field.NewInt64(tableName, "handle_code")
	_xxlJobLog.HandleMsg = field.NewString(tableName, "handle_msg")
	_xxlJobLog.AlarmStatus = field.NewInt64(tableName, "alarm_status")

	_xxlJobLog.fillFieldMap()

	return _xxlJobLog
}

type xxlJobLog struct {
	xxlJobLogDo

	ALL                    field.Asterisk
	ID                     field.Int64
	JobGroup               field.Int64 // ID
	JobID                  field.Int64 // ID
	ExecutorAddress        field.String
	ExecutorHandler        field.String // handler
	ExecutorParam          field.String
	ExecutorShardingParam  field.String //  1/2
	ExecutorFailRetryCount field.Int64
	TriggerTime            field.Time   // -
	TriggerCode            field.Int64  // -
	TriggerMsg             field.String // -
	HandleTime             field.Time   // -
	HandleCode             field.Int64  // -
	HandleMsg              field.String // -
	AlarmStatus            field.Int64  // 0-1-2-3-

	fieldMap map[string]field.Expr
}

func (x xxlJobLog) Table(newTableName string) *xxlJobLog {
	x.xxlJobLogDo.UseTable(newTableName)
	return x.updateTableName(newTableName)
}

func (x xxlJobLog) As(alias string) *xxlJobLog {
	x.xxlJobLogDo.DO = *(x.xxlJobLogDo.As(alias).(*gen.DO))
	return x.updateTableName(alias)
}

func (x *xxlJobLog) updateTableName(table string) *xxlJobLog {
	x.ALL = field.NewAsterisk(table)
	x.ID = field.NewInt64(table, "id")
	x.JobGroup = field.NewInt64(table, "job_group")
	x.JobID = field.NewInt64(table, "job_id")
	x.ExecutorAddress = field.NewString(table, "executor_address")
	x.ExecutorHandler = field.NewString(table, "executor_handler")
	x.ExecutorParam = field.NewString(table, "executor_param")
	x.ExecutorShardingParam = field.NewString(table, "executor_sharding_param")
	x.ExecutorFailRetryCount = field.NewInt64(table, "executor_fail_retry_count")
	x.TriggerTime = field.NewTime(table, "trigger_time")
	x.TriggerCode = field.NewInt64(table, "trigger_code")
	x.TriggerMsg = field.NewString(table, "trigger_msg")
	x.HandleTime = field.NewTime(table, "handle_time")
	x.HandleCode = field.NewInt64(table, "handle_code")
	x.HandleMsg = field.NewString(table, "handle_msg")
	x.AlarmStatus = field.NewInt64(table, "alarm_status")

	x.fillFieldMap()

	return x
}

func (x *xxlJobLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := x.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (x *xxlJobLog) fillFieldMap() {
	x.fieldMap = make(map[string]field.Expr, 15)
	x.fieldMap["id"] = x.ID
	x.fieldMap["job_group"] = x.JobGroup
	x.fieldMap["job_id"] = x.JobID
	x.fieldMap["executor_address"] = x.ExecutorAddress
	x.fieldMap["executor_handler"] = x.ExecutorHandler
	x.fieldMap["executor_param"] = x.ExecutorParam
	x.fieldMap["executor_sharding_param"] = x.ExecutorShardingParam
	x.fieldMap["executor_fail_retry_count"] = x.ExecutorFailRetryCount
	x.fieldMap["trigger_time"] = x.TriggerTime
	x.fieldMap["trigger_code"] = x.TriggerCode
	x.fieldMap["trigger_msg"] = x.TriggerMsg
	x.fieldMap["handle_time"] = x.HandleTime
	x.fieldMap["handle_code"] = x.HandleCode
	x.fieldMap["handle_msg"] = x.HandleMsg
	x.fieldMap["alarm_status"] = x.AlarmStatus
}

func (x xxlJobLog) clone(db *gorm.DB) xxlJobLog {
	x.xxlJobLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return x
}

func (x xxlJobLog) replaceDB(db *gorm.DB) xxlJobLog {
	x.xxlJobLogDo.ReplaceDB(db)
	return x
}

type xxlJobLogDo struct{ gen.DO }

type IXxlJobLogDo interface {
	gen.SubQuery
	Debug() IXxlJobLogDo
	WithContext(ctx context.Context) IXxlJobLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IXxlJobLogDo
	WriteDB() IXxlJobLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IXxlJobLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IXxlJobLogDo
	Not(conds ...gen.Condition) IXxlJobLogDo
	Or(conds ...gen.Condition) IXxlJobLogDo
	Select(conds ...field.Expr) IXxlJobLogDo
	Where(conds ...gen.Condition) IXxlJobLogDo
	Order(conds ...field.Expr) IXxlJobLogDo
	Distinct(cols ...field.Expr) IXxlJobLogDo
	Omit(cols ...field.Expr) IXxlJobLogDo
	Join(table schema.Tabler, on ...field.Expr) IXxlJobLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IXxlJobLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IXxlJobLogDo
	Group(cols ...field.Expr) IXxlJobLogDo
	Having(conds ...gen.Condition) IXxlJobLogDo
	Limit(limit int) IXxlJobLogDo
	Offset(offset int) IXxlJobLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IXxlJobLogDo
	Unscoped() IXxlJobLogDo
	Create(values ...*model.XxlJobLog) error
	CreateInBatches(values []*model.XxlJobLog, batchSize int) error
	Save(values ...*model.XxlJobLog) error
	First() (*model.XxlJobLog, error)
	Take() (*model.XxlJobLog, error)
	Last() (*model.XxlJobLog, error)
	Find() ([]*model.XxlJobLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XxlJobLog, err error)
	FindInBatches(result *[]*model.XxlJobLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.XxlJobLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IXxlJobLogDo
	Assign(attrs ...field.AssignExpr) IXxlJobLogDo
	Joins(fields ...field.RelationField) IXxlJobLogDo
	Preload(fields ...field.RelationField) IXxlJobLogDo
	FirstOrInit() (*model.XxlJobLog, error)
	FirstOrCreate() (*model.XxlJobLog, error)
	FindByPage(offset int, limit int) (result []*model.XxlJobLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IXxlJobLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (x xxlJobLogDo) Debug() IXxlJobLogDo {
	return x.withDO(x.DO.Debug())
}

func (x xxlJobLogDo) WithContext(ctx context.Context) IXxlJobLogDo {
	return x.withDO(x.DO.WithContext(ctx))
}

func (x xxlJobLogDo) ReadDB() IXxlJobLogDo {
	return x.Clauses(dbresolver.Read)
}

func (x xxlJobLogDo) WriteDB() IXxlJobLogDo {
	return x.Clauses(dbresolver.Write)
}

func (x xxlJobLogDo) Session(config *gorm.Session) IXxlJobLogDo {
	return x.withDO(x.DO.Session(config))
}

func (x xxlJobLogDo) Clauses(conds ...clause.Expression) IXxlJobLogDo {
	return x.withDO(x.DO.Clauses(conds...))
}

func (x xxlJobLogDo) Returning(value interface{}, columns ...string) IXxlJobLogDo {
	return x.withDO(x.DO.Returning(value, columns...))
}

func (x xxlJobLogDo) Not(conds ...gen.Condition) IXxlJobLogDo {
	return x.withDO(x.DO.Not(conds...))
}

func (x xxlJobLogDo) Or(conds ...gen.Condition) IXxlJobLogDo {
	return x.withDO(x.DO.Or(conds...))
}

func (x xxlJobLogDo) Select(conds ...field.Expr) IXxlJobLogDo {
	return x.withDO(x.DO.Select(conds...))
}

func (x xxlJobLogDo) Where(conds ...gen.Condition) IXxlJobLogDo {
	return x.withDO(x.DO.Where(conds...))
}

func (x xxlJobLogDo) Order(conds ...field.Expr) IXxlJobLogDo {
	return x.withDO(x.DO.Order(conds...))
}

func (x xxlJobLogDo) Distinct(cols ...field.Expr) IXxlJobLogDo {
	return x.withDO(x.DO.Distinct(cols...))
}

func (x xxlJobLogDo) Omit(cols ...field.Expr) IXxlJobLogDo {
	return x.withDO(x.DO.Omit(cols...))
}

func (x xxlJobLogDo) Join(table schema.Tabler, on ...field.Expr) IXxlJobLogDo {
	return x.withDO(x.DO.Join(table, on...))
}

func (x xxlJobLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IXxlJobLogDo {
	return x.withDO(x.DO.LeftJoin(table, on...))
}

func (x xxlJobLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IXxlJobLogDo {
	return x.withDO(x.DO.RightJoin(table, on...))
}

func (x xxlJobLogDo) Group(cols ...field.Expr) IXxlJobLogDo {
	return x.withDO(x.DO.Group(cols...))
}

func (x xxlJobLogDo) Having(conds ...gen.Condition) IXxlJobLogDo {
	return x.withDO(x.DO.Having(conds...))
}

func (x xxlJobLogDo) Limit(limit int) IXxlJobLogDo {
	return x.withDO(x.DO.Limit(limit))
}

func (x xxlJobLogDo) Offset(offset int) IXxlJobLogDo {
	return x.withDO(x.DO.Offset(offset))
}

func (x xxlJobLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IXxlJobLogDo {
	return x.withDO(x.DO.Scopes(funcs...))
}

func (x xxlJobLogDo) Unscoped() IXxlJobLogDo {
	return x.withDO(x.DO.Unscoped())
}

func (x xxlJobLogDo) Create(values ...*model.XxlJobLog) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Create(values)
}

func (x xxlJobLogDo) CreateInBatches(values []*model.XxlJobLog, batchSize int) error {
	return x.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (x xxlJobLogDo) Save(values ...*model.XxlJobLog) error {
	if len(values) == 0 {
		return nil
	}
	return x.DO.Save(values)
}

func (x xxlJobLogDo) First() (*model.XxlJobLog, error) {
	if result, err := x.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLog), nil
	}
}

func (x xxlJobLogDo) Take() (*model.XxlJobLog, error) {
	if result, err := x.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLog), nil
	}
}

func (x xxlJobLogDo) Last() (*model.XxlJobLog, error) {
	if result, err := x.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLog), nil
	}
}

func (x xxlJobLogDo) Find() ([]*model.XxlJobLog, error) {
	result, err := x.DO.Find()
	return result.([]*model.XxlJobLog), err
}

func (x xxlJobLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.XxlJobLog, err error) {
	buf := make([]*model.XxlJobLog, 0, batchSize)
	err = x.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (x xxlJobLogDo) FindInBatches(result *[]*model.XxlJobLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return x.DO.FindInBatches(result, batchSize, fc)
}

func (x xxlJobLogDo) Attrs(attrs ...field.AssignExpr) IXxlJobLogDo {
	return x.withDO(x.DO.Attrs(attrs...))
}

func (x xxlJobLogDo) Assign(attrs ...field.AssignExpr) IXxlJobLogDo {
	return x.withDO(x.DO.Assign(attrs...))
}

func (x xxlJobLogDo) Joins(fields ...field.RelationField) IXxlJobLogDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Joins(_f))
	}
	return &x
}

func (x xxlJobLogDo) Preload(fields ...field.RelationField) IXxlJobLogDo {
	for _, _f := range fields {
		x = *x.withDO(x.DO.Preload(_f))
	}
	return &x
}

func (x xxlJobLogDo) FirstOrInit() (*model.XxlJobLog, error) {
	if result, err := x.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLog), nil
	}
}

func (x xxlJobLogDo) FirstOrCreate() (*model.XxlJobLog, error) {
	if result, err := x.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.XxlJobLog), nil
	}
}

func (x xxlJobLogDo) FindByPage(offset int, limit int) (result []*model.XxlJobLog, count int64, err error) {
	result, err = x.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = x.Offset(-1).Limit(-1).Count()
	return
}

func (x xxlJobLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = x.Count()
	if err != nil {
		return
	}

	err = x.Offset(offset).Limit(limit).Scan(result)
	return
}

func (x xxlJobLogDo) Scan(result interface{}) (err error) {
	return x.DO.Scan(result)
}

func (x xxlJobLogDo) Delete(models ...*model.XxlJobLog) (result gen.ResultInfo, err error) {
	return x.DO.Delete(models)
}

func (x *xxlJobLogDo) withDO(do gen.Dao) *xxlJobLogDo {
	x.DO = *do.(*gen.DO)
	return x
}
