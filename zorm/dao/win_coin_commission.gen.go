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

func newWinCoinCommission(db *gorm.DB, opts ...gen.DOOption) winCoinCommission {
	_winCoinCommission := winCoinCommission{}

	_winCoinCommission.winCoinCommissionDo.UseDB(db, opts...)
	_winCoinCommission.winCoinCommissionDo.UseModel(&model.WinCoinCommission{})

	tableName := _winCoinCommission.winCoinCommissionDo.TableName()
	_winCoinCommission.ALL = field.NewAsterisk(tableName)
	_winCoinCommission.ID = field.NewInt64(tableName, "id")
	_winCoinCommission.UID = field.NewInt64(tableName, "uid")
	_winCoinCommission.Username = field.NewString(tableName, "username")
	_winCoinCommission.AgentLevel = field.NewInt64(tableName, "agent_level")
	_winCoinCommission.Riqi = field.NewInt64(tableName, "riqi")
	_winCoinCommission.Coin = field.NewField(tableName, "coin")
	_winCoinCommission.SubUID = field.NewInt64(tableName, "sub_uid")
	_winCoinCommission.SubUsername = field.NewString(tableName, "sub_username")
	_winCoinCommission.SubBetTrunover = field.NewField(tableName, "sub_bet_trunover")
	_winCoinCommission.SubBetProfit = field.NewField(tableName, "sub_bet_profit")
	_winCoinCommission.Rate = field.NewField(tableName, "rate")
	_winCoinCommission.CoinBefore = field.NewField(tableName, "coin_before")
	_winCoinCommission.Status = field.NewInt64(tableName, "status")
	_winCoinCommission.CreatedAt = field.NewInt64(tableName, "created_at")
	_winCoinCommission.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winCoinCommission.Month = field.NewInt64(tableName, "month")
	_winCoinCommission.Year = field.NewInt64(tableName, "year")

	_winCoinCommission.fillFieldMap()

	return _winCoinCommission
}

type winCoinCommission struct {
	winCoinCommissionDo

	ALL            field.Asterisk
	ID             field.Int64
	UID            field.Int64  // 代理UID
	Username       field.String // 代理用户名
	AgentLevel     field.Int64  // 代理层级
	Riqi           field.Int64  // 佣金时间
	Coin           field.Field  // 佣金金额
	SubUID         field.Int64  // 下级UID
	SubUsername    field.String // 下级用户名
	SubBetTrunover field.Field  // 下级流水总额
	SubBetProfit   field.Field  // 下级盈亏
	Rate           field.Field  // 佣金比例
	CoinBefore     field.Field  // 即时余额
	Status         field.Int64  // 状态:0-未发放 1-已发放
	CreatedAt      field.Int64  // 创建时间
	UpdatedAt      field.Int64  // 更新时间
	Month          field.Int64  // 佣金月份
	Year           field.Int64  // 佣金年份

	fieldMap map[string]field.Expr
}

func (w winCoinCommission) Table(newTableName string) *winCoinCommission {
	w.winCoinCommissionDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCoinCommission) As(alias string) *winCoinCommission {
	w.winCoinCommissionDo.DO = *(w.winCoinCommissionDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCoinCommission) updateTableName(table string) *winCoinCommission {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.AgentLevel = field.NewInt64(table, "agent_level")
	w.Riqi = field.NewInt64(table, "riqi")
	w.Coin = field.NewField(table, "coin")
	w.SubUID = field.NewInt64(table, "sub_uid")
	w.SubUsername = field.NewString(table, "sub_username")
	w.SubBetTrunover = field.NewField(table, "sub_bet_trunover")
	w.SubBetProfit = field.NewField(table, "sub_bet_profit")
	w.Rate = field.NewField(table, "rate")
	w.CoinBefore = field.NewField(table, "coin_before")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.Month = field.NewInt64(table, "month")
	w.Year = field.NewInt64(table, "year")

	w.fillFieldMap()

	return w
}

func (w *winCoinCommission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCoinCommission) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 17)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["agent_level"] = w.AgentLevel
	w.fieldMap["riqi"] = w.Riqi
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["sub_uid"] = w.SubUID
	w.fieldMap["sub_username"] = w.SubUsername
	w.fieldMap["sub_bet_trunover"] = w.SubBetTrunover
	w.fieldMap["sub_bet_profit"] = w.SubBetProfit
	w.fieldMap["rate"] = w.Rate
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["month"] = w.Month
	w.fieldMap["year"] = w.Year
}

func (w winCoinCommission) clone(db *gorm.DB) winCoinCommission {
	w.winCoinCommissionDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCoinCommission) replaceDB(db *gorm.DB) winCoinCommission {
	w.winCoinCommissionDo.ReplaceDB(db)
	return w
}

type winCoinCommissionDo struct{ gen.DO }

