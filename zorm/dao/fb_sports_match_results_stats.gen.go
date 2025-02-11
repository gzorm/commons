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

func newFbSportsMatchResultsStats(db *gorm.DB, opts ...gen.DOOption) fbSportsMatchResultsStats {
	_fbSportsMatchResultsStats := fbSportsMatchResultsStats{}

	_fbSportsMatchResultsStats.fbSportsMatchResultsStatsDo.UseDB(db, opts...)
	_fbSportsMatchResultsStats.fbSportsMatchResultsStatsDo.UseModel(&model.FbSportsMatchResultsStats{})

	tableName := _fbSportsMatchResultsStats.fbSportsMatchResultsStatsDo.TableName()
	_fbSportsMatchResultsStats.ALL = field.NewAsterisk(tableName)
	_fbSportsMatchResultsStats.ID = field.NewInt64(tableName, "id")
	_fbSportsMatchResultsStats.SportID = field.NewInt64(tableName, "sport_id")
	_fbSportsMatchResultsStats.MatchType = field.NewInt64(tableName, "match_type")
	_fbSportsMatchResultsStats.ResultCount = field.NewInt64(tableName, "result_count")
	_fbSportsMatchResultsStats.CreatedAt = field.NewTime(tableName, "created_at")
	_fbSportsMatchResultsStats.UpdatedAt = field.NewTime(tableName, "updated_at")

	_fbSportsMatchResultsStats.fillFieldMap()

	return _fbSportsMatchResultsStats
}

// fbSportsMatchResultsStats 赛果统计表
type fbSportsMatchResultsStats struct {
	fbSportsMatchResultsStatsDo

	ALL         field.Asterisk
	ID          field.Int64 // 主键ID
	SportID     field.Int64 // 运动ID，如足球、篮球
	MatchType   field.Int64 // 赛事类型，如常规赛、季后赛
	ResultCount field.Int64 // 赛果统计总数
	CreatedAt   field.Time  // 记录创建时间
	UpdatedAt   field.Time  // 记录更新时间

	fieldMap map[string]field.Expr
}

func (f fbSportsMatchResultsStats) Table(newTableName string) *fbSportsMatchResultsStats {
	f.fbSportsMatchResultsStatsDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbSportsMatchResultsStats) As(alias string) *fbSportsMatchResultsStats {
	f.fbSportsMatchResultsStatsDo.DO = *(f.fbSportsMatchResultsStatsDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbSportsMatchResultsStats) updateTableName(table string) *fbSportsMatchResultsStats {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.SportID = field.NewInt64(table, "sport_id")
	f.MatchType = field.NewInt64(table, "match_type")
	f.ResultCount = field.NewInt64(table, "result_count")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")

	f.fillFieldMap()

	return f
}

func (f *fbSportsMatchResultsStats) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbSportsMatchResultsStats) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 6)
	f.fieldMap["id"] = f.ID
	f.fieldMap["sport_id"] = f.SportID
	f.fieldMap["match_type"] = f.MatchType
	f.fieldMap["result_count"] = f.ResultCount
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
}

func (f fbSportsMatchResultsStats) clone(db *gorm.DB) fbSportsMatchResultsStats {
	f.fbSportsMatchResultsStatsDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbSportsMatchResultsStats) replaceDB(db *gorm.DB) fbSportsMatchResultsStats {
	f.fbSportsMatchResultsStatsDo.ReplaceDB(db)
	return f
}

type fbSportsMatchResultsStatsDo struct{ gen.DO }

type IFbSportsMatchResultsStatsDo interface {
	gen.SubQuery
	Debug() IFbSportsMatchResultsStatsDo
	WithContext(ctx context.Context) IFbSportsMatchResultsStatsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbSportsMatchResultsStatsDo
	WriteDB() IFbSportsMatchResultsStatsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbSportsMatchResultsStatsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbSportsMatchResultsStatsDo
	Not(conds ...gen.Condition) IFbSportsMatchResultsStatsDo
	Or(conds ...gen.Condition) IFbSportsMatchResultsStatsDo
	Select(conds ...field.Expr) IFbSportsMatchResultsStatsDo
	Where(conds ...gen.Condition) IFbSportsMatchResultsStatsDo
	Order(conds ...field.Expr) IFbSportsMatchResultsStatsDo
	Distinct(cols ...field.Expr) IFbSportsMatchResultsStatsDo
	Omit(cols ...field.Expr) IFbSportsMatchResultsStatsDo
	Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchResultsStatsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchResultsStatsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchResultsStatsDo
	Group(cols ...field.Expr) IFbSportsMatchResultsStatsDo
	Having(conds ...gen.Condition) IFbSportsMatchResultsStatsDo
	Limit(limit int) IFbSportsMatchResultsStatsDo
	Offset(offset int) IFbSportsMatchResultsStatsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchResultsStatsDo
	Unscoped() IFbSportsMatchResultsStatsDo
	Create(values ...*model.FbSportsMatchResultsStats) error
	CreateInBatches(values []*model.FbSportsMatchResultsStats, batchSize int) error
	Save(values ...*model.FbSportsMatchResultsStats) error
	First() (*model.FbSportsMatchResultsStats, error)
	Take() (*model.FbSportsMatchResultsStats, error)
	Last() (*model.FbSportsMatchResultsStats, error)
	Find() ([]*model.FbSportsMatchResultsStats, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatchResultsStats, err error)
	FindInBatches(result *[]*model.FbSportsMatchResultsStats, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbSportsMatchResultsStats) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbSportsMatchResultsStatsDo
	Assign(attrs ...field.AssignExpr) IFbSportsMatchResultsStatsDo
	Joins(fields ...field.RelationField) IFbSportsMatchResultsStatsDo
	Preload(fields ...field.RelationField) IFbSportsMatchResultsStatsDo
	FirstOrInit() (*model.FbSportsMatchResultsStats, error)
	FirstOrCreate() (*model.FbSportsMatchResultsStats, error)
	FindByPage(offset int, limit int) (result []*model.FbSportsMatchResultsStats, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbSportsMatchResultsStatsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbSportsMatchResultsStatsDo) Debug() IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Debug())
}

