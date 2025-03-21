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

func newWinCoinAdminTransfer(db *gorm.DB, opts ...gen.DOOption) winCoinAdminTransfer {
	_winCoinAdminTransfer := winCoinAdminTransfer{}

	_winCoinAdminTransfer.winCoinAdminTransferDo.UseDB(db, opts...)
	_winCoinAdminTransfer.winCoinAdminTransferDo.UseModel(&model.WinCoinAdminTransfer{})

	tableName := _winCoinAdminTransfer.winCoinAdminTransferDo.TableName()
	_winCoinAdminTransfer.ALL = field.NewAsterisk(tableName)
	_winCoinAdminTransfer.ID = field.NewInt64(tableName, "id")
	_winCoinAdminTransfer.AdminID = field.NewInt64(tableName, "admin_id")
	_winCoinAdminTransfer.Coin = field.NewField(tableName, "coin")
	_winCoinAdminTransfer.CoinBefore = field.NewField(tableName, "coin_before")
	_winCoinAdminTransfer.UID = field.NewInt64(tableName, "uid")
	_winCoinAdminTransfer.Username = field.NewString(tableName, "username")
	_winCoinAdminTransfer.Category = field.NewInt64(tableName, "category")
	_winCoinAdminTransfer.Mark = field.NewString(tableName, "mark")
	_winCoinAdminTransfer.CreatedAt = field.NewInt64(tableName, "created_at")
	_winCoinAdminTransfer.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winCoinAdminTransfer.FlowClaim = field.NewInt64(tableName, "flow_claim")
	_winCoinAdminTransfer.Message = field.NewString(tableName, "message")
	_winCoinAdminTransfer.MerchantID = field.NewInt64(tableName, "merchant_id")
	_winCoinAdminTransfer.Status = field.NewInt64(tableName, "status")

	_winCoinAdminTransfer.fillFieldMap()

	return _winCoinAdminTransfer
}

// winCoinAdminTransfer 后台调账记录
type winCoinAdminTransfer struct {
	winCoinAdminTransferDo

	ALL        field.Asterisk
	ID         field.Int64
	AdminID    field.Int64  // 管理员ID
	Coin       field.Field  // 调账金额
	CoinBefore field.Field  // 调账前金额
	UID        field.Int64  // 用户ID
	Username   field.String // 用户名
	Category   field.Int64  // 调账原因:0-其他 1-误存调账 2-活动调账
	Mark       field.String // 调账原因
	CreatedAt  field.Int64
	UpdatedAt  field.Int64
	FlowClaim  field.Int64  // 流水倍数
	Message    field.String // 通知客户信息
	MerchantID field.Int64  // 商户id
	Status     field.Int64  // 1=未使用, 2=调账完成

	fieldMap map[string]field.Expr
}

func (w winCoinAdminTransfer) Table(newTableName string) *winCoinAdminTransfer {
	w.winCoinAdminTransferDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCoinAdminTransfer) As(alias string) *winCoinAdminTransfer {
	w.winCoinAdminTransferDo.DO = *(w.winCoinAdminTransferDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCoinAdminTransfer) updateTableName(table string) *winCoinAdminTransfer {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.AdminID = field.NewInt64(table, "admin_id")
	w.Coin = field.NewField(table, "coin")
	w.CoinBefore = field.NewField(table, "coin_before")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.Category = field.NewInt64(table, "category")
	w.Mark = field.NewString(table, "mark")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.FlowClaim = field.NewInt64(table, "flow_claim")
	w.Message = field.NewString(table, "message")
	w.MerchantID = field.NewInt64(table, "merchant_id")
	w.Status = field.NewInt64(table, "status")

	w.fillFieldMap()

	return w
}

func (w *winCoinAdminTransfer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCoinAdminTransfer) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 14)
	w.fieldMap["id"] = w.ID
	w.fieldMap["admin_id"] = w.AdminID
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["category"] = w.Category
	w.fieldMap["mark"] = w.Mark
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["flow_claim"] = w.FlowClaim
	w.fieldMap["message"] = w.Message
	w.fieldMap["merchant_id"] = w.MerchantID
	w.fieldMap["status"] = w.Status
}

func (w winCoinAdminTransfer) clone(db *gorm.DB) winCoinAdminTransfer {
	w.winCoinAdminTransferDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCoinAdminTransfer) replaceDB(db *gorm.DB) winCoinAdminTransfer {
	w.winCoinAdminTransferDo.ReplaceDB(db)
	return w
}

type winCoinAdminTransferDo struct{ gen.DO }

