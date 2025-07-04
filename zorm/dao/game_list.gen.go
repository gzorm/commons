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

func newGameList(db *gorm.DB, opts ...gen.DOOption) gameList {
	_gameList := gameList{}

	_gameList.gameListDo.UseDB(db, opts...)
	_gameList.gameListDo.UseModel(&model.GameList{})

	tableName := _gameList.gameListDo.TableName()
	_gameList.ALL = field.NewAsterisk(tableName)
	_gameList.ID = field.NewInt64(tableName, "id")
	_gameList.Code = field.NewString(tableName, "code")
	_gameList.GameProviderSubtypeID = field.NewInt64(tableName, "game_provider_subtype_id")
	_gameList.GamePagcorID = field.NewInt64(tableName, "game_pagcor_id")
	_gameList.GameTypeID = field.NewInt64(tableName, "game_type_id")
	_gameList.GameProviderID = field.NewInt64(tableName, "game_provider_id")
	_gameList.GameStartParam = field.NewString(tableName, "game_start_param")
	_gameList.GameURL = field.NewString(tableName, "game_url")
	_gameList.Sort = field.NewInt64(tableName, "sort")
	_gameList.Status = field.NewInt64(tableName, "status")
	_gameList.Name = field.NewString(tableName, "name")
	_gameList.OriginalIcon = field.NewString(tableName, "original_icon")
	_gameList.LatestIcon = field.NewString(tableName, "latest_icon")
	_gameList.IsNew = field.NewInt64(tableName, "is_new")
	_gameList.IsRotated = field.NewInt64(tableName, "is_rotated")
	_gameList.FavoriteStar = field.NewInt64(tableName, "favorite_star")
	_gameList.HotStar = field.NewInt64(tableName, "hot_star")
	_gameList.Device = field.NewInt64(tableName, "device")
	_gameList.CreatedAt = field.NewInt64(tableName, "created_at")
	_gameList.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_gameList.CreatedBy = field.NewString(tableName, "created_by")
	_gameList.UpdatedBy = field.NewString(tableName, "updated_by")
	_gameList.Maintenance = field.NewString(tableName, "maintenance")
	_gameList.IsCasino = field.NewInt64(tableName, "is_casino")
	_gameList.BetTotalLimit = field.NewField(tableName, "bet_total_limit")
	_gameList.SingleUserBetPercentage = field.NewField(tableName, "single_user_bet_percentage")
	_gameList.BetTotal = field.NewField(tableName, "bet_total")

	_gameList.fillFieldMap()

	return _gameList
}

// gameList 游戏全量表
type gameList struct {
	gameListDo

	ALL                     field.Asterisk
	ID                      field.Int64
	Code                    field.String // 启动游戏编码
	GameProviderSubtypeID   field.Int64  // 关联game_provider_subtype表ID
	GamePagcorID            field.Int64  // pagcor类型id
	GameTypeID              field.Int64  // 游戏类型id
	GameProviderID          field.Int64  // 游戏供应商id
	GameStartParam          field.String // 特殊游戏三方启动参数，如: elbet
	GameURL                 field.String // 三方平台游戏url
	Sort                    field.Int64  // 排序: 从低到高
	Status                  field.Int64  // 状态: 1-启用 0-停用
	Name                    field.String // 简体名称
	OriginalIcon            field.String // 英文图片
	LatestIcon              field.String // 新版游戏图片
	IsNew                   field.Int64  // 是否新游戏:1-是 0-否
	IsRotated               field.Int64  // 1=否(横屏), 3=是(竖屏)
	FavoriteStar            field.Int64  // 收藏值
	HotStar                 field.Int64  // 热度
	Device                  field.Int64  // 设备:0-all 1-pc 2-h5
	CreatedAt               field.Int64
	UpdatedAt               field.Int64
	CreatedBy               field.String // 操作人姓名
	UpdatedBy               field.String // 最后更新人
	Maintenance             field.String // 维护时间
	IsCasino                field.Int64  // 是否isCasino 1是 0否
	BetTotalLimit           field.Field  // 比赛下注总金额限制
	SingleUserBetPercentage field.Field  // 单用户下注金额不超过总金额的百分比比例(数值范围0-100，如50表示50%)
	BetTotal                field.Field  // 比赛当前下注总金额

	fieldMap map[string]field.Expr
}

