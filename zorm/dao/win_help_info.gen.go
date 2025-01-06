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

	"github.com/gzorm/common/zorm/model"
)

func newWinHelpInfo(db *gorm.DB, opts ...gen.DOOption) winHelpInfo {
	_winHelpInfo := winHelpInfo{}

	_winHelpInfo.winHelpInfoDo.UseDB(db, opts...)
	_winHelpInfo.winHelpInfoDo.UseModel(&model.WinHelpInfo{})

	tableName := _winHelpInfo.winHelpInfoDo.TableName()
	_winHelpInfo.ALL = field.NewAsterisk(tableName)
	_winHelpInfo.ID = field.NewInt64(tableName, "id")
	_winHelpInfo.HelpTypeID = field.NewInt64(tableName, "help_type_id")
	_winHelpInfo.Language = field.NewString(tableName, "language")
	_winHelpInfo.Title = field.NewString(tableName, "title")
	_winHelpInfo.Sort = field.NewInt64(tableName, "sort")
	_winHelpInfo.Status = field.NewInt64(tableName, "status")
	_winHelpInfo.Content = field.NewString(tableName, "content")
	_winHelpInfo.CreateBy = field.NewString(tableName, "create_by")
	_winHelpInfo.UpdateBy = field.NewString(tableName, "update_by")
	_winHelpInfo.CreatedAt = field.NewInt64(tableName, "created_at")
	_winHelpInfo.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winHelpInfo.fillFieldMap()

	return _winHelpInfo
}

// winHelpInfo 帮助详情
type winHelpInfo struct {
	winHelpInfoDo

	ALL        field.Asterisk
	ID         field.Int64  // 主键编号
	HelpTypeID field.Int64  // 帮助类型id
	Language   field.String // 语言
	Title      field.String // 标题
	Sort       field.Int64  // 排序
	Status     field.Int64  // 状态:1-启用 0-停用
	Content    field.String // 内容
	CreateBy   field.String // 创建者
	UpdateBy   field.String // 更新人
	CreatedAt  field.Int64
	UpdatedAt  field.Int64

	fieldMap map[string]field.Expr
}

func (w winHelpInfo) Table(newTableName string) *winHelpInfo {
	w.winHelpInfoDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winHelpInfo) As(alias string) *winHelpInfo {
	w.winHelpInfoDo.DO = *(w.winHelpInfoDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winHelpInfo) updateTableName(table string) *winHelpInfo {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.HelpTypeID = field.NewInt64(table, "help_type_id")
	w.Language = field.NewString(table, "language")
	w.Title = field.NewString(table, "title")
	w.Sort = field.NewInt64(table, "sort")
	w.Status = field.NewInt64(table, "status")
	w.Content = field.NewString(table, "content")
	w.CreateBy = field.NewString(table, "create_by")
	w.UpdateBy = field.NewString(table, "update_by")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winHelpInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winHelpInfo) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 11)
	w.fieldMap["id"] = w.ID
	w.fieldMap["help_type_id"] = w.HelpTypeID
	w.fieldMap["language"] = w.Language
	w.fieldMap["title"] = w.Title
	w.fieldMap["sort"] = w.Sort
	w.fieldMap["status"] = w.Status
	w.fieldMap["content"] = w.Content
	w.fieldMap["create_by"] = w.CreateBy
	w.fieldMap["update_by"] = w.UpdateBy
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winHelpInfo) clone(db *gorm.DB) winHelpInfo {
	w.winHelpInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winHelpInfo) replaceDB(db *gorm.DB) winHelpInfo {
	w.winHelpInfoDo.ReplaceDB(db)
	return w
}

type winHelpInfoDo struct{ gen.DO }