type IWinCoinAdminTransferDo interface {
	gen.SubQuery
	Debug() IWinCoinAdminTransferDo
	WithContext(ctx context.Context) IWinCoinAdminTransferDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinCoinAdminTransferDo
	WriteDB() IWinCoinAdminTransferDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinCoinAdminTransferDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinCoinAdminTransferDo
	Not(conds ...gen.Condition) IWinCoinAdminTransferDo
	Or(conds ...gen.Condition) IWinCoinAdminTransferDo
	Select(conds ...field.Expr) IWinCoinAdminTransferDo
	Where(conds ...gen.Condition) IWinCoinAdminTransferDo
	Order(conds ...field.Expr) IWinCoinAdminTransferDo
	Distinct(cols ...field.Expr) IWinCoinAdminTransferDo
	Omit(cols ...field.Expr) IWinCoinAdminTransferDo
	Join(table schema.Tabler, on ...field.Expr) IWinCoinAdminTransferDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinAdminTransferDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinAdminTransferDo
	Group(cols ...field.Expr) IWinCoinAdminTransferDo
	Having(conds ...gen.Condition) IWinCoinAdminTransferDo
	Limit(limit int) IWinCoinAdminTransferDo
	Offset(offset int) IWinCoinAdminTransferDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinAdminTransferDo
	Unscoped() IWinCoinAdminTransferDo
	Create(values ...*model.WinCoinAdminTransfer) error
	CreateInBatches(values []*model.WinCoinAdminTransfer, batchSize int) error
	Save(values ...*model.WinCoinAdminTransfer) error
	First() (*model.WinCoinAdminTransfer, error)
	Take() (*model.WinCoinAdminTransfer, error)
	Last() (*model.WinCoinAdminTransfer, error)
	Find() ([]*model.WinCoinAdminTransfer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinAdminTransfer, err error)
	FindInBatches(result *[]*model.WinCoinAdminTransfer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinCoinAdminTransfer) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinCoinAdminTransferDo
	Assign(attrs ...field.AssignExpr) IWinCoinAdminTransferDo
	Joins(fields ...field.RelationField) IWinCoinAdminTransferDo
	Preload(fields ...field.RelationField) IWinCoinAdminTransferDo
	FirstOrInit() (*model.WinCoinAdminTransfer, error)
	FirstOrCreate() (*model.WinCoinAdminTransfer, error)
	FindByPage(offset int, limit int) (result []*model.WinCoinAdminTransfer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinCoinAdminTransferDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winCoinAdminTransferDo) Debug() IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Debug())
}

func (w winCoinAdminTransferDo) WithContext(ctx context.Context) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCoinAdminTransferDo) ReadDB() IWinCoinAdminTransferDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCoinAdminTransferDo) WriteDB() IWinCoinAdminTransferDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCoinAdminTransferDo) Session(config *gorm.Session) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCoinAdminTransferDo) Clauses(conds ...clause.Expression) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCoinAdminTransferDo) Returning(value interface{}, columns ...string) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCoinAdminTransferDo) Not(conds ...gen.Condition) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCoinAdminTransferDo) Or(conds ...gen.Condition) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCoinAdminTransferDo) Select(conds ...field.Expr) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCoinAdminTransferDo) Where(conds ...gen.Condition) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCoinAdminTransferDo) Order(conds ...field.Expr) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCoinAdminTransferDo) Distinct(cols ...field.Expr) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCoinAdminTransferDo) Omit(cols ...field.Expr) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCoinAdminTransferDo) Join(table schema.Tabler, on ...field.Expr) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCoinAdminTransferDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCoinAdminTransferDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCoinAdminTransferDo) Group(cols ...field.Expr) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCoinAdminTransferDo) Having(conds ...gen.Condition) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCoinAdminTransferDo) Limit(limit int) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCoinAdminTransferDo) Offset(offset int) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCoinAdminTransferDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCoinAdminTransferDo) Unscoped() IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCoinAdminTransferDo) Create(values ...*model.WinCoinAdminTransfer) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCoinAdminTransferDo) CreateInBatches(values []*model.WinCoinAdminTransfer, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCoinAdminTransferDo) Save(values ...*model.WinCoinAdminTransfer) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCoinAdminTransferDo) First() (*model.WinCoinAdminTransfer, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinAdminTransfer), nil
	}
}

func (w winCoinAdminTransferDo) Take() (*model.WinCoinAdminTransfer, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinAdminTransfer), nil
	}
}

func (w winCoinAdminTransferDo) Last() (*model.WinCoinAdminTransfer, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinAdminTransfer), nil
	}
}

func (w winCoinAdminTransferDo) Find() ([]*model.WinCoinAdminTransfer, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCoinAdminTransfer), err
}

func (w winCoinAdminTransferDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinAdminTransfer, err error) {
	buf := make([]*model.WinCoinAdminTransfer, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCoinAdminTransferDo) FindInBatches(result *[]*model.WinCoinAdminTransfer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCoinAdminTransferDo) Attrs(attrs ...field.AssignExpr) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCoinAdminTransferDo) Assign(attrs ...field.AssignExpr) IWinCoinAdminTransferDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCoinAdminTransferDo) Joins(fields ...field.RelationField) IWinCoinAdminTransferDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCoinAdminTransferDo) Preload(fields ...field.RelationField) IWinCoinAdminTransferDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCoinAdminTransferDo) FirstOrInit() (*model.WinCoinAdminTransfer, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinAdminTransfer), nil
	}
}

func (w winCoinAdminTransferDo) FirstOrCreate() (*model.WinCoinAdminTransfer, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinAdminTransfer), nil
	}
}

func (w winCoinAdminTransferDo) FindByPage(offset int, limit int) (result []*model.WinCoinAdminTransfer, count int64, err error) {
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

func (w winCoinAdminTransferDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCoinAdminTransferDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCoinAdminTransferDo) Delete(models ...*model.WinCoinAdminTransfer) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCoinAdminTransferDo) withDO(do gen.Dao) *winCoinAdminTransferDo {
	w.DO = *do.(*gen.DO)
	return w
}
