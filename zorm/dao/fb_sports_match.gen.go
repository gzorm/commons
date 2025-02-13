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

func newFbSportsMatch(db *gorm.DB, opts ...gen.DOOption) fbSportsMatch {
	_fbSportsMatch := fbSportsMatch{}

	_fbSportsMatch.fbSportsMatchDo.UseDB(db, opts...)
	_fbSportsMatch.fbSportsMatchDo.UseModel(&model.FbSportsMatch{})

	tableName := _fbSportsMatch.fbSportsMatchDo.TableName()
	_fbSportsMatch.ALL = field.NewAsterisk(tableName)
	_fbSportsMatch.ID = field.NewInt64(tableName, "id")
	_fbSportsMatch.LanguageType = field.NewString(tableName, "language_type")
	_fbSportsMatch.MatchID = field.NewInt64(tableName, "match_id")
	_fbSportsMatch.MatchName = field.NewString(tableName, "match_name")
	_fbSportsMatch.SportID = field.NewInt64(tableName, "sport_id")
	_fbSportsMatch.MatchStatus = field.NewInt64(tableName, "match_status")
	_fbSportsMatch.StartTime = field.NewInt64(tableName, "start_time")
	_fbSportsMatch.MarketCount = field.NewInt64(tableName, "market_count")
	_fbSportsMatch.Neutral = field.NewInt64(tableName, "neutral")
	_fbSportsMatch.ScoreList = field.NewString(tableName, "score_list")
	_fbSportsMatch.OddsList = field.NewString(tableName, "odds_list")
	_fbSportsMatch.LeagueInfo = field.NewString(tableName, "league_info")
	_fbSportsMatch.TeamInfo = field.NewString(tableName, "team_info")
	_fbSportsMatch.MatchClock = field.NewString(tableName, "match_clock")
	_fbSportsMatch.MarketTags = field.NewString(tableName, "market_tags")
	_fbSportsMatch.MatchFormatID = field.NewInt64(tableName, "match_format_id")
	_fbSportsMatch.MatchFormat = field.NewInt64(tableName, "match_format")
	_fbSportsMatch.LiveUrls = field.NewString(tableName, "live_urls")
	_fbSportsMatch.ServingSide = field.NewInt64(tableName, "serving_side")
	_fbSportsMatch.MatchPhase = field.NewString(tableName, "match_phase")
	_fbSportsMatch.SellingStatus = field.NewInt64(tableName, "selling_status")
	_fbSportsMatch.SubMatchType = field.NewInt64(tableName, "sub_match_type")
	_fbSportsMatch.MatchType = field.NewInt64(tableName, "match_type")
	_fbSportsMatch.Season = field.NewString(tableName, "season")
	_fbSportsMatch.Scoreboard = field.NewString(tableName, "scoreboard")
	_fbSportsMatch.IsInPlay = field.NewInt64(tableName, "is_in_play")
	_fbSportsMatch.IsRecommend = field.NewInt64(tableName, "is_recommend")
	_fbSportsMatch.CreatedAt = field.NewInt64(tableName, "created_at")
	_fbSportsMatch.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_fbSportsMatch.fillFieldMap()

	return _fbSportsMatch
}

// fbSportsMatch 赛事表
type fbSportsMatch struct {
	fbSportsMatchDo

	ALL           field.Asterisk
	ID            field.Int64  // 主键ID
	LanguageType  field.String // 语言类型
	MatchID       field.Int64  // 赛事ID
	MatchName     field.String // 冠军赛赛事名称，用于展示（原 nm）
	SportID       field.Int64  // 运动ID（原 sid）
	MatchStatus   field.Int64  // 赛事进行状态（原 ms）
	StartTime     field.Int64  // 开赛时间（原 bt）
	MarketCount   field.Int64  // 单个赛事玩法总数（原 tms）
	Neutral       field.Int64  // 中立场标记 0 非中立场 ，1 中立场（原 ne）
	ScoreList     field.String // 比分列表，提供各个赛事阶段的比分（原 nsg）
	OddsList      field.String // 赔率列表（原 mg）
	LeagueInfo    field.String // 联赛信息（原 lg）
	TeamInfo      field.String // 球队信息，包含主队和客队（原 ts）
	MatchClock    field.String // 比赛时钟信息（原 mc）
	MarketTags    field.String // 盘口组标签集合（原 tps）
	MatchFormatID field.Int64  // 赛制的场次、局数、节数（原 fid）
	MatchFormat   field.Int64  // 赛制（原 fmt）
	LiveUrls      field.String // 动画直播地址集合（原 as）
	ServingSide   field.Int64  // 主/客发球 1主队发球 2客队发球（原 ssi）
	MatchPhase    field.String // 赛事辅助标记，如附加赛、季前赛等（原 mp）
	SellingStatus field.Int64  // 销售状态 1 开售，2 暂停，其他状态都不展示（原 ss）
	SubMatchType  field.Int64  // 滚球赛事当前阶段标识：常规时间，加时赛，点球大战等（原 smt）
	MatchType     field.Int64  // 赛事类型 1 冠军投注赛事，2 正常赛事（原 ty）
	Season        field.String // 冠军赛事联赛赛季，如：2019年（原 ye）
	Scoreboard    field.String // 比分板（原 sb）
	IsInPlay      field.Int64  // 是否可以开售滚球盘口 0 否，1 是（原 pl）
	IsRecommend   field.Int64  // 是否推荐  1==否 3==是
	CreatedAt     field.Int64  // 创建时间
	UpdatedAt     field.Int64  // 更新时间

	fieldMap map[string]field.Expr
}

