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

func newFbSportsMatchStatisticsHotLeague(db *gorm.DB, opts ...gen.DOOption) fbSportsMatchStatisticsHotLeague {
	_fbSportsMatchStatisticsHotLeague := fbSportsMatchStatisticsHotLeague{}

	_fbSportsMatchStatisticsHotLeague.fbSportsMatchStatisticsHotLeagueDo.UseDB(db, opts...)
	_fbSportsMatchStatisticsHotLeague.fbSportsMatchStatisticsHotLeagueDo.UseModel(&model.FbSportsMatchStatisticsHotLeague{})

	tableName := _fbSportsMatchStatisticsHotLeague.fbSportsMatchStatisticsHotLeagueDo.TableName()
	_fbSportsMatchStatisticsHotLeague.ALL = field.NewAsterisk(tableName)
	_fbSportsMatchStatisticsHotLeague.ID = field.NewInt64(tableName, "id")
	_fbSportsMatchStatisticsHotLeague.StatisticsID = field.NewInt64(tableName, "statistics_id")
	_fbSportsMatchStatisticsHotLeague.LeagueID = field.NewInt64(tableName, "league_id")
	_fbSportsMatchStatisticsHotLeague.MatchTotal = field.NewInt64(tableName, "match_total")
	_fbSportsMatchStatisticsHotLeague.LeagueName = field.NewString(tableName, "league_name")
	_fbSportsMatchStatisticsHotLeague.LogoURL = field.NewString(tableName, "logo_url")
	_fbSportsMatchStatisticsHotLeague.SportID = field.NewInt64(tableName, "sport_id")
	_fbSportsMatchStatisticsHotLeague.CreatedAt = field.NewInt64(tableName, "created_at")
	_fbSportsMatchStatisticsHotLeague.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_fbSportsMatchStatisticsHotLeague.fillFieldMap()

	return _fbSportsMatchStatisticsHotLeague
}

// fbSportsMatchStatisticsHotLeague 热门联赛统计表，存储热门联赛的赛事统计信息
type fbSportsMatchStatisticsHotLeague struct {
	fbSportsMatchStatisticsHotLeagueDo

	ALL          field.Asterisk
	ID           field.Int64  // 主键ID
	StatisticsID field.Int64  // 关联主表的 ID（对应主表 fb_sports_match_statistics.id）
	LeagueID     field.Int64  // 热门联赛 ID（对应 data.hls.id）
	MatchTotal   field.Int64  // 赛事总数（对应 data.hls.mt）
	LeagueName   field.String // 联赛名称（对应 data.hls.na）
	LogoURL      field.String // 联赛 Logo 的 URL 地址（对应 data.hls.lurl）
	SportID      field.Int64  // 运动种类 ID（对应 data.hls.sid）
	CreatedAt    field.Int64  // 记录创建时间（时间戳）
	UpdatedAt    field.Int64  // 记录更新时间（时间戳）

	fieldMap map[string]field.Expr
}

func (f fbSportsMatchStatisticsHotLeague) Table(newTableName string) *fbSportsMatchStatisticsHotLeague {
	f.fbSportsMatchStatisticsHotLeagueDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbSportsMatchStatisticsHotLeague) As(alias string) *fbSportsMatchStatisticsHotLeague {
	f.fbSportsMatchStatisticsHotLeagueDo.DO = *(f.fbSportsMatchStatisticsHotLeagueDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbSportsMatchStatisticsHotLeague) updateTableName(table string) *fbSportsMatchStatisticsHotLeague {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.StatisticsID = field.NewInt64(table, "statistics_id")
	f.LeagueID = field.NewInt64(table, "league_id")
	f.MatchTotal = field.NewInt64(table, "match_total")
	f.LeagueName = field.NewString(table, "league_name")
	f.LogoURL = field.NewString(table, "logo_url")
	f.SportID = field.NewInt64(table, "sport_id")
	f.CreatedAt = field.NewInt64(table, "created_at")
	f.UpdatedAt = field.NewInt64(table, "updated_at")

	f.fillFieldMap()

	return f
}

func (f *fbSportsMatchStatisticsHotLeague) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbSportsMatchStatisticsHotLeague) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 9)
	f.fieldMap["id"] = f.ID
	f.fieldMap["statistics_id"] = f.StatisticsID
	f.fieldMap["league_id"] = f.LeagueID
	f.fieldMap["match_total"] = f.MatchTotal
	f.fieldMap["league_name"] = f.LeagueName
	f.fieldMap["logo_url"] = f.LogoURL
	f.fieldMap["sport_id"] = f.SportID
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
}

