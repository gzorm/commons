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

func newWinUserKyc(db *gorm.DB, opts ...gen.DOOption) winUserKyc {
	_winUserKyc := winUserKyc{}

	_winUserKyc.winUserKycDo.UseDB(db, opts...)
	_winUserKyc.winUserKycDo.UseModel(&model.WinUserKyc{})

	tableName := _winUserKyc.winUserKycDo.TableName()
	_winUserKyc.ALL = field.NewAsterisk(tableName)
	_winUserKyc.ID = field.NewInt64(tableName, "id")
	_winUserKyc.UID = field.NewInt64(tableName, "uid")
	_winUserKyc.Username = field.NewString(tableName, "username")
	_winUserKyc.MerchantID = field.NewInt64(tableName, "merchant_id")
	_winUserKyc.FirstName = field.NewString(tableName, "first_name")
	_winUserKyc.MiddleName = field.NewString(tableName, "middle_name")
	_winUserKyc.LastName = field.NewString(tableName, "last_name")
	_winUserKyc.Birthday = field.NewInt64(tableName, "birthday")
	_winUserKyc.ImgFront = field.NewString(tableName, "img_front")
	_winUserKyc.ImgBack = field.NewString(tableName, "img_back")
	_winUserKyc.IDType = field.NewInt64(tableName, "id_type")
	_winUserKyc.IDNumber = field.NewString(tableName, "id_number")
	_winUserKyc.Status = field.NewInt64(tableName, "status")
	_winUserKyc.Reason = field.NewString(tableName, "reason")
	_winUserKyc.Remark = field.NewString(tableName, "remark")
	_winUserKyc.AuditUsername = field.NewString(tableName, "audit_username")
	_winUserKyc.AuditAt = field.NewInt64(tableName, "audit_at")
	_winUserKyc.CreatedAt = field.NewInt64(tableName, "created_at")
	_winUserKyc.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winUserKyc.OperatorName = field.NewString(tableName, "operator_name")
	_winUserKyc.Nationality = field.NewString(tableName, "nationality")
	_winUserKyc.PlaceOfBirth = field.NewString(tableName, "place_of_birth")
	_winUserKyc.CurrentAddress = field.NewString(tableName, "current_address")
	_winUserKyc.PermanentAddress = field.NewString(tableName, "permanent_address")
	_winUserKyc.EmploymentType = field.NewString(tableName, "employment_type")
	_winUserKyc.MainSourceOfFunds = field.NewString(tableName, "main_source_of_funds")
	_winUserKyc.RelationStore = field.NewString(tableName, "relation_store")
	_winUserKyc.Age = field.NewInt64(tableName, "age")
	_winUserKyc.Gender = field.NewInt64(tableName, "gender")
	_winUserKyc.CivilStatus = field.NewInt64(tableName, "civil_status")
	_winUserKyc.EmailAddress = field.NewString(tableName, "email_address")
	_winUserKyc.ContactNo = field.NewString(tableName, "contact_no")
	_winUserKyc.ImgHandheld = field.NewString(tableName, "img_handheld")
	_winUserKyc.City = field.NewString(tableName, "city")
	_winUserKyc.Province = field.NewString(tableName, "province")
	_winUserKyc.ReferenceID = field.NewString(tableName, "reference_id")
	_winUserKyc.ReferenceURL = field.NewString(tableName, "reference_url")
	_winUserKyc.DeclinedReason = field.NewString(tableName, "declined_reason")

	_winUserKyc.fillFieldMap()

	return _winUserKyc
}

