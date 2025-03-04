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

func newWinAgentLink(db *gorm.DB, opts ...gen.DOOption) winAgentLink {
	_winAgentLink := winAgentLink{}

	_winAgentLink.winAgentLinkDo.UseDB(db, opts...)
	_winAgentLink.winAgentLinkDo.UseModel(&model.WinAgentLink{})

	tableName := _winAgentLink.winAgentLinkDo.TableName()
	_winAgentLink.ALL = field.NewAsterisk(tableName)
	_winAgentLink.ID = field.NewInt64(tableName, "id")
	_winAgentLink.Link = field.NewString(tableName, "link")
	_winAgentLink.UID = field.NewInt64(tableName, "uid")
	_winAgentLink.Status = field.NewInt64(tableName, "status")
	_winAgentLink.CreatedAt = field.NewInt64(tableName, "created_at")
	_winAgentLink.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winAgentLink.fillFieldMap()

	return _winAgentLink
}

// winAgentLink 代理专属域名
type winAgentLink struct {
	winAgentLinkDo

	ALL       field.Asterisk
	ID        field.Int64
	Link      field.String // 推广域名
	UID       field.Int64  // 代理ID
	Status    field.Int64  // 状态:1-启用 0-停用
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (w winAgentLink) Table(newTableName string) *winAgentLink {
	w.winAgentLinkDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winAgentLink) As(alias string) *winAgentLink {
	w.winAgentLinkDo.DO = *(w.winAgentLinkDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winAgentLink) updateTableName(table string) *winAgentLink {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Link = field.NewString(table, "link")
	w.UID = field.NewInt64(table, "uid")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winAgentLink) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winAgentLink) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 6)
	w.fieldMap["id"] = w.ID
	w.fieldMap["link"] = w.Link
	w.fieldMap["uid"] = w.UID
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winAgentLink) clone(db *gorm.DB) winAgentLink {
	w.winAgentLinkDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winAgentLink) replaceDB(db *gorm.DB) winAgentLink {
	w.winAgentLinkDo.ReplaceDB(db)
	return w
}

type winAgentLinkDo struct{ gen.DO }

type IWinAgentLinkDo interface {
	gen.SubQuery
	Debug() IWinAgentLinkDo
	WithContext(ctx context.Context) IWinAgentLinkDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinAgentLinkDo
	WriteDB() IWinAgentLinkDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinAgentLinkDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinAgentLinkDo
	Not(conds ...gen.Condition) IWinAgentLinkDo
	Or(conds ...gen.Condition) IWinAgentLinkDo
	Select(conds ...field.Expr) IWinAgentLinkDo
	Where(conds ...gen.Condition) IWinAgentLinkDo
	Order(conds ...field.Expr) IWinAgentLinkDo
	Distinct(cols ...field.Expr) IWinAgentLinkDo
	Omit(cols ...field.Expr) IWinAgentLinkDo
	Join(table schema.Tabler, on ...field.Expr) IWinAgentLinkDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinAgentLinkDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinAgentLinkDo
	Group(cols ...field.Expr) IWinAgentLinkDo
	Having(conds ...gen.Condition) IWinAgentLinkDo
	Limit(limit int) IWinAgentLinkDo
	Offset(offset int) IWinAgentLinkDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinAgentLinkDo
	Unscoped() IWinAgentLinkDo
	Create(values ...*model.WinAgentLink) error
	CreateInBatches(values []*model.WinAgentLink, batchSize int) error
	Save(values ...*model.WinAgentLink) error
	First() (*model.WinAgentLink, error)
	Take() (*model.WinAgentLink, error)
	Last() (*model.WinAgentLink, error)
	Find() ([]*model.WinAgentLink, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAgentLink, err error)
	FindInBatches(result *[]*model.WinAgentLink, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinAgentLink) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinAgentLinkDo
	Assign(attrs ...field.AssignExpr) IWinAgentLinkDo
	Joins(fields ...field.RelationField) IWinAgentLinkDo
	Preload(fields ...field.RelationField) IWinAgentLinkDo
	FirstOrInit() (*model.WinAgentLink, error)
	FirstOrCreate() (*model.WinAgentLink, error)
	FindByPage(offset int, limit int) (result []*model.WinAgentLink, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinAgentLinkDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winAgentLinkDo) Debug() IWinAgentLinkDo {
	return w.withDO(w.DO.Debug())
}

func (w winAgentLinkDo) WithContext(ctx context.Context) IWinAgentLinkDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winAgentLinkDo) ReadDB() IWinAgentLinkDo {
	return w.Clauses(dbresolver.Read)
}

