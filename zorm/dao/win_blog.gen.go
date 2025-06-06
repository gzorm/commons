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

func newWinBlog(db *gorm.DB, opts ...gen.DOOption) winBlog {
	_winBlog := winBlog{}

	_winBlog.winBlogDo.UseDB(db, opts...)
	_winBlog.winBlogDo.UseModel(&model.WinBlog{})

	tableName := _winBlog.winBlogDo.TableName()
	_winBlog.ALL = field.NewAsterisk(tableName)
	_winBlog.ID = field.NewInt64(tableName, "id")
	_winBlog.Lang = field.NewString(tableName, "lang")
	_winBlog.Category = field.NewString(tableName, "category")
	_winBlog.Image = field.NewString(tableName, "image")
	_winBlog.Title = field.NewString(tableName, "title")
	_winBlog.TitleSub = field.NewString(tableName, "title_sub")
	_winBlog.Content = field.NewString(tableName, "content")
	_winBlog.Recommend = field.NewInt64(tableName, "recommend")
	_winBlog.Status = field.NewInt64(tableName, "status")
	_winBlog.Sort = field.NewInt64(tableName, "sort")
	_winBlog.CreateBy = field.NewString(tableName, "create_by")
	_winBlog.UpdateBy = field.NewString(tableName, "update_by")
	_winBlog.CreatedAt = field.NewInt64(tableName, "created_at")
	_winBlog.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winBlog.fillFieldMap()

	return _winBlog
}

// winBlog 博客
type winBlog struct {
	winBlogDo

	ALL       field.Asterisk
	ID        field.Int64
	Lang      field.String // 语言
	Category  field.String // 组别名称:SPORTS-体育 DIG_CURRENCY-数字货币 ENTERTAINMENT-娱乐 1XWIN-1xWin新闻 GAMING-博彩
	Image     field.String // 图片
	Title     field.String // 主标题
	TitleSub  field.String // 副标题
	Content   field.String // 内容
	Recommend field.Int64  // 推荐:0-否 1-是
	Status    field.Int64  // 状态:1-启用 0-停用
	Sort      field.Int64  // 排序
	CreateBy  field.String // 创建者
	UpdateBy  field.String // 更新者
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (w winBlog) Table(newTableName string) *winBlog {
	w.winBlogDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winBlog) As(alias string) *winBlog {
	w.winBlogDo.DO = *(w.winBlogDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winBlog) updateTableName(table string) *winBlog {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Lang = field.NewString(table, "lang")
	w.Category = field.NewString(table, "category")
	w.Image = field.NewString(table, "image")
	w.Title = field.NewString(table, "title")
	w.TitleSub = field.NewString(table, "title_sub")
	w.Content = field.NewString(table, "content")
	w.Recommend = field.NewInt64(table, "recommend")
	w.Status = field.NewInt64(table, "status")
	w.Sort = field.NewInt64(table, "sort")
	w.CreateBy = field.NewString(table, "create_by")
	w.UpdateBy = field.NewString(table, "update_by")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winBlog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winBlog) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 14)
	w.fieldMap["id"] = w.ID
	w.fieldMap["lang"] = w.Lang
	w.fieldMap["category"] = w.Category
	w.fieldMap["image"] = w.Image
	w.fieldMap["title"] = w.Title
	w.fieldMap["title_sub"] = w.TitleSub
	w.fieldMap["content"] = w.Content
	w.fieldMap["recommend"] = w.Recommend
	w.fieldMap["status"] = w.Status
	w.fieldMap["sort"] = w.Sort
	w.fieldMap["create_by"] = w.CreateBy
	w.fieldMap["update_by"] = w.UpdateBy
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winBlog) clone(db *gorm.DB) winBlog {
	w.winBlogDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winBlog) replaceDB(db *gorm.DB) winBlog {
	w.winBlogDo.ReplaceDB(db)
	return w
}

type winBlogDo struct{ gen.DO }