func (f fbSportsMatchStatisticsHotLeague) clone(db *gorm.DB) fbSportsMatchStatisticsHotLeague {
	f.fbSportsMatchStatisticsHotLeagueDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbSportsMatchStatisticsHotLeague) replaceDB(db *gorm.DB) fbSportsMatchStatisticsHotLeague {
	f.fbSportsMatchStatisticsHotLeagueDo.ReplaceDB(db)
	return f
}

type fbSportsMatchStatisticsHotLeagueDo struct{ gen.DO }

type IFbSportsMatchStatisticsHotLeagueDo interface {
	gen.SubQuery
	Debug() IFbSportsMatchStatisticsHotLeagueDo
	WithContext(ctx context.Context) IFbSportsMatchStatisticsHotLeagueDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbSportsMatchStatisticsHotLeagueDo
	WriteDB() IFbSportsMatchStatisticsHotLeagueDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbSportsMatchStatisticsHotLeagueDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbSportsMatchStatisticsHotLeagueDo
	Not(conds ...gen.Condition) IFbSportsMatchStatisticsHotLeagueDo
	Or(conds ...gen.Condition) IFbSportsMatchStatisticsHotLeagueDo
	Select(conds ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo
	Where(conds ...gen.Condition) IFbSportsMatchStatisticsHotLeagueDo
	Order(conds ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo
	Distinct(cols ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo
	Omit(cols ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo
	Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo
	Group(cols ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo
	Having(conds ...gen.Condition) IFbSportsMatchStatisticsHotLeagueDo
	Limit(limit int) IFbSportsMatchStatisticsHotLeagueDo
	Offset(offset int) IFbSportsMatchStatisticsHotLeagueDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchStatisticsHotLeagueDo
	Unscoped() IFbSportsMatchStatisticsHotLeagueDo
	Create(values ...*model.FbSportsMatchStatisticsHotLeague) error
	CreateInBatches(values []*model.FbSportsMatchStatisticsHotLeague, batchSize int) error
	Save(values ...*model.FbSportsMatchStatisticsHotLeague) error
	First() (*model.FbSportsMatchStatisticsHotLeague, error)
	Take() (*model.FbSportsMatchStatisticsHotLeague, error)
	Last() (*model.FbSportsMatchStatisticsHotLeague, error)
	Find() ([]*model.FbSportsMatchStatisticsHotLeague, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatchStatisticsHotLeague, err error)
	FindInBatches(result *[]*model.FbSportsMatchStatisticsHotLeague, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbSportsMatchStatisticsHotLeague) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbSportsMatchStatisticsHotLeagueDo
	Assign(attrs ...field.AssignExpr) IFbSportsMatchStatisticsHotLeagueDo
	Joins(fields ...field.RelationField) IFbSportsMatchStatisticsHotLeagueDo
	Preload(fields ...field.RelationField) IFbSportsMatchStatisticsHotLeagueDo
	FirstOrInit() (*model.FbSportsMatchStatisticsHotLeague, error)
	FirstOrCreate() (*model.FbSportsMatchStatisticsHotLeague, error)
	FindByPage(offset int, limit int) (result []*model.FbSportsMatchStatisticsHotLeague, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbSportsMatchStatisticsHotLeagueDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbSportsMatchStatisticsHotLeagueDo) Debug() IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Debug())
}

func (f fbSportsMatchStatisticsHotLeagueDo) WithContext(ctx context.Context) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbSportsMatchStatisticsHotLeagueDo) ReadDB() IFbSportsMatchStatisticsHotLeagueDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbSportsMatchStatisticsHotLeagueDo) WriteDB() IFbSportsMatchStatisticsHotLeagueDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbSportsMatchStatisticsHotLeagueDo) Session(config *gorm.Session) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Clauses(conds ...clause.Expression) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Returning(value interface{}, columns ...string) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Not(conds ...gen.Condition) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Or(conds ...gen.Condition) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Select(conds ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Where(conds ...gen.Condition) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Order(conds ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Distinct(cols ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Omit(cols ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Group(cols ...field.Expr) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Having(conds ...gen.Condition) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Limit(limit int) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Offset(offset int) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Unscoped() IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbSportsMatchStatisticsHotLeagueDo) Create(values ...*model.FbSportsMatchStatisticsHotLeague) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbSportsMatchStatisticsHotLeagueDo) CreateInBatches(values []*model.FbSportsMatchStatisticsHotLeague, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbSportsMatchStatisticsHotLeagueDo) Save(values ...*model.FbSportsMatchStatisticsHotLeague) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbSportsMatchStatisticsHotLeagueDo) First() (*model.FbSportsMatchStatisticsHotLeague, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatisticsHotLeague), nil
	}
}

