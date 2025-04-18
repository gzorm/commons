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

func newWinCodeRecord(db *gorm.DB, opts ...gen.DOOption) winCodeRecord {
	_winCodeRecord := winCodeRecord{}

	_winCodeRecord.winCodeRecordDo.UseDB(db, opts...)
	_winCodeRecord.winCodeRecordDo.UseModel(&model.WinCodeRecord{})

	tableName := _winCodeRecord.winCodeRecordDo.TableName()
	_winCodeRecord.ALL = field.NewAsterisk(tableName)
	_winCodeRecord.ID = field.NewInt64(tableName, "id")
	_winCodeRecord.UID = field.NewInt64(tableName, "uid")
	_winCodeRecord.Username = field.NewString(tableName, "username")
	_winCodeRecord.Coin = field.NewField(tableName, "coin")
	_winCodeRecord.CodeRequire = field.NewField(tableName, "code_require")
	_winCodeRecord.FlowClaim = field.NewField(tableName, "flow_claim")
	_winCodeRecord.RealCode = field.NewField(tableName, "real_code")
	_winCodeRecord.Category = field.NewInt64(tableName, "category")
	_winCodeRecord.ReferID = field.NewInt64(tableName, "refer_id")
	_winCodeRecord.ReferWithdrawalID = field.NewInt64(tableName, "refer_withdrawal_id")
	_winCodeRecord.Remarks = field.NewString(tableName, "remarks")
	_winCodeRecord.Status = field.NewInt64(tableName, "status")
	_winCodeRecord.CreatedAt = field.NewInt64(tableName, "created_at")
	_winCodeRecord.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winCodeRecord.fillFieldMap()

	return _winCodeRecord
}

// winCodeRecord 打码量记录表
type winCodeRecord struct {
	winCodeRecordDo

	ALL               field.Asterisk
	ID                field.Int64
	UID               field.Int64  // UID
	Username          field.String // 用户名
	Coin              field.Field  // 金额
	CodeRequire       field.Field  // 需求打码量
	FlowClaim         field.Field  // 需求打码倍数
	RealCode          field.Field  // 实际打码量
	Category          field.Int64  // 类型:1-充值 2-签到活动 3- 系统调账 4注册活动 5充值活动
	ReferID           field.Int64  // 关联ID
	ReferWithdrawalID field.Int64  // 关联提款ID
	Remarks           field.String // 备注
	Status            field.Int64  // 状态:0-未结算 1-结算
	CreatedAt         field.Int64
	UpdatedAt         field.Int64

	fieldMap map[string]field.Expr
}

func (w winCodeRecord) Table(newTableName string) *winCodeRecord {
	w.winCodeRecordDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCodeRecord) As(alias string) *winCodeRecord {
	w.winCodeRecordDo.DO = *(w.winCodeRecordDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCodeRecord) updateTableName(table string) *winCodeRecord {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.Coin = field.NewField(table, "coin")
	w.CodeRequire = field.NewField(table, "code_require")
	w.FlowClaim = field.NewField(table, "flow_claim")
	w.RealCode = field.NewField(table, "real_code")
	w.Category = field.NewInt64(table, "category")
	w.ReferID = field.NewInt64(table, "refer_id")
	w.ReferWithdrawalID = field.NewInt64(table, "refer_withdrawal_id")
	w.Remarks = field.NewString(table, "remarks")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winCodeRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCodeRecord) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 14)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["code_require"] = w.CodeRequire
	w.fieldMap["flow_claim"] = w.FlowClaim
	w.fieldMap["real_code"] = w.RealCode
	w.fieldMap["category"] = w.Category
	w.fieldMap["refer_id"] = w.ReferID
	w.fieldMap["refer_withdrawal_id"] = w.ReferWithdrawalID
	w.fieldMap["remarks"] = w.Remarks
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winCodeRecord) clone(db *gorm.DB) winCodeRecord {
	w.winCodeRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCodeRecord) replaceDB(db *gorm.DB) winCodeRecord {
	w.winCodeRecordDo.ReplaceDB(db)
	return w
}

type winCodeRecordDo struct{ gen.DO }

