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

func newWinPromotionRequirements(db *gorm.DB, opts ...gen.DOOption) winPromotionRequirements {
	_winPromotionRequirements := winPromotionRequirements{}

	_winPromotionRequirements.winPromotionRequirementsDo.UseDB(db, opts...)
	_winPromotionRequirements.winPromotionRequirementsDo.UseModel(&model.WinPromotionRequirements{})

	tableName := _winPromotionRequirements.winPromotionRequirementsDo.TableName()
	_winPromotionRequirements.ALL = field.NewAsterisk(tableName)
	_winPromotionRequirements.ID = field.NewInt64(tableName, "id")
	_winPromotionRequirements.PeopleNum = field.NewString(tableName, "people_num")
	_winPromotionRequirements.MonthRegisteredUsers = field.NewInt64(tableName, "month_registered_users")
	_winPromotionRequirements.MonthPaidUsers = field.NewInt64(tableName, "month_paid_users")
	_winPromotionRequirements.MonthNegativeProfit = field.NewField(tableName, "month_negative_profit")
	_winPromotionRequirements.CreatedAt = field.NewInt64(tableName, "created_at")
	_winPromotionRequirements.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winPromotionRequirements.fillFieldMap()

	return _winPromotionRequirements
}

// winPromotionRequirements 团队长团队人员平均绩效要求
type winPromotionRequirements struct {
	winPromotionRequirementsDo

	ALL                  field.Asterisk
	ID                   field.Int64
	PeopleNum            field.String // 人员
	MonthRegisteredUsers field.Int64  // 注册用户/月
	MonthPaidUsers       field.Int64  // 付费用户/月
	MonthNegativeProfit  field.Field  // 负盈利总额/月
	CreatedAt            field.Int64  // 创建时间
	UpdatedAt            field.Int64  // 更新时间

	fieldMap map[string]field.Expr
}

func (w winPromotionRequirements) Table(newTableName string) *winPromotionRequirements {
	w.winPromotionRequirementsDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winPromotionRequirements) As(alias string) *winPromotionRequirements {
	w.winPromotionRequirementsDo.DO = *(w.winPromotionRequirementsDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winPromotionRequirements) updateTableName(table string) *winPromotionRequirements {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.PeopleNum = field.NewString(table, "people_num")
	w.MonthRegisteredUsers = field.NewInt64(table, "month_registered_users")
	w.MonthPaidUsers = field.NewInt64(table, "month_paid_users")
	w.MonthNegativeProfit = field.NewField(table, "month_negative_profit")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winPromotionRequirements) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winPromotionRequirements) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 7)
	w.fieldMap["id"] = w.ID
	w.fieldMap["people_num"] = w.PeopleNum
	w.fieldMap["month_registered_users"] = w.MonthRegisteredUsers
	w.fieldMap["month_paid_users"] = w.MonthPaidUsers
	w.fieldMap["month_negative_profit"] = w.MonthNegativeProfit
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winPromotionRequirements) clone(db *gorm.DB) winPromotionRequirements {
	w.winPromotionRequirementsDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winPromotionRequirements) replaceDB(db *gorm.DB) winPromotionRequirements {
	w.winPromotionRequirementsDo.ReplaceDB(db)
	return w
}

type winPromotionRequirementsDo struct{ gen.DO }

type IWinPromotionRequirementsDo interface {
	gen.SubQuery
	Debug() IWinPromotionRequirementsDo
	WithContext(ctx context.Context) IWinPromotionRequirementsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinPromotionRequirementsDo
	WriteDB() IWinPromotionRequirementsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinPromotionRequirementsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinPromotionRequirementsDo
	Not(conds ...gen.Condition) IWinPromotionRequirementsDo
	Or(conds ...gen.Condition) IWinPromotionRequirementsDo
	Select(conds ...field.Expr) IWinPromotionRequirementsDo
	Where(conds ...gen.Condition) IWinPromotionRequirementsDo
	Order(conds ...field.Expr) IWinPromotionRequirementsDo
	Distinct(cols ...field.Expr) IWinPromotionRequirementsDo
	Omit(cols ...field.Expr) IWinPromotionRequirementsDo
	Join(table schema.Tabler, on ...field.Expr) IWinPromotionRequirementsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinPromotionRequirementsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinPromotionRequirementsDo
	Group(cols ...field.Expr) IWinPromotionRequirementsDo
	Having(conds ...gen.Condition) IWinPromotionRequirementsDo
	Limit(limit int) IWinPromotionRequirementsDo
	Offset(offset int) IWinPromotionRequirementsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinPromotionRequirementsDo
	Unscoped() IWinPromotionRequirementsDo
	Create(values ...*model.WinPromotionRequirements) error
	CreateInBatches(values []*model.WinPromotionRequirements, batchSize int) error
	Save(values ...*model.WinPromotionRequirements) error
	First() (*model.WinPromotionRequirements, error)
	Take() (*model.WinPromotionRequirements, error)
	Last() (*model.WinPromotionRequirements, error)
	Find() ([]*model.WinPromotionRequirements, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinPromotionRequirements, err error)
	FindInBatches(result *[]*model.WinPromotionRequirements, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinPromotionRequirements) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinPromotionRequirementsDo
	Assign(attrs ...field.AssignExpr) IWinPromotionRequirementsDo
	Joins(fields ...field.RelationField) IWinPromotionRequirementsDo
	Preload(fields ...field.RelationField) IWinPromotionRequirementsDo
	FirstOrInit() (*model.WinPromotionRequirements, error)
	FirstOrCreate() (*model.WinPromotionRequirements, error)
	FindByPage(offset int, limit int) (result []*model.WinPromotionRequirements, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinPromotionRequirementsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winPromotionRequirementsDo) Debug() IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Debug())
}

