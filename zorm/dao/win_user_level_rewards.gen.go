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

func newWinUserLevelReward(db *gorm.DB, opts ...gen.DOOption) winUserLevelReward {
	_winUserLevelReward := winUserLevelReward{}

	_winUserLevelReward.winUserLevelRewardDo.UseDB(db, opts...)
	_winUserLevelReward.winUserLevelRewardDo.UseModel(&model.WinUserLevelReward{})

	tableName := _winUserLevelReward.winUserLevelRewardDo.TableName()
	_winUserLevelReward.ALL = field.NewAsterisk(tableName)
	_winUserLevelReward.ID = field.NewInt64(tableName, "id")
	_winUserLevelReward.UID = field.NewInt64(tableName, "uid")
	_winUserLevelReward.Username = field.NewString(tableName, "username")
	_winUserLevelReward.Coin = field.NewField(tableName, "coin")
	_winUserLevelReward.CoinBefore = field.NewField(tableName, "coin_before")
	_winUserLevelReward.Category = field.NewInt64(tableName, "category")
	_winUserLevelReward.ReceiveAt = field.NewInt64(tableName, "receive_at")
	_winUserLevelReward.UserLevelID = field.NewInt64(tableName, "user_level_id")
	_winUserLevelReward.FlowClaim = field.NewInt64(tableName, "flow_claim")
	_winUserLevelReward.Codes = field.NewField(tableName, "codes")
	_winUserLevelReward.EndedAt = field.NewInt64(tableName, "ended_at")
	_winUserLevelReward.Status = field.NewInt64(tableName, "status")
	_winUserLevelReward.CreatedAt = field.NewInt64(tableName, "created_at")
	_winUserLevelReward.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winUserLevelReward.fillFieldMap()

	return _winUserLevelReward
}

// winUserLevelReward vip奖金表
type winUserLevelReward struct {
	winUserLevelRewardDo

	ALL         field.Asterisk
	ID          field.Int64
	UID         field.Int64  // UID
	Username    field.String // 用户名
	Coin        field.Field  // 彩金金额
	CoinBefore  field.Field  // 即时金额
	Category    field.Int64  // 福利类型:0-晋级彩金 1-生日礼金 2-周彩金3 -月彩金
	ReceiveAt   field.Int64  // 领取时间
	UserLevelID field.Int64  // 用户等级编号
	FlowClaim   field.Int64  // 流水倍数
	Codes       field.Field  // 所需打码量
	EndedAt     field.Int64  // 领取结束时间
	Status      field.Int64  // 状态:0-待领取 1-已领取 2-已过期
	CreatedAt   field.Int64
	UpdatedAt   field.Int64

	fieldMap map[string]field.Expr
}

func (w winUserLevelReward) Table(newTableName string) *winUserLevelReward {
	w.winUserLevelRewardDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winUserLevelReward) As(alias string) *winUserLevelReward {
	w.winUserLevelRewardDo.DO = *(w.winUserLevelRewardDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winUserLevelReward) updateTableName(table string) *winUserLevelReward {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.Coin = field.NewField(table, "coin")
	w.CoinBefore = field.NewField(table, "coin_before")
	w.Category = field.NewInt64(table, "category")
	w.ReceiveAt = field.NewInt64(table, "receive_at")
	w.UserLevelID = field.NewInt64(table, "user_level_id")
	w.FlowClaim = field.NewInt64(table, "flow_claim")
	w.Codes = field.NewField(table, "codes")
	w.EndedAt = field.NewInt64(table, "ended_at")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winUserLevelReward) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winUserLevelReward) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 14)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["category"] = w.Category
	w.fieldMap["receive_at"] = w.ReceiveAt
	w.fieldMap["user_level_id"] = w.UserLevelID
	w.fieldMap["flow_claim"] = w.FlowClaim
	w.fieldMap["codes"] = w.Codes
	w.fieldMap["ended_at"] = w.EndedAt
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winUserLevelReward) clone(db *gorm.DB) winUserLevelReward {
	w.winUserLevelRewardDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winUserLevelReward) replaceDB(db *gorm.DB) winUserLevelReward {
	w.winUserLevelRewardDo.ReplaceDB(db)
	return w
}

type winUserLevelRewardDo struct{ gen.DO }