// winUserKyc kyc表
type winUserKyc struct {
	winUserKycDo

	ALL               field.Asterisk
	ID                field.Int64
	UID               field.Int64  // uid
	Username          field.String // 用户名
	MerchantID        field.Int64  // 商户id
	FirstName         field.String // first_name
	MiddleName        field.String // middle_name
	LastName          field.String // last_name
	Birthday          field.Int64  // birthday
	ImgFront          field.String // 证件正面图片地址
	ImgBack           field.String // 证件背面图片地址
	IDType            field.Int64  // 证件类型:1-身份证 2-护照 3-驾照
	IDNumber          field.String // 证件编号
	Status            field.Int64  // 状态:1-待审核 2-通过 3-拒绝
	Reason            field.String // 拒绝原因
	Remark            field.String // 备注
	AuditUsername     field.String // 审核人
	AuditAt           field.Int64  // 审核时间
	CreatedAt         field.Int64
	UpdatedAt         field.Int64
	OperatorName      field.String // 操作人姓名
	Nationality       field.String // 国籍
	PlaceOfBirth      field.String // 出生地
	CurrentAddress    field.String // 当前地址
	PermanentAddress  field.String // 长期地址
	EmploymentType    field.String // 工作性质1,Government Employee 2,OFW 3,Private Sector Employee 4,Self-Employed 5,Student 6,Unemployed
	MainSourceOfFunds field.String // 收入来源1，Allowance/Pension/Stipend/Honoraria2，Campaign Funds/Donation3，Cash on Hand4，Commission or Incentives5，E-Money6，Income from Business/Profession7，Inheritance8，Investment9，Loan Availments10，Personal Savings11，Proceeds of Personal or Real property sale12，Remittance from relatives13，Rental Income14，Salary
	RelationStore     field.String // 关联商铺
	Age               field.Int64  // 年龄
	Gender            field.Int64  // 性别：1.Male2.Female 3.Prefer not to say
	CivilStatus       field.Int64  // 婚姻状态：1. Single2. Married 3. Divorced 4. Widowed
	EmailAddress      field.String // 邮箱地址
	ContactNo         field.String // 联系方式
	ImgHandheld       field.String // 手持证件照
	City              field.String // 城市
	Province          field.String // 当前所在省
	ReferenceID       field.String // 三方订单ID
	ReferenceURL      field.String // 三方kyc URL
	DeclinedReason    field.String // 三方拒绝理由

	fieldMap map[string]field.Expr
}

