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

func newWinBetslips(db *gorm.DB, opts ...gen.DOOption) winBetslips {
	_winBetslips := winBetslips{}

	_winBetslips.winBetslipsDo.UseDB(db, opts...)
	_winBetslips.winBetslipsDo.UseModel(&model.WinBetslips{})

	tableName := _winBetslips.winBetslipsDo.TableName()
	_winBetslips.ALL = field.NewAsterisk(tableName)
	_winBetslips.ID = field.NewInt64(tableName, "id")
	_winBetslips.RoundID = field.NewString(tableName, "round_id")
	_winBetslips.TransactionID = field.NewString(tableName, "transaction_id")
	_winBetslips.XbStatus = field.NewInt64(tableName, "xb_status")
	_winBetslips.XbUID = field.NewInt64(tableName, "xb_uid")
	_winBetslips.XbUsername = field.NewString(tableName, "xb_username")
	_winBetslips.MerchantID = field.NewInt64(tableName, "merchant_id")
	_winBetslips.XbProfit = field.NewField(tableName, "xb_profit")
	_winBetslips.Stake = field.NewField(tableName, "stake")
	_winBetslips.ValidStake = field.NewField(tableName, "valid_stake")
	_winBetslips.Payout = field.NewField(tableName, "payout")
	_winBetslips.CoinRefund = field.NewField(tableName, "coin_refund")
	_winBetslips.CoinBefore = field.NewField(tableName, "coin_before")
	_winBetslips.GameProviderSubtypeID = field.NewInt64(tableName, "game_provider_subtype_id")
	_winBetslips.GameListID = field.NewInt64(tableName, "game_list_id")
	_winBetslips.GamePagcorID = field.NewInt64(tableName, "game_pagcor_id")
	_winBetslips.GameProviderID = field.NewInt64(tableName, "game_provider_id")
	_winBetslips.AmountType = field.NewInt64(tableName, "amount_type")
	_winBetslips.DtStarted = field.NewInt64(tableName, "dt_started")
	_winBetslips.DtCompleted = field.NewInt64(tableName, "dt_completed")
	_winBetslips.WinTransactionID = field.NewString(tableName, "win_transaction_id")
	_winBetslips.CreateTimeStr = field.NewString(tableName, "create_time_str")
	_winBetslips.CreatedAt = field.NewInt64(tableName, "created_at")
	_winBetslips.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winBetslips.GameTypeID = field.NewInt64(tableName, "game_type_id")
	_winBetslips.IsCounted = field.NewInt64(tableName, "is_counted")

	_winBetslips.fillFieldMap()

	return _winBetslips
}

type winBetslips struct {
	winBetslipsDo

	ALL                   field.Asterisk
	ID                    field.Int64  // 主键
	RoundID               field.String // 回合id
	TransactionID         field.String // 注单号 对应三方拉单transaction_id
	XbStatus              field.Int64  // 注单状态 1:待开彩  2:完成  3: 退款
	XbUID                 field.Int64  // 对应user表id
	XbUsername            field.String // 对应user表username
	MerchantID            field.Int64  // 商户id
	XbProfit              field.Field  // 盈亏金额
	Stake                 field.Field  // 投注
	ValidStake            field.Field  // 有效投注金额
	Payout                field.Field  // 派彩
	CoinRefund            field.Field  // 退款金额
	CoinBefore            field.Field  // 投注前金额
	GameProviderSubtypeID field.Int64  // 游戏id对应game_provider_subtype表id
	GameListID            field.Int64  // 游戏id对应game_list表id
	GamePagcorID          field.Int64  // pagcor分组id
	GameProviderID        field.Int64  // 游戏供应商id
	AmountType            field.Int64  // 投注方式 1:现金，2:奖金 3:免费旋转 4:活动免费旋转
	DtStarted             field.Int64  // 游戏开始时间
	DtCompleted           field.Int64  // 游戏结束时间
	WinTransactionID      field.String // 开奖交易单号
	CreateTimeStr         field.String // 投注时间
	CreatedAt             field.Int64
	UpdatedAt             field.Int64
	GameTypeID            field.Int64 // 游戏分组id
	IsCounted             field.Int64 // 是否统计过：1=否，3=是

	fieldMap map[string]field.Expr
}

