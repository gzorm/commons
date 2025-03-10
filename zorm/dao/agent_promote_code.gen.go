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

func newAgentPromoteCode(db *gorm.DB, opts ...gen.DOOption) agentPromoteCode {
	_agentPromoteCode := agentPromoteCode{}

	_agentPromoteCode.agentPromoteCodeDo.UseDB(db, opts...)
	_agentPromoteCode.agentPromoteCodeDo.UseModel(&model.AgentPromoteCode{})

	tableName := _agentPromoteCode.agentPromoteCodeDo.TableName()
	_agentPromoteCode.ALL = field.NewAsterisk(tableName)
	_agentPromoteCode.ID = field.NewInt64(tableName, "id")
	_agentPromoteCode.AgentID = field.NewInt64(tableName, "agent_id")
	_agentPromoteCode.PromoteCodeImg = field.NewString(tableName, "promote_code_img")
	_agentPromoteCode.PromoteCode = field.NewString(tableName, "promote_code")
	_agentPromoteCode.CreatedAt = field.NewInt64(tableName, "created_at")
	_agentPromoteCode.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_agentPromoteCode.fillFieldMap()

	return _agentPromoteCode
}

// agentPromoteCode 代理邀請二維碼
type agentPromoteCode struct {
	agentPromoteCodeDo

	ALL            field.Asterisk
	ID             field.Int64  // id
	AgentID        field.Int64  // 代理uid
	PromoteCodeImg field.String // 代理上傳的二維碼圖片網址
	PromoteCode    field.String // 代理上傳的邀請碼
	CreatedAt      field.Int64  // 創建時間
	UpdatedAt      field.Int64  // 更新時間

	fieldMap map[string]field.Expr
}

func (a agentPromoteCode) Table(newTableName string) *agentPromoteCode {
	a.agentPromoteCodeDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a agentPromoteCode) As(alias string) *agentPromoteCode {
	a.agentPromoteCodeDo.DO = *(a.agentPromoteCodeDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *agentPromoteCode) updateTableName(table string) *agentPromoteCode {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.AgentID = field.NewInt64(table, "agent_id")
	a.PromoteCodeImg = field.NewString(table, "promote_code_img")
	a.PromoteCode = field.NewString(table, "promote_code")
	a.CreatedAt = field.NewInt64(table, "created_at")
	a.UpdatedAt = field.NewInt64(table, "updated_at")

	a.fillFieldMap()

	return a
}

func (a *agentPromoteCode) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *agentPromoteCode) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["id"] = a.ID
	a.fieldMap["agent_id"] = a.AgentID
	a.fieldMap["promote_code_img"] = a.PromoteCodeImg
	a.fieldMap["promote_code"] = a.PromoteCode
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
}

func (a agentPromoteCode) clone(db *gorm.DB) agentPromoteCode {
	a.agentPromoteCodeDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a agentPromoteCode) replaceDB(db *gorm.DB) agentPromoteCode {
	a.agentPromoteCodeDo.ReplaceDB(db)
	return a
}

type agentPromoteCodeDo struct{ gen.DO }

type IAgentPromoteCodeDo interface {
	gen.SubQuery
	Debug() IAgentPromoteCodeDo
	WithContext(ctx context.Context) IAgentPromoteCodeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAgentPromoteCodeDo
	WriteDB() IAgentPromoteCodeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAgentPromoteCodeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAgentPromoteCodeDo
	Not(conds ...gen.Condition) IAgentPromoteCodeDo
	Or(conds ...gen.Condition) IAgentPromoteCodeDo
	Select(conds ...field.Expr) IAgentPromoteCodeDo
	Where(conds ...gen.Condition) IAgentPromoteCodeDo
	Order(conds ...field.Expr) IAgentPromoteCodeDo
	Distinct(cols ...field.Expr) IAgentPromoteCodeDo
	Omit(cols ...field.Expr) IAgentPromoteCodeDo
	Join(table schema.Tabler, on ...field.Expr) IAgentPromoteCodeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAgentPromoteCodeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAgentPromoteCodeDo
	Group(cols ...field.Expr) IAgentPromoteCodeDo
	Having(conds ...gen.Condition) IAgentPromoteCodeDo
	Limit(limit int) IAgentPromoteCodeDo
	Offset(offset int) IAgentPromoteCodeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAgentPromoteCodeDo
	Unscoped() IAgentPromoteCodeDo
	Create(values ...*model.AgentPromoteCode) error
	CreateInBatches(values []*model.AgentPromoteCode, batchSize int) error
	Save(values ...*model.AgentPromoteCode) error
	First() (*model.AgentPromoteCode, error)
	Take() (*model.AgentPromoteCode, error)
	Last() (*model.AgentPromoteCode, error)
	Find() ([]*model.AgentPromoteCode, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AgentPromoteCode, err error)
	FindInBatches(result *[]*model.AgentPromoteCode, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AgentPromoteCode) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAgentPromoteCodeDo
	Assign(attrs ...field.AssignExpr) IAgentPromoteCodeDo
	Joins(fields ...field.RelationField) IAgentPromoteCodeDo
	Preload(fields ...field.RelationField) IAgentPromoteCodeDo
	FirstOrInit() (*model.AgentPromoteCode, error)
	FirstOrCreate() (*model.AgentPromoteCode, error)
	FindByPage(offset int, limit int) (result []*model.AgentPromoteCode, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAgentPromoteCodeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a agentPromoteCodeDo) Debug() IAgentPromoteCodeDo {
	return a.withDO(a.DO.Debug())
}

func (a agentPromoteCodeDo) WithContext(ctx context.Context) IAgentPromoteCodeDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a agentPromoteCodeDo) ReadDB() IAgentPromoteCodeDo {
	return a.Clauses(dbresolver.Read)
}