func (g gameList) Table(newTableName string) *gameList {
	g.gameListDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g gameList) As(alias string) *gameList {
	g.gameListDo.DO = *(g.gameListDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *gameList) updateTableName(table string) *gameList {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt64(table, "id")
	g.Code = field.NewString(table, "code")
	g.GameProviderSubtypeID = field.NewInt64(table, "game_provider_subtype_id")
	g.GamePagcorID = field.NewInt64(table, "game_pagcor_id")
	g.GameTypeID = field.NewInt64(table, "game_type_id")
	g.GameProviderID = field.NewInt64(table, "game_provider_id")
	g.GameStartParam = field.NewString(table, "game_start_param")
	g.GameURL = field.NewString(table, "game_url")
	g.Sort = field.NewInt64(table, "sort")
	g.Status = field.NewInt64(table, "status")
	g.Name = field.NewString(table, "name")
	g.OriginalIcon = field.NewString(table, "original_icon")
	g.LatestIcon = field.NewString(table, "latest_icon")
	g.IsNew = field.NewInt64(table, "is_new")
	g.IsRotated = field.NewInt64(table, "is_rotated")
	g.FavoriteStar = field.NewInt64(table, "favorite_star")
	g.HotStar = field.NewInt64(table, "hot_star")
	g.Device = field.NewInt64(table, "device")
	g.CreatedAt = field.NewInt64(table, "created_at")
	g.UpdatedAt = field.NewInt64(table, "updated_at")
	g.CreatedBy = field.NewString(table, "created_by")
	g.UpdatedBy = field.NewString(table, "updated_by")
	g.Maintenance = field.NewString(table, "maintenance")
	g.IsCasino = field.NewInt64(table, "is_casino")
	g.BetTotalLimit = field.NewField(table, "bet_total_limit")
	g.SingleUserBetPercentage = field.NewField(table, "single_user_bet_percentage")
	g.BetTotal = field.NewField(table, "bet_total")

	g.fillFieldMap()

	return g
}

func (g *gameList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *gameList) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 27)
	g.fieldMap["id"] = g.ID
	g.fieldMap["code"] = g.Code
	g.fieldMap["game_provider_subtype_id"] = g.GameProviderSubtypeID
	g.fieldMap["game_pagcor_id"] = g.GamePagcorID
	g.fieldMap["game_type_id"] = g.GameTypeID
	g.fieldMap["game_provider_id"] = g.GameProviderID
	g.fieldMap["game_start_param"] = g.GameStartParam
	g.fieldMap["game_url"] = g.GameURL
	g.fieldMap["sort"] = g.Sort
	g.fieldMap["status"] = g.Status
	g.fieldMap["name"] = g.Name
	g.fieldMap["original_icon"] = g.OriginalIcon
	g.fieldMap["latest_icon"] = g.LatestIcon
	g.fieldMap["is_new"] = g.IsNew
	g.fieldMap["is_rotated"] = g.IsRotated
	g.fieldMap["favorite_star"] = g.FavoriteStar
	g.fieldMap["hot_star"] = g.HotStar
	g.fieldMap["device"] = g.Device
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
	g.fieldMap["created_by"] = g.CreatedBy
	g.fieldMap["updated_by"] = g.UpdatedBy
	g.fieldMap["maintenance"] = g.Maintenance
	g.fieldMap["is_casino"] = g.IsCasino
	g.fieldMap["bet_total_limit"] = g.BetTotalLimit
	g.fieldMap["single_user_bet_percentage"] = g.SingleUserBetPercentage
	g.fieldMap["bet_total"] = g.BetTotal
}

func (g gameList) clone(db *gorm.DB) gameList {
	g.gameListDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g gameList) replaceDB(db *gorm.DB) gameList {
	g.gameListDo.ReplaceDB(db)
	return g
}

type gameListDo struct{ gen.DO }