type IWinBlogDo interface {
	gen.SubQuery
	Debug() IWinBlogDo
	WithContext(ctx context.Context) IWinBlogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinBlogDo
	WriteDB() IWinBlogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinBlogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinBlogDo
	Not(conds ...gen.Condition) IWinBlogDo
	Or(conds ...gen.Condition) IWinBlogDo
	Select(conds ...field.Expr) IWinBlogDo
	Where(conds ...gen.Condition) IWinBlogDo
	Order(conds ...field.Expr) IWinBlogDo
	Distinct(cols ...field.Expr) IWinBlogDo
	Omit(cols ...field.Expr) IWinBlogDo
	Join(table schema.Tabler, on ...field.Expr) IWinBlogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinBlogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinBlogDo
	Group(cols ...field.Expr) IWinBlogDo
	Having(conds ...gen.Condition) IWinBlogDo
	Limit(limit int) IWinBlogDo
	Offset(offset int) IWinBlogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBlogDo
	Unscoped() IWinBlogDo
	Create(values ...*model.WinBlog) error
	CreateInBatches(values []*model.WinBlog, batchSize int) error
	Save(values ...*model.WinBlog) error
	First() (*model.WinBlog, error)
	Take() (*model.WinBlog, error)
	Last() (*model.WinBlog, error)
	Find() ([]*model.WinBlog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBlog, err error)
	FindInBatches(result *[]*model.WinBlog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinBlog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinBlogDo
	Assign(attrs ...field.AssignExpr) IWinBlogDo
	Joins(fields ...field.RelationField) IWinBlogDo
	Preload(fields ...field.RelationField) IWinBlogDo
	FirstOrInit() (*model.WinBlog, error)
	FirstOrCreate() (*model.WinBlog, error)
	FindByPage(offset int, limit int) (result []*model.WinBlog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinBlogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winBlogDo) Debug() IWinBlogDo {
	return w.withDO(w.DO.Debug())
}

func (w winBlogDo) WithContext(ctx context.Context) IWinBlogDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winBlogDo) ReadDB() IWinBlogDo {
	return w.Clauses(dbresolver.Read)
}

func (w winBlogDo) WriteDB() IWinBlogDo {
	return w.Clauses(dbresolver.Write)
}

func (w winBlogDo) Session(config *gorm.Session) IWinBlogDo {
	return w.withDO(w.DO.Session(config))
}

func (w winBlogDo) Clauses(conds ...clause.Expression) IWinBlogDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winBlogDo) Returning(value interface{}, columns ...string) IWinBlogDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winBlogDo) Not(conds ...gen.Condition) IWinBlogDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winBlogDo) Or(conds ...gen.Condition) IWinBlogDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winBlogDo) Select(conds ...field.Expr) IWinBlogDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winBlogDo) Where(conds ...gen.Condition) IWinBlogDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winBlogDo) Order(conds ...field.Expr) IWinBlogDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winBlogDo) Distinct(cols ...field.Expr) IWinBlogDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winBlogDo) Omit(cols ...field.Expr) IWinBlogDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winBlogDo) Join(table schema.Tabler, on ...field.Expr) IWinBlogDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winBlogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinBlogDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winBlogDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinBlogDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winBlogDo) Group(cols ...field.Expr) IWinBlogDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winBlogDo) Having(conds ...gen.Condition) IWinBlogDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winBlogDo) Limit(limit int) IWinBlogDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winBlogDo) Offset(offset int) IWinBlogDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winBlogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBlogDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winBlogDo) Unscoped() IWinBlogDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winBlogDo) Create(values ...*model.WinBlog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winBlogDo) CreateInBatches(values []*model.WinBlog, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winBlogDo) Save(values ...*model.WinBlog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winBlogDo) First() (*model.WinBlog, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBlog), nil
	}
}

func (w winBlogDo) Take() (*model.WinBlog, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBlog), nil
	}
}

func (w winBlogDo) Last() (*model.WinBlog, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBlog), nil
	}
}

func (w winBlogDo) Find() ([]*model.WinBlog, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinBlog), err
}

func (w winBlogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBlog, err error) {
	buf := make([]*model.WinBlog, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winBlogDo) FindInBatches(result *[]*model.WinBlog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winBlogDo) Attrs(attrs ...field.AssignExpr) IWinBlogDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winBlogDo) Assign(attrs ...field.AssignExpr) IWinBlogDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winBlogDo) Joins(fields ...field.RelationField) IWinBlogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winBlogDo) Preload(fields ...field.RelationField) IWinBlogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winBlogDo) FirstOrInit() (*model.WinBlog, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBlog), nil
	}
}

func (w winBlogDo) FirstOrCreate() (*model.WinBlog, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBlog), nil
	}
}

func (w winBlogDo) FindByPage(offset int, limit int) (result []*model.WinBlog, count int64, err error) {
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

func (w winBlogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winBlogDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winBlogDo) Delete(models ...*model.WinBlog) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winBlogDo) withDO(do gen.Dao) *winBlogDo {
	w.DO = *do.(*gen.DO)
	return w
}