func (f fbSportsMatch) Table(newTableName string) *fbSportsMatch {
	f.fbSportsMatchDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbSportsMatch) As(alias string) *fbSportsMatch {
	f.fbSportsMatchDo.DO = *(f.fbSportsMatchDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbSportsMatch) updateTableName(table string) *fbSportsMatch {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.LanguageType = field.NewString(table, "language_type")
	f.MatchID = field.NewInt64(table, "match_id")
	f.MatchName = field.NewString(table, "match_name")
	f.SportID = field.NewInt64(table, "sport_id")
	f.MatchStatus = field.NewInt64(table, "match_status")
	f.StartTime = field.NewInt64(table, "start_time")
	f.MarketCount = field.NewInt64(table, "market_count")
	f.Neutral = field.NewInt64(table, "neutral")
	f.ScoreList = field.NewString(table, "score_list")
	f.OddsList = field.NewString(table, "odds_list")
	f.LeagueInfo = field.NewString(table, "league_info")
	f.TeamInfo = field.NewString(table, "team_info")
	f.MatchClock = field.NewString(table, "match_clock")
	f.MarketTags = field.NewString(table, "market_tags")
	f.MatchFormatID = field.NewInt64(table, "match_format_id")
	f.MatchFormat = field.NewInt64(table, "match_format")
	f.LiveUrls = field.NewString(table, "live_urls")
	f.ServingSide = field.NewInt64(table, "serving_side")
	f.MatchPhase = field.NewString(table, "match_phase")
	f.SellingStatus = field.NewInt64(table, "selling_status")
	f.SubMatchType = field.NewInt64(table, "sub_match_type")
	f.MatchType = field.NewInt64(table, "match_type")
	f.Season = field.NewString(table, "season")
	f.Scoreboard = field.NewString(table, "scoreboard")
	f.IsInPlay = field.NewInt64(table, "is_in_play")
	f.IsRecommend = field.NewInt64(table, "is_recommend")
	f.CreatedAt = field.NewInt64(table, "created_at")
	f.UpdatedAt = field.NewInt64(table, "updated_at")

	f.fillFieldMap()

	return f
}

func (f *fbSportsMatch) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbSportsMatch) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 29)
	f.fieldMap["id"] = f.ID
	f.fieldMap["language_type"] = f.LanguageType
	f.fieldMap["match_id"] = f.MatchID
	f.fieldMap["match_name"] = f.MatchName
	f.fieldMap["sport_id"] = f.SportID
	f.fieldMap["match_status"] = f.MatchStatus
	f.fieldMap["start_time"] = f.StartTime
	f.fieldMap["market_count"] = f.MarketCount
	f.fieldMap["neutral"] = f.Neutral
	f.fieldMap["score_list"] = f.ScoreList
	f.fieldMap["odds_list"] = f.OddsList
	f.fieldMap["league_info"] = f.LeagueInfo
	f.fieldMap["team_info"] = f.TeamInfo
	f.fieldMap["match_clock"] = f.MatchClock
	f.fieldMap["market_tags"] = f.MarketTags
	f.fieldMap["match_format_id"] = f.MatchFormatID
	f.fieldMap["match_format"] = f.MatchFormat
	f.fieldMap["live_urls"] = f.LiveUrls
	f.fieldMap["serving_side"] = f.ServingSide
	f.fieldMap["match_phase"] = f.MatchPhase
	f.fieldMap["selling_status"] = f.SellingStatus
	f.fieldMap["sub_match_type"] = f.SubMatchType
	f.fieldMap["match_type"] = f.MatchType
	f.fieldMap["season"] = f.Season
	f.fieldMap["scoreboard"] = f.Scoreboard
	f.fieldMap["is_in_play"] = f.IsInPlay
	f.fieldMap["is_recommend"] = f.IsRecommend
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
}

func (f fbSportsMatch) clone(db *gorm.DB) fbSportsMatch {
	f.fbSportsMatchDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbSportsMatch) replaceDB(db *gorm.DB) fbSportsMatch {
	f.fbSportsMatchDo.ReplaceDB(db)
	return f
}

type fbSportsMatchDo struct{ gen.DO }

