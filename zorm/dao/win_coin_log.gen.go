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

func newWinCoinLog(db *gorm.DB, opts ...gen.DOOption) winCoinLog {
	_winCoinLog := winCoinLog{}

	_winCoinLog.winCoinLogDo.UseDB(db, opts...)
	_winCoinLog.winCoinLogDo.UseModel(&model.WinCoinLog{})

	tableName := _winCoinLog.winCoinLogDo.TableName()
	_winCoinLog.ALL = field.NewAsterisk(tableName)
	_winCoinLog.ID = field.NewInt64(tableName, "id")
	_winCoinLog.UID = field.NewInt64(tableName, "uid")
	_winCoinLog.Username = field.NewString(tableName, "username")
	_winCoinLog.MerchantID = field.NewInt64(tableName, "merchant_id")
	_winCoinLog.WalletID = field.NewInt64(tableName, "wallet_id")
	_winCoinLog.Currency = field.NewInt64(tableName, "currency")
	_winCoinLog.Category = field.NewInt64(tableName, "category")
	_winCoinLog.Code = field.NewString(tableName, "code")
	_winCoinLog.PlatName = field.NewString(tableName, "plat_name")
	_winCoinLog.PlatNickName = field.NewString(tableName, "plat_nick_name")
	_winCoinLog.ReferID = field.NewInt64(tableName, "refer_id")
	_winCoinLog.OrderID = field.NewString(tableName, "order_id")
	_winCoinLog.Coin = field.NewField(tableName, "coin")
	_winCoinLog.CoinReal = field.NewField(tableName, "coin_real")
	_winCoinLog.PlatID = field.NewInt64(tableName, "plat_id")
	_winCoinLog.OutIn = field.NewInt64(tableName, "out_in")
	_winCoinLog.GameID = field.NewInt64(tableName, "game_id")
	_winCoinLog.CoinBefore = field.NewField(tableName, "coin_before")
	_winCoinLog.CoinAfter = field.NewField(tableName, "coin_after")
	_winCoinLog.Status = field.NewInt64(tableName, "status")
	_winCoinLog.CreatedAt = field.NewInt64(tableName, "created_at")
	_winCoinLog.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winCoinLog.fillFieldMap()

	return _winCoinLog
}

type winCoinLog struct {
	winCoinLogDo

	ALL          field.Asterisk
	ID           field.Int64
	UID          field.Int64  // UID
	Username     field.String // 用户名
	MerchantID   field.Int64  // 商户id
	WalletID     field.Int64  // 钱包id
	Currency     field.Int64  // 币种
	Category     field.Int64  // 类型:1-存款 2-提款 3-投注 4-派彩 5-返水 6-佣金 7-活动(奖励) 8-系统调账 9-退款 10-佣金钱包转主账户余额 11-小费,12-提款退款,13-系统调账转出,14-系统调账转入,15-余额带出入游戏,16-游戏调整余额
	Code         field.String // 支付通道编码
	PlatName     field.String // 平台名称
	PlatNickName field.String // 平台自定义名称
	ReferID      field.Int64  // 关联ID
	OrderID      field.String // 订单id
	Coin         field.Field  // 金额
	CoinReal     field.Field  // 实际金额
	PlatID       field.Int64  // 游戏平台ID
	OutIn        field.Int64  // 收支类型:0-支出 1-收入
	GameID       field.Int64  // 三方游戏ID
	CoinBefore   field.Field  // 前金额
	CoinAfter    field.Field  // 帐变后金额
	Status       field.Int64  // 状态:0-处理中 1-成功 2-失败
	CreatedAt    field.Int64
	UpdatedAt    field.Int64

	fieldMap map[string]field.Expr
}

func (w winCoinLog) Table(newTableName string) *winCoinLog {
	w.winCoinLogDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCoinLog) As(alias string) *winCoinLog {
	w.winCoinLogDo.DO = *(w.winCoinLogDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCoinLog) updateTableName(table string) *winCoinLog {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.MerchantID = field.NewInt64(table, "merchant_id")
	w.WalletID = field.NewInt64(table, "wallet_id")
	w.Currency = field.NewInt64(table, "currency")
	w.Category = field.NewInt64(table, "category")
	w.Code = field.NewString(table, "code")
	w.PlatName = field.NewString(table, "plat_name")
	w.PlatNickName = field.NewString(table, "plat_nick_name")
	w.ReferID = field.NewInt64(table, "refer_id")
	w.OrderID = field.NewString(table, "order_id")
	w.Coin = field.NewField(table, "coin")
	w.CoinReal = field.NewField(table, "coin_real")
	w.PlatID = field.NewInt64(table, "plat_id")
	w.OutIn = field.NewInt64(table, "out_in")
	w.GameID = field.NewInt64(table, "game_id")
	w.CoinBefore = field.NewField(table, "coin_before")
	w.CoinAfter = field.NewField(table, "coin_after")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winCoinLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCoinLog) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 22)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["merchant_id"] = w.MerchantID
	w.fieldMap["wallet_id"] = w.WalletID
	w.fieldMap["currency"] = w.Currency
	w.fieldMap["category"] = w.Category
	w.fieldMap["code"] = w.Code
	w.fieldMap["plat_name"] = w.PlatName
	w.fieldMap["plat_nick_name"] = w.PlatNickName
	w.fieldMap["refer_id"] = w.ReferID
	w.fieldMap["order_id"] = w.OrderID
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["coin_real"] = w.CoinReal
	w.fieldMap["plat_id"] = w.PlatID
	w.fieldMap["out_in"] = w.OutIn
	w.fieldMap["game_id"] = w.GameID
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["coin_after"] = w.CoinAfter
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winCoinLog) clone(db *gorm.DB) winCoinLog {
	w.winCoinLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCoinLog) replaceDB(db *gorm.DB) winCoinLog {
	w.winCoinLogDo.ReplaceDB(db)
	return w
}