type IWinUserLevelRewardDo interface {
	gen.SubQuery
	Debug() IWinUserLevelRewardDo
	WithContext(ctx context.Context) IWinUserLevelRewardDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinUserLevelRewardDo
	WriteDB() IWinUserLevelRewardDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinUserLevelRewardDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinUserLevelRewardDo
	Not(conds ...gen.Condition) IWinUserLevelRewardDo
	Or(conds ...gen.Condition) IWinUserLevelRewardDo
	Select(conds ...field.Expr) IWinUserLevelRewardDo
	Where(conds ...gen.Condition) IWinUserLevelRewardDo
	Order(conds ...field.Expr) IWinUserLevelRewardDo
	Distinct(cols ...field.Expr) IWinUserLevelRewardDo
	Omit(cols ...field.Expr) IWinUserLevelRewardDo
	Join(table schema.Tabler, on ...field.Expr) IWinUserLevelRewardDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelRewardDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelRewardDo
	Group(cols ...field.Expr) IWinUserLevelRewardDo
	Having(conds ...gen.Condition) IWinUserLevelRewardDo
	Limit(limit int) IWinUserLevelRewardDo
	Offset(offset int) IWinUserLevelRewardDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserLevelRewardDo
	Unscoped() IWinUserLevelRewardDo
	Create(values ...*model.WinUserLevelReward) error
	CreateInBatches(values []*model.WinUserLevelReward, batchSize int) error
	Save(values ...*model.WinUserLevelReward) error
	First() (*model.WinUserLevelReward, error)
	Take() (*model.WinUserLevelReward, error)
	Last() (*model.WinUserLevelReward, error)
	Find() ([]*model.WinUserLevelReward, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserLevelReward, err error)
	FindInBatches(result *[]*model.WinUserLevelReward, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinUserLevelReward) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinUserLevelRewardDo
	Assign(attrs ...field.AssignExpr) IWinUserLevelRewardDo
	Joins(fields ...field.RelationField) IWinUserLevelRewardDo
	Preload(fields ...field.RelationField) IWinUserLevelRewardDo
	FirstOrInit() (*model.WinUserLevelReward, error)
	FirstOrCreate() (*model.WinUserLevelReward, error)
	FindByPage(offset int, limit int) (result []*model.WinUserLevelReward, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinUserLevelRewardDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winUserLevelRewardDo) Debug() IWinUserLevelRewardDo {
	return w.withDO(w.DO.Debug())
}

func (w winUserLevelRewardDo) WithContext(ctx context.Context) IWinUserLevelRewardDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winUserLevelRewardDo) ReadDB() IWinUserLevelRewardDo {
	return w.Clauses(dbresolver.Read)
}

func (w winUserLevelRewardDo) WriteDB() IWinUserLevelRewardDo {
	return w.Clauses(dbresolver.Write)
}

func (w winUserLevelRewardDo) Session(config *gorm.Session) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Session(config))
}

func (w winUserLevelRewardDo) Clauses(conds ...clause.Expression) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winUserLevelRewardDo) Returning(value interface{}, columns ...string) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winUserLevelRewardDo) Not(conds ...gen.Condition) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winUserLevelRewardDo) Or(conds ...gen.Condition) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winUserLevelRewardDo) Select(conds ...field.Expr) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winUserLevelRewardDo) Where(conds ...gen.Condition) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winUserLevelRewardDo) Order(conds ...field.Expr) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winUserLevelRewardDo) Distinct(cols ...field.Expr) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winUserLevelRewardDo) Omit(cols ...field.Expr) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winUserLevelRewardDo) Join(table schema.Tabler, on ...field.Expr) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winUserLevelRewardDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelRewardDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winUserLevelRewardDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelRewardDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winUserLevelRewardDo) Group(cols ...field.Expr) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winUserLevelRewardDo) Having(conds ...gen.Condition) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winUserLevelRewardDo) Limit(limit int) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winUserLevelRewardDo) Offset(offset int) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winUserLevelRewardDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winUserLevelRewardDo) Unscoped() IWinUserLevelRewardDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winUserLevelRewardDo) Create(values ...*model.WinUserLevelReward) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winUserLevelRewardDo) CreateInBatches(values []*model.WinUserLevelReward, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winUserLevelRewardDo) Save(values ...*model.WinUserLevelReward) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winUserLevelRewardDo) First() (*model.WinUserLevelReward, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelReward), nil
	}
}

func (w winUserLevelRewardDo) Take() (*model.WinUserLevelReward, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelReward), nil
	}
}

func (w winUserLevelRewardDo) Last() (*model.WinUserLevelReward, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelReward), nil
	}
}

func (w winUserLevelRewardDo) Find() ([]*model.WinUserLevelReward, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinUserLevelReward), err
}

func (w winUserLevelRewardDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserLevelReward, err error) {
	buf := make([]*model.WinUserLevelReward, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winUserLevelRewardDo) FindInBatches(result *[]*model.WinUserLevelReward, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winUserLevelRewardDo) Attrs(attrs ...field.AssignExpr) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winUserLevelRewardDo) Assign(attrs ...field.AssignExpr) IWinUserLevelRewardDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winUserLevelRewardDo) Joins(fields ...field.RelationField) IWinUserLevelRewardDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winUserLevelRewardDo) Preload(fields ...field.RelationField) IWinUserLevelRewardDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winUserLevelRewardDo) FirstOrInit() (*model.WinUserLevelReward, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelReward), nil
	}
}

func (w winUserLevelRewardDo) FirstOrCreate() (*model.WinUserLevelReward, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelReward), nil
	}
}

func (w winUserLevelRewardDo) FindByPage(offset int, limit int) (result []*model.WinUserLevelReward, count int64, err error) {
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

func (w winUserLevelRewardDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winUserLevelRewardDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winUserLevelRewardDo) Delete(models ...*model.WinUserLevelReward) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winUserLevelRewardDo) withDO(do gen.Dao) *winUserLevelRewardDo {
	w.DO = *do.(*gen.DO)
	return w
}
