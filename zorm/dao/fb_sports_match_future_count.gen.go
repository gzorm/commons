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

func newFbSportsMatchFutureCount(db *gorm.DB, opts ...gen.DOOption) fbSportsMatchFutureCount {
	_fbSportsMatchFutureCount := fbSportsMatchFutureCount{}

	_fbSportsMatchFutureCount.fbSportsMatchFutureCountDo.UseDB(db, opts...)
	_fbSportsMatchFutureCount.fbSportsMatchFutureCountDo.UseModel(&model.FbSportsMatchFutureCount{})

	tableName := _fbSportsMatchFutureCount.fbSportsMatchFutureCountDo.TableName()
	_fbSportsMatchFutureCount.ALL = field.NewAsterisk(tableName)
	_fbSportsMatchFutureCount.ID = field.NewInt64(tableName, "id")
	_fbSportsMatchFutureCount.SportID = field.NewInt64(tableName, "sport_id")
	_fbSportsMatchFutureCount.FutureMatchData = field.NewString(tableName, "future_match_data")
	_fbSportsMatchFutureCount.BeyondFutureMatchData = field.NewString(tableName, "beyond_future_match_data")
	_fbSportsMatchFutureCount.FutureDays = field.NewInt64(tableName, "future_days")
	_fbSportsMatchFutureCount.CreatedAt = field.NewInt64(tableName, "created_at")
	_fbSportsMatchFutureCount.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_fbSportsMatchFutureCount.fillFieldMap()

	return _fbSportsMatchFutureCount
}

// fbSportsMatchFutureCount 存储未来天数的赛事统计数据
type fbSportsMatchFutureCount struct {
	fbSportsMatchFutureCountDo

	ALL                   field.Asterisk
	ID                    field.Int64  // 主键ID
	SportID               field.Int64  // 运动类型ID
	FutureMatchData       field.String // 未来 {future_days} 天的赛事统计数据
	BeyondFutureMatchData field.String // 超过 {future_days} 天后的赛事统计数据
	FutureDays            field.Int64  // 统计的未来天数
	CreatedAt             field.Int64  // 创建时间（时间戳）
	UpdatedAt             field.Int64  // 更新时间（时间戳）

	fieldMap map[string]field.Expr
}

func (f fbSportsMatchFutureCount) Table(newTableName string) *fbSportsMatchFutureCount {
	f.fbSportsMatchFutureCountDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbSportsMatchFutureCount) As(alias string) *fbSportsMatchFutureCount {
	f.fbSportsMatchFutureCountDo.DO = *(f.fbSportsMatchFutureCountDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbSportsMatchFutureCount) updateTableName(table string) *fbSportsMatchFutureCount {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.SportID = field.NewInt64(table, "sport_id")
	f.FutureMatchData = field.NewString(table, "future_match_data")
	f.BeyondFutureMatchData = field.NewString(table, "beyond_future_match_data")
	f.FutureDays = field.NewInt64(table, "future_days")
	f.CreatedAt = field.NewInt64(table, "created_at")
	f.UpdatedAt = field.NewInt64(table, "updated_at")

	f.fillFieldMap()

	return f
}

func (f *fbSportsMatchFutureCount) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbSportsMatchFutureCount) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 7)
	f.fieldMap["id"] = f.ID
	f.fieldMap["sport_id"] = f.SportID
	f.fieldMap["future_match_data"] = f.FutureMatchData
	f.fieldMap["beyond_future_match_data"] = f.BeyondFutureMatchData
	f.fieldMap["future_days"] = f.FutureDays
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
}

func (f fbSportsMatchFutureCount) clone(db *gorm.DB) fbSportsMatchFutureCount {
	f.fbSportsMatchFutureCountDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbSportsMatchFutureCount) replaceDB(db *gorm.DB) fbSportsMatchFutureCount {
	f.fbSportsMatchFutureCountDo.ReplaceDB(db)
	return f
}

type fbSportsMatchFutureCountDo struct{ gen.DO }