type IFbSportsMatchDo interface {
	gen.SubQuery
	Debug() IFbSportsMatchDo
	WithContext(ctx context.Context) IFbSportsMatchDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbSportsMatchDo
	WriteDB() IFbSportsMatchDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbSportsMatchDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbSportsMatchDo
	Not(conds ...gen.Condition) IFbSportsMatchDo
	Or(conds ...gen.Condition) IFbSportsMatchDo
	Select(conds ...field.Expr) IFbSportsMatchDo
	Where(conds ...gen.Condition) IFbSportsMatchDo
	Order(conds ...field.Expr) IFbSportsMatchDo
	Distinct(cols ...field.Expr) IFbSportsMatchDo
	Omit(cols ...field.Expr) IFbSportsMatchDo
	Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchDo
	Group(cols ...field.Expr) IFbSportsMatchDo
	Having(conds ...gen.Condition) IFbSportsMatchDo
	Limit(limit int) IFbSportsMatchDo
	Offset(offset int) IFbSportsMatchDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchDo
	Unscoped() IFbSportsMatchDo
	Create(values ...*model.FbSportsMatch) error
	CreateInBatches(values []*model.FbSportsMatch, batchSize int) error
	Save(values ...*model.FbSportsMatch) error
	First() (*model.FbSportsMatch, error)
	Take() (*model.FbSportsMatch, error)
	Last() (*model.FbSportsMatch, error)
	Find() ([]*model.FbSportsMatch, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatch, err error)
	FindInBatches(result *[]*model.FbSportsMatch, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbSportsMatch) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbSportsMatchDo
	Assign(attrs ...field.AssignExpr) IFbSportsMatchDo
	Joins(fields ...field.RelationField) IFbSportsMatchDo
	Preload(fields ...field.RelationField) IFbSportsMatchDo
	FirstOrInit() (*model.FbSportsMatch, error)
	FirstOrCreate() (*model.FbSportsMatch, error)
	FindByPage(offset int, limit int) (result []*model.FbSportsMatch, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbSportsMatchDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbSportsMatchDo) Debug() IFbSportsMatchDo {
	return f.withDO(f.DO.Debug())
}

func (f fbSportsMatchDo) WithContext(ctx context.Context) IFbSportsMatchDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbSportsMatchDo) ReadDB() IFbSportsMatchDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbSportsMatchDo) WriteDB() IFbSportsMatchDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbSportsMatchDo) Session(config *gorm.Session) IFbSportsMatchDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbSportsMatchDo) Clauses(conds ...clause.Expression) IFbSportsMatchDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbSportsMatchDo) Returning(value interface{}, columns ...string) IFbSportsMatchDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbSportsMatchDo) Not(conds ...gen.Condition) IFbSportsMatchDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbSportsMatchDo) Or(conds ...gen.Condition) IFbSportsMatchDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbSportsMatchDo) Select(conds ...field.Expr) IFbSportsMatchDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbSportsMatchDo) Where(conds ...gen.Condition) IFbSportsMatchDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbSportsMatchDo) Order(conds ...field.Expr) IFbSportsMatchDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbSportsMatchDo) Distinct(cols ...field.Expr) IFbSportsMatchDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbSportsMatchDo) Omit(cols ...field.Expr) IFbSportsMatchDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbSportsMatchDo) Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbSportsMatchDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbSportsMatchDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbSportsMatchDo) Group(cols ...field.Expr) IFbSportsMatchDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbSportsMatchDo) Having(conds ...gen.Condition) IFbSportsMatchDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbSportsMatchDo) Limit(limit int) IFbSportsMatchDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbSportsMatchDo) Offset(offset int) IFbSportsMatchDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbSportsMatchDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbSportsMatchDo) Unscoped() IFbSportsMatchDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbSportsMatchDo) Create(values ...*model.FbSportsMatch) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbSportsMatchDo) CreateInBatches(values []*model.FbSportsMatch, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbSportsMatchDo) Save(values ...*model.FbSportsMatch) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbSportsMatchDo) First() (*model.FbSportsMatch, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatch), nil
	}
}

func (f fbSportsMatchDo) Take() (*model.FbSportsMatch, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatch), nil
	}
}

func (f fbSportsMatchDo) Last() (*model.FbSportsMatch, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatch), nil
	}
}

func (f fbSportsMatchDo) Find() ([]*model.FbSportsMatch, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbSportsMatch), err
}

func (f fbSportsMatchDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatch, err error) {
	buf := make([]*model.FbSportsMatch, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbSportsMatchDo) FindInBatches(result *[]*model.FbSportsMatch, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbSportsMatchDo) Attrs(attrs ...field.AssignExpr) IFbSportsMatchDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbSportsMatchDo) Assign(attrs ...field.AssignExpr) IFbSportsMatchDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbSportsMatchDo) Joins(fields ...field.RelationField) IFbSportsMatchDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbSportsMatchDo) Preload(fields ...field.RelationField) IFbSportsMatchDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbSportsMatchDo) FirstOrInit() (*model.FbSportsMatch, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatch), nil
	}
}

func (f fbSportsMatchDo) FirstOrCreate() (*model.FbSportsMatch, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatch), nil
	}
}

func (f fbSportsMatchDo) FindByPage(offset int, limit int) (result []*model.FbSportsMatch, count int64, err error) {
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

func (f fbSportsMatchDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbSportsMatchDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbSportsMatchDo) Delete(models ...*model.FbSportsMatch) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbSportsMatchDo) withDO(do gen.Dao) *fbSportsMatchDo {
	f.DO = *do.(*gen.DO)
	return f
}