func (w winBetslips) Table(newTableName string) *winBetslips {
	w.winBetslipsDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winBetslips) As(alias string) *winBetslips {
	w.winBetslipsDo.DO = *(w.winBetslipsDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winBetslips) updateTableName(table string) *winBetslips {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.RoundID = field.NewString(table, "round_id")
	w.TransactionID = field.NewString(table, "transaction_id")
	w.XbStatus = field.NewInt64(table, "xb_status")
	w.XbUID = field.NewInt64(table, "xb_uid")
	w.XbUsername = field.NewString(table, "xb_username")
	w.MerchantID = field.NewInt64(table, "merchant_id")
	w.XbProfit = field.NewField(table, "xb_profit")
	w.Stake = field.NewField(table, "stake")
	w.ValidStake = field.NewField(table, "valid_stake")
	w.Payout = field.NewField(table, "payout")
	w.CoinRefund = field.NewField(table, "coin_refund")
	w.CoinBefore = field.NewField(table, "coin_before")
	w.GameProviderSubtypeID = field.NewInt64(table, "game_provider_subtype_id")
	w.GameListID = field.NewInt64(table, "game_list_id")
	w.GamePagcorID = field.NewInt64(table, "game_pagcor_id")
	w.GameProviderID = field.NewInt64(table, "game_provider_id")
	w.AmountType = field.NewInt64(table, "amount_type")
	w.DtStarted = field.NewInt64(table, "dt_started")
	w.DtCompleted = field.NewInt64(table, "dt_completed")
	w.WinTransactionID = field.NewString(table, "win_transaction_id")
	w.CreateTimeStr = field.NewString(table, "create_time_str")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.GameTypeID = field.NewInt64(table, "game_type_id")
	w.IsCounted = field.NewInt64(table, "is_counted")

	w.fillFieldMap()

	return w
}

func (w *winBetslips) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winBetslips) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 26)
	w.fieldMap["id"] = w.ID
	w.fieldMap["round_id"] = w.RoundID
	w.fieldMap["transaction_id"] = w.TransactionID
	w.fieldMap["xb_status"] = w.XbStatus
	w.fieldMap["xb_uid"] = w.XbUID
	w.fieldMap["xb_username"] = w.XbUsername
	w.fieldMap["merchant_id"] = w.MerchantID
	w.fieldMap["xb_profit"] = w.XbProfit
	w.fieldMap["stake"] = w.Stake
	w.fieldMap["valid_stake"] = w.ValidStake
	w.fieldMap["payout"] = w.Payout
	w.fieldMap["coin_refund"] = w.CoinRefund
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["game_provider_subtype_id"] = w.GameProviderSubtypeID
	w.fieldMap["game_list_id"] = w.GameListID
	w.fieldMap["game_pagcor_id"] = w.GamePagcorID
	w.fieldMap["game_provider_id"] = w.GameProviderID
	w.fieldMap["amount_type"] = w.AmountType
	w.fieldMap["dt_started"] = w.DtStarted
	w.fieldMap["dt_completed"] = w.DtCompleted
	w.fieldMap["win_transaction_id"] = w.WinTransactionID
	w.fieldMap["create_time_str"] = w.CreateTimeStr
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["game_type_id"] = w.GameTypeID
	w.fieldMap["is_counted"] = w.IsCounted
}

func (w winBetslips) clone(db *gorm.DB) winBetslips {
	w.winBetslipsDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winBetslips) replaceDB(db *gorm.DB) winBetslips {
	w.winBetslipsDo.ReplaceDB(db)
	return w
}

type winBetslipsDo struct{ gen.DO }