type IWinCoinCommissionDo interface {
	gen.SubQuery
	Debug() IWinCoinCommissionDo
	WithContext(ctx context.Context) IWinCoinCommissionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinCoinCommissionDo
	WriteDB() IWinCoinCommissionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinCoinCommissionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinCoinCommissionDo
	Not(conds ...gen.Condition) IWinCoinCommissionDo
	Or(conds ...gen.Condition) IWinCoinCommissionDo
	Select(conds ...field.Expr) IWinCoinCommissionDo
	Where(conds ...gen.Condition) IWinCoinCommissionDo
	Order(conds ...field.Expr) IWinCoinCommissionDo
	Distinct(cols ...field.Expr) IWinCoinCommissionDo
	Omit(cols ...field.Expr) IWinCoinCommissionDo
	Join(table schema.Tabler, on ...field.Expr) IWinCoinCommissionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinCommissionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinCommissionDo
	Group(cols ...field.Expr) IWinCoinCommissionDo
	Having(conds ...gen.Condition) IWinCoinCommissionDo
	Limit(limit int) IWinCoinCommissionDo
	Offset(offset int) IWinCoinCommissionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinCommissionDo
	Unscoped() IWinCoinCommissionDo
	Create(values ...*model.WinCoinCommission) error
	CreateInBatches(values []*model.WinCoinCommission, batchSize int) error
	Save(values ...*model.WinCoinCommission) error
	First() (*model.WinCoinCommission, error)
	Take() (*model.WinCoinCommission, error)
	Last() (*model.WinCoinCommission, error)
	Find() ([]*model.WinCoinCommission, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinCommission, err error)
	FindInBatches(result *[]*model.WinCoinCommission, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinCoinCommission) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinCoinCommissionDo
	Assign(attrs ...field.AssignExpr) IWinCoinCommissionDo
	Joins(fields ...field.RelationField) IWinCoinCommissionDo
	Preload(fields ...field.RelationField) IWinCoinCommissionDo
	FirstOrInit() (*model.WinCoinCommission, error)
	FirstOrCreate() (*model.WinCoinCommission, error)
	FindByPage(offset int, limit int) (result []*model.WinCoinCommission, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinCoinCommissionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winCoinCommissionDo) Debug() IWinCoinCommissionDo {
	return w.withDO(w.DO.Debug())
}

func (w winCoinCommissionDo) WithContext(ctx context.Context) IWinCoinCommissionDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCoinCommissionDo) ReadDB() IWinCoinCommissionDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCoinCommissionDo) WriteDB() IWinCoinCommissionDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCoinCommissionDo) Session(config *gorm.Session) IWinCoinCommissionDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCoinCommissionDo) Clauses(conds ...clause.Expression) IWinCoinCommissionDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCoinCommissionDo) Returning(value interface{}, columns ...string) IWinCoinCommissionDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCoinCommissionDo) Not(conds ...gen.Condition) IWinCoinCommissionDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCoinCommissionDo) Or(conds ...gen.Condition) IWinCoinCommissionDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCoinCommissionDo) Select(conds ...field.Expr) IWinCoinCommissionDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCoinCommissionDo) Where(conds ...gen.Condition) IWinCoinCommissionDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCoinCommissionDo) Order(conds ...field.Expr) IWinCoinCommissionDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCoinCommissionDo) Distinct(cols ...field.Expr) IWinCoinCommissionDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCoinCommissionDo) Omit(cols ...field.Expr) IWinCoinCommissionDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCoinCommissionDo) Join(table schema.Tabler, on ...field.Expr) IWinCoinCommissionDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCoinCommissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinCommissionDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCoinCommissionDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinCommissionDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCoinCommissionDo) Group(cols ...field.Expr) IWinCoinCommissionDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCoinCommissionDo) Having(conds ...gen.Condition) IWinCoinCommissionDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCoinCommissionDo) Limit(limit int) IWinCoinCommissionDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCoinCommissionDo) Offset(offset int) IWinCoinCommissionDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCoinCommissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinCommissionDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCoinCommissionDo) Unscoped() IWinCoinCommissionDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCoinCommissionDo) Create(values ...*model.WinCoinCommission) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCoinCommissionDo) CreateInBatches(values []*model.WinCoinCommission, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCoinCommissionDo) Save(values ...*model.WinCoinCommission) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCoinCommissionDo) First() (*model.WinCoinCommission, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommission), nil
	}
}

func (w winCoinCommissionDo) Take() (*model.WinCoinCommission, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommission), nil
	}
}

func (w winCoinCommissionDo) Last() (*model.WinCoinCommission, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommission), nil
	}
}

func (w winCoinCommissionDo) Find() ([]*model.WinCoinCommission, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCoinCommission), err
}

func (w winCoinCommissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinCommission, err error) {
	buf := make([]*model.WinCoinCommission, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCoinCommissionDo) FindInBatches(result *[]*model.WinCoinCommission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCoinCommissionDo) Attrs(attrs ...field.AssignExpr) IWinCoinCommissionDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCoinCommissionDo) Assign(attrs ...field.AssignExpr) IWinCoinCommissionDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCoinCommissionDo) Joins(fields ...field.RelationField) IWinCoinCommissionDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCoinCommissionDo) Preload(fields ...field.RelationField) IWinCoinCommissionDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCoinCommissionDo) FirstOrInit() (*model.WinCoinCommission, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommission), nil
	}
}

func (w winCoinCommissionDo) FirstOrCreate() (*model.WinCoinCommission, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommission), nil
	}
}

func (w winCoinCommissionDo) FindByPage(offset int, limit int) (result []*model.WinCoinCommission, count int64, err error) {
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

func (w winCoinCommissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCoinCommissionDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCoinCommissionDo) Delete(models ...*model.WinCoinCommission) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCoinCommissionDo) withDO(do gen.Dao) *winCoinCommissionDo {
	w.DO = *do.(*gen.DO)
	return w
}