type IWinHelpInfoDo interface {
	gen.SubQuery
	Debug() IWinHelpInfoDo
	WithContext(ctx context.Context) IWinHelpInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinHelpInfoDo
	WriteDB() IWinHelpInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinHelpInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinHelpInfoDo
	Not(conds ...gen.Condition) IWinHelpInfoDo
	Or(conds ...gen.Condition) IWinHelpInfoDo
	Select(conds ...field.Expr) IWinHelpInfoDo
	Where(conds ...gen.Condition) IWinHelpInfoDo
	Order(conds ...field.Expr) IWinHelpInfoDo
	Distinct(cols ...field.Expr) IWinHelpInfoDo
	Omit(cols ...field.Expr) IWinHelpInfoDo
	Join(table schema.Tabler, on ...field.Expr) IWinHelpInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinHelpInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinHelpInfoDo
	Group(cols ...field.Expr) IWinHelpInfoDo
	Having(conds ...gen.Condition) IWinHelpInfoDo
	Limit(limit int) IWinHelpInfoDo
	Offset(offset int) IWinHelpInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinHelpInfoDo
	Unscoped() IWinHelpInfoDo
	Create(values ...*model.WinHelpInfo) error
	CreateInBatches(values []*model.WinHelpInfo, batchSize int) error
	Save(values ...*model.WinHelpInfo) error
	First() (*model.WinHelpInfo, error)
	Take() (*model.WinHelpInfo, error)
	Last() (*model.WinHelpInfo, error)
	Find() ([]*model.WinHelpInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinHelpInfo, err error)
	FindInBatches(result *[]*model.WinHelpInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinHelpInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinHelpInfoDo
	Assign(attrs ...field.AssignExpr) IWinHelpInfoDo
	Joins(fields ...field.RelationField) IWinHelpInfoDo
	Preload(fields ...field.RelationField) IWinHelpInfoDo
	FirstOrInit() (*model.WinHelpInfo, error)
	FirstOrCreate() (*model.WinHelpInfo, error)
	FindByPage(offset int, limit int) (result []*model.WinHelpInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinHelpInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winHelpInfoDo) Debug() IWinHelpInfoDo {
	return w.withDO(w.DO.Debug())
}

func (w winHelpInfoDo) WithContext(ctx context.Context) IWinHelpInfoDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winHelpInfoDo) ReadDB() IWinHelpInfoDo {
	return w.Clauses(dbresolver.Read)
}

func (w winHelpInfoDo) WriteDB() IWinHelpInfoDo {
	return w.Clauses(dbresolver.Write)
}

func (w winHelpInfoDo) Session(config *gorm.Session) IWinHelpInfoDo {
	return w.withDO(w.DO.Session(config))
}

func (w winHelpInfoDo) Clauses(conds ...clause.Expression) IWinHelpInfoDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winHelpInfoDo) Returning(value interface{}, columns ...string) IWinHelpInfoDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winHelpInfoDo) Not(conds ...gen.Condition) IWinHelpInfoDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winHelpInfoDo) Or(conds ...gen.Condition) IWinHelpInfoDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winHelpInfoDo) Select(conds ...field.Expr) IWinHelpInfoDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winHelpInfoDo) Where(conds ...gen.Condition) IWinHelpInfoDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winHelpInfoDo) Order(conds ...field.Expr) IWinHelpInfoDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winHelpInfoDo) Distinct(cols ...field.Expr) IWinHelpInfoDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winHelpInfoDo) Omit(cols ...field.Expr) IWinHelpInfoDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winHelpInfoDo) Join(table schema.Tabler, on ...field.Expr) IWinHelpInfoDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winHelpInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinHelpInfoDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winHelpInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinHelpInfoDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winHelpInfoDo) Group(cols ...field.Expr) IWinHelpInfoDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winHelpInfoDo) Having(conds ...gen.Condition) IWinHelpInfoDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winHelpInfoDo) Limit(limit int) IWinHelpInfoDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winHelpInfoDo) Offset(offset int) IWinHelpInfoDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winHelpInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinHelpInfoDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winHelpInfoDo) Unscoped() IWinHelpInfoDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winHelpInfoDo) Create(values ...*model.WinHelpInfo) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winHelpInfoDo) CreateInBatches(values []*model.WinHelpInfo, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winHelpInfoDo) Save(values ...*model.WinHelpInfo) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winHelpInfoDo) First() (*model.WinHelpInfo, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinHelpInfo), nil
	}
}

func (w winHelpInfoDo) Take() (*model.WinHelpInfo, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinHelpInfo), nil
	}
}

func (w winHelpInfoDo) Last() (*model.WinHelpInfo, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinHelpInfo), nil
	}
}

func (w winHelpInfoDo) Find() ([]*model.WinHelpInfo, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinHelpInfo), err
}

func (w winHelpInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinHelpInfo, err error) {
	buf := make([]*model.WinHelpInfo, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winHelpInfoDo) FindInBatches(result *[]*model.WinHelpInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winHelpInfoDo) Attrs(attrs ...field.AssignExpr) IWinHelpInfoDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winHelpInfoDo) Assign(attrs ...field.AssignExpr) IWinHelpInfoDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winHelpInfoDo) Joins(fields ...field.RelationField) IWinHelpInfoDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winHelpInfoDo) Preload(fields ...field.RelationField) IWinHelpInfoDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winHelpInfoDo) FirstOrInit() (*model.WinHelpInfo, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinHelpInfo), nil
	}
}

func (w winHelpInfoDo) FirstOrCreate() (*model.WinHelpInfo, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinHelpInfo), nil
	}
}

func (w winHelpInfoDo) FindByPage(offset int, limit int) (result []*model.WinHelpInfo, count int64, err error) {
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

func (w winHelpInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winHelpInfoDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winHelpInfoDo) Delete(models ...*model.WinHelpInfo) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winHelpInfoDo) withDO(do gen.Dao) *winHelpInfoDo {
	w.DO = *do.(*gen.DO)
	return w
}