type IFbSportsMatchFutureCountDo interface {
	gen.SubQuery
	Debug() IFbSportsMatchFutureCountDo
	WithContext(ctx context.Context) IFbSportsMatchFutureCountDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbSportsMatchFutureCountDo
	WriteDB() IFbSportsMatchFutureCountDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbSportsMatchFutureCountDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbSportsMatchFutureCountDo
	Not(conds ...gen.Condition) IFbSportsMatchFutureCountDo
	Or(conds ...gen.Condition) IFbSportsMatchFutureCountDo
	Select(conds ...field.Expr) IFbSportsMatchFutureCountDo
	Where(conds ...gen.Condition) IFbSportsMatchFutureCountDo
	Order(conds ...field.Expr) IFbSportsMatchFutureCountDo
	Distinct(cols ...field.Expr) IFbSportsMatchFutureCountDo
	Omit(cols ...field.Expr) IFbSportsMatchFutureCountDo
	Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchFutureCountDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchFutureCountDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchFutureCountDo
	Group(cols ...field.Expr) IFbSportsMatchFutureCountDo
	Having(conds ...gen.Condition) IFbSportsMatchFutureCountDo
	Limit(limit int) IFbSportsMatchFutureCountDo
	Offset(offset int) IFbSportsMatchFutureCountDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchFutureCountDo
	Unscoped() IFbSportsMatchFutureCountDo
	Create(values ...*model.FbSportsMatchFutureCount) error
	CreateInBatches(values []*model.FbSportsMatchFutureCount, batchSize int) error
	Save(values ...*model.FbSportsMatchFutureCount) error
	First() (*model.FbSportsMatchFutureCount, error)
	Take() (*model.FbSportsMatchFutureCount, error)
	Last() (*model.FbSportsMatchFutureCount, error)
	Find() ([]*model.FbSportsMatchFutureCount, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatchFutureCount, err error)
	FindInBatches(result *[]*model.FbSportsMatchFutureCount, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbSportsMatchFutureCount) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbSportsMatchFutureCountDo
	Assign(attrs ...field.AssignExpr) IFbSportsMatchFutureCountDo
	Joins(fields ...field.RelationField) IFbSportsMatchFutureCountDo
	Preload(fields ...field.RelationField) IFbSportsMatchFutureCountDo
	FirstOrInit() (*model.FbSportsMatchFutureCount, error)
	FirstOrCreate() (*model.FbSportsMatchFutureCount, error)
	FindByPage(offset int, limit int) (result []*model.FbSportsMatchFutureCount, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbSportsMatchFutureCountDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbSportsMatchFutureCountDo) Debug() IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Debug())
}

func (f fbSportsMatchFutureCountDo) WithContext(ctx context.Context) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbSportsMatchFutureCountDo) ReadDB() IFbSportsMatchFutureCountDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbSportsMatchFutureCountDo) WriteDB() IFbSportsMatchFutureCountDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbSportsMatchFutureCountDo) Session(config *gorm.Session) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbSportsMatchFutureCountDo) Clauses(conds ...clause.Expression) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbSportsMatchFutureCountDo) Returning(value interface{}, columns ...string) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbSportsMatchFutureCountDo) Not(conds ...gen.Condition) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbSportsMatchFutureCountDo) Or(conds ...gen.Condition) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbSportsMatchFutureCountDo) Select(conds ...field.Expr) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbSportsMatchFutureCountDo) Where(conds ...gen.Condition) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbSportsMatchFutureCountDo) Order(conds ...field.Expr) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbSportsMatchFutureCountDo) Distinct(cols ...field.Expr) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbSportsMatchFutureCountDo) Omit(cols ...field.Expr) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbSportsMatchFutureCountDo) Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbSportsMatchFutureCountDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbSportsMatchFutureCountDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbSportsMatchFutureCountDo) Group(cols ...field.Expr) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbSportsMatchFutureCountDo) Having(conds ...gen.Condition) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbSportsMatchFutureCountDo) Limit(limit int) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbSportsMatchFutureCountDo) Offset(offset int) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbSportsMatchFutureCountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbSportsMatchFutureCountDo) Unscoped() IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbSportsMatchFutureCountDo) Create(values ...*model.FbSportsMatchFutureCount) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbSportsMatchFutureCountDo) CreateInBatches(values []*model.FbSportsMatchFutureCount, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbSportsMatchFutureCountDo) Save(values ...*model.FbSportsMatchFutureCount) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbSportsMatchFutureCountDo) First() (*model.FbSportsMatchFutureCount, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchFutureCount), nil
	}
}

func (f fbSportsMatchFutureCountDo) Take() (*model.FbSportsMatchFutureCount, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchFutureCount), nil
	}
}

func (f fbSportsMatchFutureCountDo) Last() (*model.FbSportsMatchFutureCount, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchFutureCount), nil
	}
}

func (f fbSportsMatchFutureCountDo) Find() ([]*model.FbSportsMatchFutureCount, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbSportsMatchFutureCount), err
}

func (f fbSportsMatchFutureCountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatchFutureCount, err error) {
	buf := make([]*model.FbSportsMatchFutureCount, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbSportsMatchFutureCountDo) FindInBatches(result *[]*model.FbSportsMatchFutureCount, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbSportsMatchFutureCountDo) Attrs(attrs ...field.AssignExpr) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbSportsMatchFutureCountDo) Assign(attrs ...field.AssignExpr) IFbSportsMatchFutureCountDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbSportsMatchFutureCountDo) Joins(fields ...field.RelationField) IFbSportsMatchFutureCountDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbSportsMatchFutureCountDo) Preload(fields ...field.RelationField) IFbSportsMatchFutureCountDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbSportsMatchFutureCountDo) FirstOrInit() (*model.FbSportsMatchFutureCount, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchFutureCount), nil
	}
}

func (f fbSportsMatchFutureCountDo) FirstOrCreate() (*model.FbSportsMatchFutureCount, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchFutureCount), nil
	}
}

func (f fbSportsMatchFutureCountDo) FindByPage(offset int, limit int) (result []*model.FbSportsMatchFutureCount, count int64, err error) {
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

func (f fbSportsMatchFutureCountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbSportsMatchFutureCountDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbSportsMatchFutureCountDo) Delete(models ...*model.FbSportsMatchFutureCount) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbSportsMatchFutureCountDo) withDO(do gen.Dao) *fbSportsMatchFutureCountDo {
	f.DO = *do.(*gen.DO)
	return f
}