func (w winPromotionRequirementsDo) WithContext(ctx context.Context) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winPromotionRequirementsDo) ReadDB() IWinPromotionRequirementsDo {
	return w.Clauses(dbresolver.Read)
}

func (w winPromotionRequirementsDo) WriteDB() IWinPromotionRequirementsDo {
	return w.Clauses(dbresolver.Write)
}

func (w winPromotionRequirementsDo) Session(config *gorm.Session) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Session(config))
}

func (w winPromotionRequirementsDo) Clauses(conds ...clause.Expression) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winPromotionRequirementsDo) Returning(value interface{}, columns ...string) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winPromotionRequirementsDo) Not(conds ...gen.Condition) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winPromotionRequirementsDo) Or(conds ...gen.Condition) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winPromotionRequirementsDo) Select(conds ...field.Expr) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winPromotionRequirementsDo) Where(conds ...gen.Condition) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winPromotionRequirementsDo) Order(conds ...field.Expr) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winPromotionRequirementsDo) Distinct(cols ...field.Expr) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winPromotionRequirementsDo) Omit(cols ...field.Expr) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winPromotionRequirementsDo) Join(table schema.Tabler, on ...field.Expr) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winPromotionRequirementsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winPromotionRequirementsDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winPromotionRequirementsDo) Group(cols ...field.Expr) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winPromotionRequirementsDo) Having(conds ...gen.Condition) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winPromotionRequirementsDo) Limit(limit int) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winPromotionRequirementsDo) Offset(offset int) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winPromotionRequirementsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winPromotionRequirementsDo) Unscoped() IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winPromotionRequirementsDo) Create(values ...*model.WinPromotionRequirements) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winPromotionRequirementsDo) CreateInBatches(values []*model.WinPromotionRequirements, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winPromotionRequirementsDo) Save(values ...*model.WinPromotionRequirements) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winPromotionRequirementsDo) First() (*model.WinPromotionRequirements, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotionRequirements), nil
	}
}

func (w winPromotionRequirementsDo) Take() (*model.WinPromotionRequirements, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotionRequirements), nil
	}
}

func (w winPromotionRequirementsDo) Last() (*model.WinPromotionRequirements, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotionRequirements), nil
	}
}

func (w winPromotionRequirementsDo) Find() ([]*model.WinPromotionRequirements, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinPromotionRequirements), err
}

func (w winPromotionRequirementsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinPromotionRequirements, err error) {
	buf := make([]*model.WinPromotionRequirements, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winPromotionRequirementsDo) FindInBatches(result *[]*model.WinPromotionRequirements, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winPromotionRequirementsDo) Attrs(attrs ...field.AssignExpr) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winPromotionRequirementsDo) Assign(attrs ...field.AssignExpr) IWinPromotionRequirementsDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winPromotionRequirementsDo) Joins(fields ...field.RelationField) IWinPromotionRequirementsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winPromotionRequirementsDo) Preload(fields ...field.RelationField) IWinPromotionRequirementsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winPromotionRequirementsDo) FirstOrInit() (*model.WinPromotionRequirements, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotionRequirements), nil
	}
}

func (w winPromotionRequirementsDo) FirstOrCreate() (*model.WinPromotionRequirements, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotionRequirements), nil
	}
}

func (w winPromotionRequirementsDo) FindByPage(offset int, limit int) (result []*model.WinPromotionRequirements, count int64, err error) {
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

func (w winPromotionRequirementsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winPromotionRequirementsDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winPromotionRequirementsDo) Delete(models ...*model.WinPromotionRequirements) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winPromotionRequirementsDo) withDO(do gen.Dao) *winPromotionRequirementsDo {
	w.DO = *do.(*gen.DO)
	return w
}