func (w winUserKyc) Table(newTableName string) *winUserKyc {
	w.winUserKycDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winUserKyc) As(alias string) *winUserKyc {
	w.winUserKycDo.DO = *(w.winUserKycDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winUserKyc) updateTableName(table string) *winUserKyc {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.MerchantID = field.NewInt64(table, "merchant_id")
	w.FirstName = field.NewString(table, "first_name")
	w.MiddleName = field.NewString(table, "middle_name")
	w.LastName = field.NewString(table, "last_name")
	w.Birthday = field.NewInt64(table, "birthday")
	w.ImgFront = field.NewString(table, "img_front")
	w.ImgBack = field.NewString(table, "img_back")
	w.IDType = field.NewInt64(table, "id_type")
	w.IDNumber = field.NewString(table, "id_number")
	w.Status = field.NewInt64(table, "status")
	w.Reason = field.NewString(table, "reason")
	w.Remark = field.NewString(table, "remark")
	w.AuditUsername = field.NewString(table, "audit_username")
	w.AuditAt = field.NewInt64(table, "audit_at")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.OperatorName = field.NewString(table, "operator_name")
	w.Nationality = field.NewString(table, "nationality")
	w.PlaceOfBirth = field.NewString(table, "place_of_birth")
	w.CurrentAddress = field.NewString(table, "current_address")
	w.PermanentAddress = field.NewString(table, "permanent_address")
	w.EmploymentType = field.NewString(table, "employment_type")
	w.MainSourceOfFunds = field.NewString(table, "main_source_of_funds")
	w.RelationStore = field.NewString(table, "relation_store")
	w.Age = field.NewInt64(table, "age")
	w.Gender = field.NewInt64(table, "gender")
	w.CivilStatus = field.NewInt64(table, "civil_status")
	w.EmailAddress = field.NewString(table, "email_address")
	w.ContactNo = field.NewString(table, "contact_no")
	w.ImgHandheld = field.NewString(table, "img_handheld")
	w.City = field.NewString(table, "city")
	w.Province = field.NewString(table, "province")
	w.ReferenceID = field.NewString(table, "reference_id")
	w.ReferenceURL = field.NewString(table, "reference_url")
	w.DeclinedReason = field.NewString(table, "declined_reason")

	w.fillFieldMap()

	return w
}

func (w *winUserKyc) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winUserKyc) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 38)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["merchant_id"] = w.MerchantID
	w.fieldMap["first_name"] = w.FirstName
	w.fieldMap["middle_name"] = w.MiddleName
	w.fieldMap["last_name"] = w.LastName
	w.fieldMap["birthday"] = w.Birthday
	w.fieldMap["img_front"] = w.ImgFront
	w.fieldMap["img_back"] = w.ImgBack
	w.fieldMap["id_type"] = w.IDType
	w.fieldMap["id_number"] = w.IDNumber
	w.fieldMap["status"] = w.Status
	w.fieldMap["reason"] = w.Reason
	w.fieldMap["remark"] = w.Remark
	w.fieldMap["audit_username"] = w.AuditUsername
	w.fieldMap["audit_at"] = w.AuditAt
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["operator_name"] = w.OperatorName
	w.fieldMap["nationality"] = w.Nationality
	w.fieldMap["place_of_birth"] = w.PlaceOfBirth
	w.fieldMap["current_address"] = w.CurrentAddress
	w.fieldMap["permanent_address"] = w.PermanentAddress
	w.fieldMap["employment_type"] = w.EmploymentType
	w.fieldMap["main_source_of_funds"] = w.MainSourceOfFunds
	w.fieldMap["relation_store"] = w.RelationStore
	w.fieldMap["age"] = w.Age
	w.fieldMap["gender"] = w.Gender
	w.fieldMap["civil_status"] = w.CivilStatus
	w.fieldMap["email_address"] = w.EmailAddress
	w.fieldMap["contact_no"] = w.ContactNo
	w.fieldMap["img_handheld"] = w.ImgHandheld
	w.fieldMap["city"] = w.City
	w.fieldMap["province"] = w.Province
	w.fieldMap["reference_id"] = w.ReferenceID
	w.fieldMap["reference_url"] = w.ReferenceURL
	w.fieldMap["declined_reason"] = w.DeclinedReason
}

func (w winUserKyc) clone(db *gorm.DB) winUserKyc {
	w.winUserKycDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winUserKyc) replaceDB(db *gorm.DB) winUserKyc {
	w.winUserKycDo.ReplaceDB(db)
	return w
}

type winUserKycDo struct{ gen.DO }

type IWinUserKycDo interface {
	gen.SubQuery
	Debug() IWinUserKycDo
	WithContext(ctx context.Context) IWinUserKycDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinUserKycDo
	WriteDB() IWinUserKycDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinUserKycDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinUserKycDo
	Not(conds ...gen.Condition) IWinUserKycDo
	Or(conds ...gen.Condition) IWinUserKycDo
	Select(conds ...field.Expr) IWinUserKycDo
	Where(conds ...gen.Condition) IWinUserKycDo
	Order(conds ...field.Expr) IWinUserKycDo
	Distinct(cols ...field.Expr) IWinUserKycDo
	Omit(cols ...field.Expr) IWinUserKycDo
	Join(table schema.Tabler, on ...field.Expr) IWinUserKycDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserKycDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinUserKycDo
	Group(cols ...field.Expr) IWinUserKycDo
	Having(conds ...gen.Condition) IWinUserKycDo
	Limit(limit int) IWinUserKycDo
	Offset(offset int) IWinUserKycDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserKycDo
	Unscoped() IWinUserKycDo
	Create(values ...*model.WinUserKyc) error
	CreateInBatches(values []*model.WinUserKyc, batchSize int) error
	Save(values ...*model.WinUserKyc) error
	First() (*model.WinUserKyc, error)
	Take() (*model.WinUserKyc, error)
	Last() (*model.WinUserKyc, error)
	Find() ([]*model.WinUserKyc, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserKyc, err error)
	FindInBatches(result *[]*model.WinUserKyc, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinUserKyc) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinUserKycDo
	Assign(attrs ...field.AssignExpr) IWinUserKycDo
	Joins(fields ...field.RelationField) IWinUserKycDo
	Preload(fields ...field.RelationField) IWinUserKycDo
	FirstOrInit() (*model.WinUserKyc, error)
	FirstOrCreate() (*model.WinUserKyc, error)
	FindByPage(offset int, limit int) (result []*model.WinUserKyc, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinUserKycDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winUserKycDo) Debug() IWinUserKycDo {
	return w.withDO(w.DO.Debug())
}

func (w winUserKycDo) WithContext(ctx context.Context) IWinUserKycDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winUserKycDo) ReadDB() IWinUserKycDo {
	return w.Clauses(dbresolver.Read)
}