func (f fbSportsMatchStatisticsHotLeagueDo) Take() (*model.FbSportsMatchStatisticsHotLeague, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatisticsHotLeague), nil
	}
}

func (f fbSportsMatchStatisticsHotLeagueDo) Last() (*model.FbSportsMatchStatisticsHotLeague, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatisticsHotLeague), nil
	}
}

func (f fbSportsMatchStatisticsHotLeagueDo) Find() ([]*model.FbSportsMatchStatisticsHotLeague, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbSportsMatchStatisticsHotLeague), err
}

func (f fbSportsMatchStatisticsHotLeagueDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatchStatisticsHotLeague, err error) {
	buf := make([]*model.FbSportsMatchStatisticsHotLeague, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbSportsMatchStatisticsHotLeagueDo) FindInBatches(result *[]*model.FbSportsMatchStatisticsHotLeague, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbSportsMatchStatisticsHotLeagueDo) Attrs(attrs ...field.AssignExpr) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Assign(attrs ...field.AssignExpr) IFbSportsMatchStatisticsHotLeagueDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbSportsMatchStatisticsHotLeagueDo) Joins(fields ...field.RelationField) IFbSportsMatchStatisticsHotLeagueDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbSportsMatchStatisticsHotLeagueDo) Preload(fields ...field.RelationField) IFbSportsMatchStatisticsHotLeagueDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbSportsMatchStatisticsHotLeagueDo) FirstOrInit() (*model.FbSportsMatchStatisticsHotLeague, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatisticsHotLeague), nil
	}
}

func (f fbSportsMatchStatisticsHotLeagueDo) FirstOrCreate() (*model.FbSportsMatchStatisticsHotLeague, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatisticsHotLeague), nil
	}
}

func (f fbSportsMatchStatisticsHotLeagueDo) FindByPage(offset int, limit int) (result []*model.FbSportsMatchStatisticsHotLeague, count int64, err error) {
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

func (f fbSportsMatchStatisticsHotLeagueDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbSportsMatchStatisticsHotLeagueDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbSportsMatchStatisticsHotLeagueDo) Delete(models ...*model.FbSportsMatchStatisticsHotLeague) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbSportsMatchStatisticsHotLeagueDo) withDO(do gen.Dao) *fbSportsMatchStatisticsHotLeagueDo {
	f.DO = *do.(*gen.DO)
	return f
}