type IGameListDo interface {
	gen.SubQuery
	Debug() IGameListDo
	WithContext(ctx context.Context) IGameListDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGameListDo
	WriteDB() IGameListDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGameListDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGameListDo
	Not(conds ...gen.Condition) IGameListDo
	Or(conds ...gen.Condition) IGameListDo
	Select(conds ...field.Expr) IGameListDo
	Where(conds ...gen.Condition) IGameListDo
	Order(conds ...field.Expr) IGameListDo
	Distinct(cols ...field.Expr) IGameListDo
	Omit(cols ...field.Expr) IGameListDo
	Join(table schema.Tabler, on ...field.Expr) IGameListDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGameListDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGameListDo
	Group(cols ...field.Expr) IGameListDo
	Having(conds ...gen.Condition) IGameListDo
	Limit(limit int) IGameListDo
	Offset(offset int) IGameListDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGameListDo
	Unscoped() IGameListDo
	Create(values ...*model.GameList) error
	CreateInBatches(values []*model.GameList, batchSize int) error
	Save(values ...*model.GameList) error
	First() (*model.GameList, error)
	Take() (*model.GameList, error)
	Last() (*model.GameList, error)
	Find() ([]*model.GameList, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GameList, err error)
	FindInBatches(result *[]*model.GameList, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GameList) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGameListDo
	Assign(attrs ...field.AssignExpr) IGameListDo
	Joins(fields ...field.RelationField) IGameListDo
	Preload(fields ...field.RelationField) IGameListDo
	FirstOrInit() (*model.GameList, error)
	FirstOrCreate() (*model.GameList, error)
	FindByPage(offset int, limit int) (result []*model.GameList, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGameListDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g gameListDo) Debug() IGameListDo {
	return g.withDO(g.DO.Debug())
}

func (g gameListDo) WithContext(ctx context.Context) IGameListDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gameListDo) ReadDB() IGameListDo {
	return g.Clauses(dbresolver.Read)
}

func (g gameListDo) WriteDB() IGameListDo {
	return g.Clauses(dbresolver.Write)
}

func (g gameListDo) Session(config *gorm.Session) IGameListDo {
	return g.withDO(g.DO.Session(config))
}

func (g gameListDo) Clauses(conds ...clause.Expression) IGameListDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gameListDo) Returning(value interface{}, columns ...string) IGameListDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g gameListDo) Not(conds ...gen.Condition) IGameListDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gameListDo) Or(conds ...gen.Condition) IGameListDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gameListDo) Select(conds ...field.Expr) IGameListDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gameListDo) Where(conds ...gen.Condition) IGameListDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gameListDo) Order(conds ...field.Expr) IGameListDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gameListDo) Distinct(cols ...field.Expr) IGameListDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gameListDo) Omit(cols ...field.Expr) IGameListDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gameListDo) Join(table schema.Tabler, on ...field.Expr) IGameListDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gameListDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGameListDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gameListDo) RightJoin(table schema.Tabler, on ...field.Expr) IGameListDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gameListDo) Group(cols ...field.Expr) IGameListDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gameListDo) Having(conds ...gen.Condition) IGameListDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gameListDo) Limit(limit int) IGameListDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gameListDo) Offset(offset int) IGameListDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gameListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGameListDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gameListDo) Unscoped() IGameListDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gameListDo) Create(values ...*model.GameList) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gameListDo) CreateInBatches(values []*model.GameList, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gameListDo) Save(values ...*model.GameList) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gameListDo) First() (*model.GameList, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameList), nil
	}
}

func (g gameListDo) Take() (*model.GameList, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameList), nil
	}
}

func (g gameListDo) Last() (*model.GameList, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameList), nil
	}
}

func (g gameListDo) Find() ([]*model.GameList, error) {
	result, err := g.DO.Find()
	return result.([]*model.GameList), err
}

func (g gameListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GameList, err error) {
	buf := make([]*model.GameList, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gameListDo) FindInBatches(result *[]*model.GameList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gameListDo) Attrs(attrs ...field.AssignExpr) IGameListDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gameListDo) Assign(attrs ...field.AssignExpr) IGameListDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gameListDo) Joins(fields ...field.RelationField) IGameListDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g gameListDo) Preload(fields ...field.RelationField) IGameListDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g gameListDo) FirstOrInit() (*model.GameList, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameList), nil
	}
}

func (g gameListDo) FirstOrCreate() (*model.GameList, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameList), nil
	}
}

func (g gameListDo) FindByPage(offset int, limit int) (result []*model.GameList, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g gameListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g gameListDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g gameListDo) Delete(models ...*model.GameList) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *gameListDo) withDO(do gen.Dao) *gameListDo {
	g.DO = *do.(*gen.DO)
	return g
}