type IWinCodeRecordDo interface {
	gen.SubQuery
	Debug() IWinCodeRecordDo
	WithContext(ctx context.Context) IWinCodeRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinCodeRecordDo
	WriteDB() IWinCodeRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinCodeRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinCodeRecordDo
	Not(conds ...gen.Condition) IWinCodeRecordDo
	Or(conds ...gen.Condition) IWinCodeRecordDo
	Select(conds ...field.Expr) IWinCodeRecordDo
	Where(conds ...gen.Condition) IWinCodeRecordDo
	Order(conds ...field.Expr) IWinCodeRecordDo
	Distinct(cols ...field.Expr) IWinCodeRecordDo
	Omit(cols ...field.Expr) IWinCodeRecordDo
	Join(table schema.Tabler, on ...field.Expr) IWinCodeRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinCodeRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinCodeRecordDo
	Group(cols ...field.Expr) IWinCodeRecordDo
	Having(conds ...gen.Condition) IWinCodeRecordDo
	Limit(limit int) IWinCodeRecordDo
	Offset(offset int) IWinCodeRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCodeRecordDo
	Unscoped() IWinCodeRecordDo
	Create(values ...*model.WinCodeRecord) error
	CreateInBatches(values []*model.WinCodeRecord, batchSize int) error
	Save(values ...*model.WinCodeRecord) error
	First() (*model.WinCodeRecord, error)
	Take() (*model.WinCodeRecord, error)
	Last() (*model.WinCodeRecord, error)
	Find() ([]*model.WinCodeRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCodeRecord, err error)
	FindInBatches(result *[]*model.WinCodeRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinCodeRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinCodeRecordDo
	Assign(attrs ...field.AssignExpr) IWinCodeRecordDo
	Joins(fields ...field.RelationField) IWinCodeRecordDo
	Preload(fields ...field.RelationField) IWinCodeRecordDo
	FirstOrInit() (*model.WinCodeRecord, error)
	FirstOrCreate() (*model.WinCodeRecord, error)
	FindByPage(offset int, limit int) (result []*model.WinCodeRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinCodeRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winCodeRecordDo) Debug() IWinCodeRecordDo {
	return w.withDO(w.DO.Debug())
}

func (w winCodeRecordDo) WithContext(ctx context.Context) IWinCodeRecordDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCodeRecordDo) ReadDB() IWinCodeRecordDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCodeRecordDo) WriteDB() IWinCodeRecordDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCodeRecordDo) Session(config *gorm.Session) IWinCodeRecordDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCodeRecordDo) Clauses(conds ...clause.Expression) IWinCodeRecordDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCodeRecordDo) Returning(value interface{}, columns ...string) IWinCodeRecordDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCodeRecordDo) Not(conds ...gen.Condition) IWinCodeRecordDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCodeRecordDo) Or(conds ...gen.Condition) IWinCodeRecordDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCodeRecordDo) Select(conds ...field.Expr) IWinCodeRecordDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCodeRecordDo) Where(conds ...gen.Condition) IWinCodeRecordDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCodeRecordDo) Order(conds ...field.Expr) IWinCodeRecordDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCodeRecordDo) Distinct(cols ...field.Expr) IWinCodeRecordDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCodeRecordDo) Omit(cols ...field.Expr) IWinCodeRecordDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCodeRecordDo) Join(table schema.Tabler, on ...field.Expr) IWinCodeRecordDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCodeRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinCodeRecordDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCodeRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinCodeRecordDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCodeRecordDo) Group(cols ...field.Expr) IWinCodeRecordDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCodeRecordDo) Having(conds ...gen.Condition) IWinCodeRecordDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCodeRecordDo) Limit(limit int) IWinCodeRecordDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCodeRecordDo) Offset(offset int) IWinCodeRecordDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCodeRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCodeRecordDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCodeRecordDo) Unscoped() IWinCodeRecordDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCodeRecordDo) Create(values ...*model.WinCodeRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCodeRecordDo) CreateInBatches(values []*model.WinCodeRecord, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCodeRecordDo) Save(values ...*model.WinCodeRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCodeRecordDo) First() (*model.WinCodeRecord, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecord), nil
	}
}

func (w winCodeRecordDo) Take() (*model.WinCodeRecord, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecord), nil
	}
}

func (w winCodeRecordDo) Last() (*model.WinCodeRecord, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecord), nil
	}
}

func (w winCodeRecordDo) Find() ([]*model.WinCodeRecord, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCodeRecord), err
}

func (w winCodeRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCodeRecord, err error) {
	buf := make([]*model.WinCodeRecord, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCodeRecordDo) FindInBatches(result *[]*model.WinCodeRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCodeRecordDo) Attrs(attrs ...field.AssignExpr) IWinCodeRecordDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCodeRecordDo) Assign(attrs ...field.AssignExpr) IWinCodeRecordDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCodeRecordDo) Joins(fields ...field.RelationField) IWinCodeRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCodeRecordDo) Preload(fields ...field.RelationField) IWinCodeRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCodeRecordDo) FirstOrInit() (*model.WinCodeRecord, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecord), nil
	}
}

func (w winCodeRecordDo) FirstOrCreate() (*model.WinCodeRecord, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCodeRecord), nil
	}
}

func (w winCodeRecordDo) FindByPage(offset int, limit int) (result []*model.WinCodeRecord, count int64, err error) {
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

func (w winCodeRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCodeRecordDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCodeRecordDo) Delete(models ...*model.WinCodeRecord) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCodeRecordDo) withDO(do gen.Dao) *winCodeRecordDo {
	w.DO = *do.(*gen.DO)
	return w
}