func (f fbSportsMatchResultsStatsDo) WithContext(ctx context.Context) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbSportsMatchResultsStatsDo) ReadDB() IFbSportsMatchResultsStatsDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbSportsMatchResultsStatsDo) WriteDB() IFbSportsMatchResultsStatsDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbSportsMatchResultsStatsDo) Session(config *gorm.Session) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbSportsMatchResultsStatsDo) Clauses(conds ...clause.Expression) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbSportsMatchResultsStatsDo) Returning(value interface{}, columns ...string) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbSportsMatchResultsStatsDo) Not(conds ...gen.Condition) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbSportsMatchResultsStatsDo) Or(conds ...gen.Condition) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbSportsMatchResultsStatsDo) Select(conds ...field.Expr) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbSportsMatchResultsStatsDo) Where(conds ...gen.Condition) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbSportsMatchResultsStatsDo) Order(conds ...field.Expr) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbSportsMatchResultsStatsDo) Distinct(cols ...field.Expr) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbSportsMatchResultsStatsDo) Omit(cols ...field.Expr) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbSportsMatchResultsStatsDo) Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbSportsMatchResultsStatsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbSportsMatchResultsStatsDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbSportsMatchResultsStatsDo) Group(cols ...field.Expr) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbSportsMatchResultsStatsDo) Having(conds ...gen.Condition) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbSportsMatchResultsStatsDo) Limit(limit int) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbSportsMatchResultsStatsDo) Offset(offset int) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbSportsMatchResultsStatsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbSportsMatchResultsStatsDo) Unscoped() IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbSportsMatchResultsStatsDo) Create(values ...*model.FbSportsMatchResultsStats) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbSportsMatchResultsStatsDo) CreateInBatches(values []*model.FbSportsMatchResultsStats, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbSportsMatchResultsStatsDo) Save(values ...*model.FbSportsMatchResultsStats) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbSportsMatchResultsStatsDo) First() (*model.FbSportsMatchResultsStats, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchResultsStats), nil
	}
}

func (f fbSportsMatchResultsStatsDo) Take() (*model.FbSportsMatchResultsStats, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchResultsStats), nil
	}
}

func (f fbSportsMatchResultsStatsDo) Last() (*model.FbSportsMatchResultsStats, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchResultsStats), nil
	}
}

func (f fbSportsMatchResultsStatsDo) Find() ([]*model.FbSportsMatchResultsStats, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbSportsMatchResultsStats), err
}

func (f fbSportsMatchResultsStatsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatchResultsStats, err error) {
	buf := make([]*model.FbSportsMatchResultsStats, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbSportsMatchResultsStatsDo) FindInBatches(result *[]*model.FbSportsMatchResultsStats, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbSportsMatchResultsStatsDo) Attrs(attrs ...field.AssignExpr) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbSportsMatchResultsStatsDo) Assign(attrs ...field.AssignExpr) IFbSportsMatchResultsStatsDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbSportsMatchResultsStatsDo) Joins(fields ...field.RelationField) IFbSportsMatchResultsStatsDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbSportsMatchResultsStatsDo) Preload(fields ...field.RelationField) IFbSportsMatchResultsStatsDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbSportsMatchResultsStatsDo) FirstOrInit() (*model.FbSportsMatchResultsStats, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchResultsStats), nil
	}
}

func (f fbSportsMatchResultsStatsDo) FirstOrCreate() (*model.FbSportsMatchResultsStats, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchResultsStats), nil
	}
}

func (f fbSportsMatchResultsStatsDo) FindByPage(offset int, limit int) (result []*model.FbSportsMatchResultsStats, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fbSportsMatchResultsStatsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbSportsMatchResultsStatsDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbSportsMatchResultsStatsDo) Delete(models ...*model.FbSportsMatchResultsStats) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbSportsMatchResultsStatsDo) withDO(do gen.Dao) *fbSportsMatchResultsStatsDo {
	f.DO = *do.(*gen.DO)
	return f
}