type winCoinLogDo struct{ gen.DO }

type IWinCoinLogDo interface {
	gen.SubQuery
	Debug() IWinCoinLogDo
	WithContext(ctx context.Context) IWinCoinLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinCoinLogDo
	WriteDB() IWinCoinLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinCoinLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinCoinLogDo
	Not(conds ...gen.Condition) IWinCoinLogDo
	Or(conds ...gen.Condition) IWinCoinLogDo
	Select(conds ...field.Expr) IWinCoinLogDo
	Where(conds ...gen.Condition) IWinCoinLogDo
	Order(conds ...field.Expr) IWinCoinLogDo
	Distinct(cols ...field.Expr) IWinCoinLogDo
	Omit(cols ...field.Expr) IWinCoinLogDo
	Join(table schema.Tabler, on ...field.Expr) IWinCoinLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinLogDo
	Group(cols ...field.Expr) IWinCoinLogDo
	Having(conds ...gen.Condition) IWinCoinLogDo
	Limit(limit int) IWinCoinLogDo
	Offset(offset int) IWinCoinLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinLogDo
	Unscoped() IWinCoinLogDo
	Create(values ...*model.WinCoinLog) error
	CreateInBatches(values []*model.WinCoinLog, batchSize int) error
	Save(values ...*model.WinCoinLog) error
	First() (*model.WinCoinLog, error)
	Take() (*model.WinCoinLog, error)
	Last() (*model.WinCoinLog, error)
	Find() ([]*model.WinCoinLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinLog, err error)
	FindInBatches(result *[]*model.WinCoinLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinCoinLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinCoinLogDo
	Assign(attrs ...field.AssignExpr) IWinCoinLogDo
	Joins(fields ...field.RelationField) IWinCoinLogDo
	Preload(fields ...field.RelationField) IWinCoinLogDo
	FirstOrInit() (*model.WinCoinLog, error)
	FirstOrCreate() (*model.WinCoinLog, error)
	FindByPage(offset int, limit int) (result []*model.WinCoinLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinCoinLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winCoinLogDo) Debug() IWinCoinLogDo {
	return w.withDO(w.DO.Debug())
}

func (w winCoinLogDo) WithContext(ctx context.Context) IWinCoinLogDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCoinLogDo) ReadDB() IWinCoinLogDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCoinLogDo) WriteDB() IWinCoinLogDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCoinLogDo) Session(config *gorm.Session) IWinCoinLogDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCoinLogDo) Clauses(conds ...clause.Expression) IWinCoinLogDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCoinLogDo) Returning(value interface{}, columns ...string) IWinCoinLogDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCoinLogDo) Not(conds ...gen.Condition) IWinCoinLogDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCoinLogDo) Or(conds ...gen.Condition) IWinCoinLogDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCoinLogDo) Select(conds ...field.Expr) IWinCoinLogDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCoinLogDo) Where(conds ...gen.Condition) IWinCoinLogDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCoinLogDo) Order(conds ...field.Expr) IWinCoinLogDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCoinLogDo) Distinct(cols ...field.Expr) IWinCoinLogDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCoinLogDo) Omit(cols ...field.Expr) IWinCoinLogDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCoinLogDo) Join(table schema.Tabler, on ...field.Expr) IWinCoinLogDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCoinLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinLogDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCoinLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinLogDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCoinLogDo) Group(cols ...field.Expr) IWinCoinLogDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCoinLogDo) Having(conds ...gen.Condition) IWinCoinLogDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCoinLogDo) Limit(limit int) IWinCoinLogDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCoinLogDo) Offset(offset int) IWinCoinLogDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCoinLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinLogDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCoinLogDo) Unscoped() IWinCoinLogDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCoinLogDo) Create(values ...*model.WinCoinLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCoinLogDo) CreateInBatches(values []*model.WinCoinLog, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCoinLogDo) Save(values ...*model.WinCoinLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCoinLogDo) First() (*model.WinCoinLog, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinLog), nil
	}
}

func (w winCoinLogDo) Take() (*model.WinCoinLog, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinLog), nil
	}
}

func (w winCoinLogDo) Last() (*model.WinCoinLog, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinLog), nil
	}
}

func (w winCoinLogDo) Find() ([]*model.WinCoinLog, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCoinLog), err
}

func (w winCoinLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinLog, err error) {
	buf := make([]*model.WinCoinLog, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCoinLogDo) FindInBatches(result *[]*model.WinCoinLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCoinLogDo) Attrs(attrs ...field.AssignExpr) IWinCoinLogDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCoinLogDo) Assign(attrs ...field.AssignExpr) IWinCoinLogDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCoinLogDo) Joins(fields ...field.RelationField) IWinCoinLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCoinLogDo) Preload(fields ...field.RelationField) IWinCoinLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCoinLogDo) FirstOrInit() (*model.WinCoinLog, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinLog), nil
	}
}

func (w winCoinLogDo) FirstOrCreate() (*model.WinCoinLog, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinLog), nil
	}
}

func (w winCoinLogDo) FindByPage(offset int, limit int) (result []*model.WinCoinLog, count int64, err error) {
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

func (w winCoinLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCoinLogDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCoinLogDo) Delete(models ...*model.WinCoinLog) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCoinLogDo) withDO(do gen.Dao) *winCoinLogDo {
	w.DO = *do.(*gen.DO)
	return w
}