type IWinBetslipsDo interface {
	gen.SubQuery
	Debug() IWinBetslipsDo
	WithContext(ctx context.Context) IWinBetslipsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinBetslipsDo
	WriteDB() IWinBetslipsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinBetslipsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinBetslipsDo
	Not(conds ...gen.Condition) IWinBetslipsDo
	Or(conds ...gen.Condition) IWinBetslipsDo
	Select(conds ...field.Expr) IWinBetslipsDo
	Where(conds ...gen.Condition) IWinBetslipsDo
	Order(conds ...field.Expr) IWinBetslipsDo
	Distinct(cols ...field.Expr) IWinBetslipsDo
	Omit(cols ...field.Expr) IWinBetslipsDo
	Join(table schema.Tabler, on ...field.Expr) IWinBetslipsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDo
	Group(cols ...field.Expr) IWinBetslipsDo
	Having(conds ...gen.Condition) IWinBetslipsDo
	Limit(limit int) IWinBetslipsDo
	Offset(offset int) IWinBetslipsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetslipsDo
	Unscoped() IWinBetslipsDo
	Create(values ...*model.WinBetslips) error
	CreateInBatches(values []*model.WinBetslips, batchSize int) error
	Save(values ...*model.WinBetslips) error
	First() (*model.WinBetslips, error)
	Take() (*model.WinBetslips, error)
	Last() (*model.WinBetslips, error)
	Find() ([]*model.WinBetslips, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetslips, err error)
	FindInBatches(result *[]*model.WinBetslips, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinBetslips) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinBetslipsDo
	Assign(attrs ...field.AssignExpr) IWinBetslipsDo
	Joins(fields ...field.RelationField) IWinBetslipsDo
	Preload(fields ...field.RelationField) IWinBetslipsDo
	FirstOrInit() (*model.WinBetslips, error)
	FirstOrCreate() (*model.WinBetslips, error)
	FindByPage(offset int, limit int) (result []*model.WinBetslips, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinBetslipsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winBetslipsDo) Debug() IWinBetslipsDo {
	return w.withDO(w.DO.Debug())
}

func (w winBetslipsDo) WithContext(ctx context.Context) IWinBetslipsDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winBetslipsDo) ReadDB() IWinBetslipsDo {
	return w.Clauses(dbresolver.Read)
}

func (w winBetslipsDo) WriteDB() IWinBetslipsDo {
	return w.Clauses(dbresolver.Write)
}

func (w winBetslipsDo) Session(config *gorm.Session) IWinBetslipsDo {
	return w.withDO(w.DO.Session(config))
}

func (w winBetslipsDo) Clauses(conds ...clause.Expression) IWinBetslipsDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winBetslipsDo) Returning(value interface{}, columns ...string) IWinBetslipsDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winBetslipsDo) Not(conds ...gen.Condition) IWinBetslipsDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winBetslipsDo) Or(conds ...gen.Condition) IWinBetslipsDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winBetslipsDo) Select(conds ...field.Expr) IWinBetslipsDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winBetslipsDo) Where(conds ...gen.Condition) IWinBetslipsDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winBetslipsDo) Order(conds ...field.Expr) IWinBetslipsDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winBetslipsDo) Distinct(cols ...field.Expr) IWinBetslipsDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winBetslipsDo) Omit(cols ...field.Expr) IWinBetslipsDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winBetslipsDo) Join(table schema.Tabler, on ...field.Expr) IWinBetslipsDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winBetslipsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winBetslipsDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winBetslipsDo) Group(cols ...field.Expr) IWinBetslipsDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winBetslipsDo) Having(conds ...gen.Condition) IWinBetslipsDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winBetslipsDo) Limit(limit int) IWinBetslipsDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winBetslipsDo) Offset(offset int) IWinBetslipsDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winBetslipsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetslipsDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winBetslipsDo) Unscoped() IWinBetslipsDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winBetslipsDo) Create(values ...*model.WinBetslips) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winBetslipsDo) CreateInBatches(values []*model.WinBetslips, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winBetslipsDo) Save(values ...*model.WinBetslips) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winBetslipsDo) First() (*model.WinBetslips, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips), nil
	}
}

func (w winBetslipsDo) Take() (*model.WinBetslips, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips), nil
	}
}

func (w winBetslipsDo) Last() (*model.WinBetslips, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips), nil
	}
}

func (w winBetslipsDo) Find() ([]*model.WinBetslips, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinBetslips), err
}

func (w winBetslipsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetslips, err error) {
	buf := make([]*model.WinBetslips, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winBetslipsDo) FindInBatches(result *[]*model.WinBetslips, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winBetslipsDo) Attrs(attrs ...field.AssignExpr) IWinBetslipsDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winBetslipsDo) Assign(attrs ...field.AssignExpr) IWinBetslipsDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winBetslipsDo) Joins(fields ...field.RelationField) IWinBetslipsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winBetslipsDo) Preload(fields ...field.RelationField) IWinBetslipsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winBetslipsDo) FirstOrInit() (*model.WinBetslips, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips), nil
	}
}

func (w winBetslipsDo) FirstOrCreate() (*model.WinBetslips, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips), nil
	}
}

func (w winBetslipsDo) FindByPage(offset int, limit int) (result []*model.WinBetslips, count int64, err error) {
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

func (w winBetslipsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winBetslipsDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winBetslipsDo) Delete(models ...*model.WinBetslips) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winBetslipsDo) withDO(do gen.Dao) *winBetslipsDo {
	w.DO = *do.(*gen.DO)
	return w
}