func (a agentPromoteCodeDo) WriteDB() IAgentPromoteCodeDo {
	return a.Clauses(dbresolver.Write)
}

func (a agentPromoteCodeDo) Session(config *gorm.Session) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Session(config))
}

func (a agentPromoteCodeDo) Clauses(conds ...clause.Expression) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a agentPromoteCodeDo) Returning(value interface{}, columns ...string) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a agentPromoteCodeDo) Not(conds ...gen.Condition) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a agentPromoteCodeDo) Or(conds ...gen.Condition) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a agentPromoteCodeDo) Select(conds ...field.Expr) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a agentPromoteCodeDo) Where(conds ...gen.Condition) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a agentPromoteCodeDo) Order(conds ...field.Expr) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a agentPromoteCodeDo) Distinct(cols ...field.Expr) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a agentPromoteCodeDo) Omit(cols ...field.Expr) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a agentPromoteCodeDo) Join(table schema.Tabler, on ...field.Expr) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a agentPromoteCodeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAgentPromoteCodeDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a agentPromoteCodeDo) RightJoin(table schema.Tabler, on ...field.Expr) IAgentPromoteCodeDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a agentPromoteCodeDo) Group(cols ...field.Expr) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a agentPromoteCodeDo) Having(conds ...gen.Condition) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a agentPromoteCodeDo) Limit(limit int) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a agentPromoteCodeDo) Offset(offset int) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a agentPromoteCodeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a agentPromoteCodeDo) Unscoped() IAgentPromoteCodeDo {
	return a.withDO(a.DO.Unscoped())
}

func (a agentPromoteCodeDo) Create(values ...*model.AgentPromoteCode) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a agentPromoteCodeDo) CreateInBatches(values []*model.AgentPromoteCode, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a agentPromoteCodeDo) Save(values ...*model.AgentPromoteCode) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a agentPromoteCodeDo) First() (*model.AgentPromoteCode, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentPromoteCode), nil
	}
}

func (a agentPromoteCodeDo) Take() (*model.AgentPromoteCode, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentPromoteCode), nil
	}
}

func (a agentPromoteCodeDo) Last() (*model.AgentPromoteCode, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentPromoteCode), nil
	}
}

func (a agentPromoteCodeDo) Find() ([]*model.AgentPromoteCode, error) {
	result, err := a.DO.Find()
	return result.([]*model.AgentPromoteCode), err
}

func (a agentPromoteCodeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AgentPromoteCode, err error) {
	buf := make([]*model.AgentPromoteCode, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a agentPromoteCodeDo) FindInBatches(result *[]*model.AgentPromoteCode, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a agentPromoteCodeDo) Attrs(attrs ...field.AssignExpr) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a agentPromoteCodeDo) Assign(attrs ...field.AssignExpr) IAgentPromoteCodeDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a agentPromoteCodeDo) Joins(fields ...field.RelationField) IAgentPromoteCodeDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a agentPromoteCodeDo) Preload(fields ...field.RelationField) IAgentPromoteCodeDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a agentPromoteCodeDo) FirstOrInit() (*model.AgentPromoteCode, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentPromoteCode), nil
	}
}

func (a agentPromoteCodeDo) FirstOrCreate() (*model.AgentPromoteCode, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentPromoteCode), nil
	}
}

func (a agentPromoteCodeDo) FindByPage(offset int, limit int) (result []*model.AgentPromoteCode, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a agentPromoteCodeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a agentPromoteCodeDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a agentPromoteCodeDo) Delete(models ...*model.AgentPromoteCode) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *agentPromoteCodeDo) withDO(do gen.Dao) *agentPromoteCodeDo {
	a.DO = *do.(*gen.DO)
	return a
}