func (w winAgentLinkDo) WriteDB() IWinAgentLinkDo {
	return w.Clauses(dbresolver.Write)
}

func (w winAgentLinkDo) Session(config *gorm.Session) IWinAgentLinkDo {
	return w.withDO(w.DO.Session(config))
}

func (w winAgentLinkDo) Clauses(conds ...clause.Expression) IWinAgentLinkDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winAgentLinkDo) Returning(value interface{}, columns ...string) IWinAgentLinkDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winAgentLinkDo) Not(conds ...gen.Condition) IWinAgentLinkDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winAgentLinkDo) Or(conds ...gen.Condition) IWinAgentLinkDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winAgentLinkDo) Select(conds ...field.Expr) IWinAgentLinkDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winAgentLinkDo) Where(conds ...gen.Condition) IWinAgentLinkDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winAgentLinkDo) Order(conds ...field.Expr) IWinAgentLinkDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winAgentLinkDo) Distinct(cols ...field.Expr) IWinAgentLinkDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winAgentLinkDo) Omit(cols ...field.Expr) IWinAgentLinkDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winAgentLinkDo) Join(table schema.Tabler, on ...field.Expr) IWinAgentLinkDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winAgentLinkDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinAgentLinkDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winAgentLinkDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinAgentLinkDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winAgentLinkDo) Group(cols ...field.Expr) IWinAgentLinkDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winAgentLinkDo) Having(conds ...gen.Condition) IWinAgentLinkDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winAgentLinkDo) Limit(limit int) IWinAgentLinkDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winAgentLinkDo) Offset(offset int) IWinAgentLinkDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winAgentLinkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinAgentLinkDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winAgentLinkDo) Unscoped() IWinAgentLinkDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winAgentLinkDo) Create(values ...*model.WinAgentLink) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winAgentLinkDo) CreateInBatches(values []*model.WinAgentLink, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winAgentLinkDo) Save(values ...*model.WinAgentLink) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winAgentLinkDo) First() (*model.WinAgentLink, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentLink), nil
	}
}

func (w winAgentLinkDo) Take() (*model.WinAgentLink, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentLink), nil
	}
}

func (w winAgentLinkDo) Last() (*model.WinAgentLink, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentLink), nil
	}
}

func (w winAgentLinkDo) Find() ([]*model.WinAgentLink, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinAgentLink), err
}

func (w winAgentLinkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAgentLink, err error) {
	buf := make([]*model.WinAgentLink, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winAgentLinkDo) FindInBatches(result *[]*model.WinAgentLink, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winAgentLinkDo) Attrs(attrs ...field.AssignExpr) IWinAgentLinkDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winAgentLinkDo) Assign(attrs ...field.AssignExpr) IWinAgentLinkDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winAgentLinkDo) Joins(fields ...field.RelationField) IWinAgentLinkDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winAgentLinkDo) Preload(fields ...field.RelationField) IWinAgentLinkDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winAgentLinkDo) FirstOrInit() (*model.WinAgentLink, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentLink), nil
	}
}

func (w winAgentLinkDo) FirstOrCreate() (*model.WinAgentLink, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAgentLink), nil
	}
}

func (w winAgentLinkDo) FindByPage(offset int, limit int) (result []*model.WinAgentLink, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w winAgentLinkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winAgentLinkDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winAgentLinkDo) Delete(models ...*model.WinAgentLink) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winAgentLinkDo) withDO(do gen.Dao) *winAgentLinkDo {
	w.DO = *do.(*gen.DO)
	return w
}