func (w winUserKycDo) WriteDB() IWinUserKycDo {
	return w.Clauses(dbresolver.Write)
}

func (w winUserKycDo) Session(config *gorm.Session) IWinUserKycDo {
	return w.withDO(w.DO.Session(config))
}

func (w winUserKycDo) Clauses(conds ...clause.Expression) IWinUserKycDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winUserKycDo) Returning(value interface{}, columns ...string) IWinUserKycDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winUserKycDo) Not(conds ...gen.Condition) IWinUserKycDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winUserKycDo) Or(conds ...gen.Condition) IWinUserKycDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winUserKycDo) Select(conds ...field.Expr) IWinUserKycDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winUserKycDo) Where(conds ...gen.Condition) IWinUserKycDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winUserKycDo) Order(conds ...field.Expr) IWinUserKycDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winUserKycDo) Distinct(cols ...field.Expr) IWinUserKycDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winUserKycDo) Omit(cols ...field.Expr) IWinUserKycDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winUserKycDo) Join(table schema.Tabler, on ...field.Expr) IWinUserKycDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winUserKycDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserKycDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winUserKycDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinUserKycDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winUserKycDo) Group(cols ...field.Expr) IWinUserKycDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winUserKycDo) Having(conds ...gen.Condition) IWinUserKycDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winUserKycDo) Limit(limit int) IWinUserKycDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winUserKycDo) Offset(offset int) IWinUserKycDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winUserKycDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserKycDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winUserKycDo) Unscoped() IWinUserKycDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winUserKycDo) Create(values ...*model.WinUserKyc) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winUserKycDo) CreateInBatches(values []*model.WinUserKyc, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winUserKycDo) Save(values ...*model.WinUserKyc) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winUserKycDo) First() (*model.WinUserKyc, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserKyc), nil
	}
}

func (w winUserKycDo) Take() (*model.WinUserKyc, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserKyc), nil
	}
}

func (w winUserKycDo) Last() (*model.WinUserKyc, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserKyc), nil
	}
}

func (w winUserKycDo) Find() ([]*model.WinUserKyc, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinUserKyc), err
}

func (w winUserKycDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserKyc, err error) {
	buf := make([]*model.WinUserKyc, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winUserKycDo) FindInBatches(result *[]*model.WinUserKyc, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winUserKycDo) Attrs(attrs ...field.AssignExpr) IWinUserKycDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winUserKycDo) Assign(attrs ...field.AssignExpr) IWinUserKycDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winUserKycDo) Joins(fields ...field.RelationField) IWinUserKycDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winUserKycDo) Preload(fields ...field.RelationField) IWinUserKycDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winUserKycDo) FirstOrInit() (*model.WinUserKyc, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserKyc), nil
	}
}

func (w winUserKycDo) FirstOrCreate() (*model.WinUserKyc, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserKyc), nil
	}
}

func (w winUserKycDo) FindByPage(offset int, limit int) (result []*model.WinUserKyc, count int64, err error) {
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

func (w winUserKycDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winUserKycDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winUserKycDo) Delete(models ...*model.WinUserKyc) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winUserKycDo) withDO(do gen.Dao) *winUserKycDo {
	w.DO = *do.(*gen.DO)
	return w
}
