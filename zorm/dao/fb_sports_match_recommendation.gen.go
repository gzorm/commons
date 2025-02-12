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

func newFbSportsMatchRecommendation(db *gorm.DB, opts ...gen.DOOption) fbSportsMatchRecommendation {
	_fbSportsMatchRecommendation := fbSportsMatchRecommendation{}

	_fbSportsMatchRecommendation.fbSportsMatchRecommendationDo.UseDB(db, opts...)
	_fbSportsMatchRecommendation.fbSportsMatchRecommendationDo.UseModel(&model.FbSportsMatchRecommendation{})

	tableName := _fbSportsMatchRecommendation.fbSportsMatchRecommendationDo.TableName()
	_fbSportsMatchRecommendation.ALL = field.NewAsterisk(tableName)
	_fbSportsMatchRecommendation.ID = field.NewInt64(tableName, "id")
	_fbSportsMatchRecommendation.MatchID = field.NewInt64(tableName, "match_id")
	_fbSportsMatchRecommendation.MatchName = field.NewString(tableName, "match_name")
	_fbSportsMatchRecommendation.LeagueID = field.NewInt64(tableName, "league_id")
	_fbSportsMatchRecommendation.LeagueName = field.NewString(tableName, "league_name")
	_fbSportsMatchRecommendation.SportID = field.NewInt64(tableName, "sport_id")
	_fbSportsMatchRecommendation.MatchStatus = field.NewInt64(tableName, "match_status")
	_fbSportsMatchRecommendation.StartTime = field.NewInt64(tableName, "start_time")
	_fbSportsMatchRecommendation.MarketCount = field.NewInt64(tableName, "market_count")
	_fbSportsMatchRecommendation.Neutral = field.NewInt64(tableName, "neutral")
	_fbSportsMatchRecommendation.ScoreList = field.NewString(tableName, "score_list")
	_fbSportsMatchRecommendation.OddsList = field.NewString(tableName, "odds_list")
	_fbSportsMatchRecommendation.LeagueInfo = field.NewString(tableName, "league_info")
	_fbSportsMatchRecommendation.TeamInfo = field.NewString(tableName, "team_info")
	_fbSportsMatchRecommendation.MatchClock = field.NewString(tableName, "match_clock")
	_fbSportsMatchRecommendation.MatchFormatID = field.NewInt64(tableName, "match_format_id")
	_fbSportsMatchRecommendation.MatchFormat = field.NewInt64(tableName, "match_format")
	_fbSportsMatchRecommendation.LiveUrls = field.NewString(tableName, "live_urls")
	_fbSportsMatchRecommendation.ServingSide = field.NewInt64(tableName, "serving_side")
	_fbSportsMatchRecommendation.MatchPhase = field.NewString(tableName, "match_phase")
	_fbSportsMatchRecommendation.SubMatchType = field.NewInt64(tableName, "sub_match_type")
	_fbSportsMatchRecommendation.MatchType = field.NewInt64(tableName, "match_type")
	_fbSportsMatchRecommendation.Season = field.NewString(tableName, "season")
	_fbSportsMatchRecommendation.Scoreboard = field.NewString(tableName, "scoreboard")
	_fbSportsMatchRecommendation.MarketTags = field.NewString(tableName, "market_tags")
	_fbSportsMatchRecommendation.SellingStatus = field.NewInt64(tableName, "selling_status")
	_fbSportsMatchRecommendation.OddsType = field.NewInt64(tableName, "odds_type")
	_fbSportsMatchRecommendation.MarketType = field.NewInt64(tableName, "market_type")
	_fbSportsMatchRecommendation.PlayStage = field.NewInt64(tableName, "play_stage")
	_fbSportsMatchRecommendation.BetOption = field.NewString(tableName, "bet_option")
	_fbSportsMatchRecommendation.BetShortname = field.NewString(tableName, "bet_shortname")
	_fbSportsMatchRecommendation.SelectionType = field.NewInt64(tableName, "selection_type")
	_fbSportsMatchRecommendation.EuropeanOdds = field.NewField(tableName, "european_odds")
	_fbSportsMatchRecommendation.BackupOdds = field.NewField(tableName, "backup_odds")
	_fbSportsMatchRecommendation.SettleResult = field.NewInt64(tableName, "settle_result")
	_fbSportsMatchRecommendation.MarketStatus = field.NewInt64(tableName, "market_status")
	_fbSportsMatchRecommendation.IsParlay = field.NewInt64(tableName, "is_parlay")
	_fbSportsMatchRecommendation.IsBestLine = field.NewInt64(tableName, "is_best_line")
	_fbSportsMatchRecommendation.LineValue = field.NewString(tableName, "line_value")
	_fbSportsMatchRecommendation.LeagueLevel = field.NewInt64(tableName, "league_level")
	_fbSportsMatchRecommendation.RegionID = field.NewInt64(tableName, "region_id")
	_fbSportsMatchRecommendation.RegionName = field.NewString(tableName, "region_name")
	_fbSportsMatchRecommendation.LeagueLogo = field.NewString(tableName, "league_logo")
	_fbSportsMatchRecommendation.CreateTime = field.NewInt64(tableName, "create_time")
	_fbSportsMatchRecommendation.UpdateTime = field.NewInt64(tableName, "update_time")

	_fbSportsMatchRecommendation.fillFieldMap()

	return _fbSportsMatchRecommendation
}

// fbSportsMatchRecommendation 赛事推荐表
type fbSportsMatchRecommendation struct {
	fbSportsMatchRecommendationDo

	ALL           field.Asterisk
	ID            field.Int64  // 主键ID
	MatchID       field.Int64  // 赛事ID
	MatchName     field.String // 赛事名称
	LeagueID      field.Int64  // 联赛ID
	LeagueName    field.String // 联赛名称
	SportID       field.Int64  // 运动种类ID
	MatchStatus   field.Int64  // 赛事进行状态
	StartTime     field.Int64  // 开赛时间
	MarketCount   field.Int64  // 单个赛事玩法总数
	Neutral       field.Int64  // 是否为中立场（0 否，1 是）
	ScoreList     field.String // 比分列表（原 nsg）
	OddsList      field.String // 赔率列表（原 mg）
	LeagueInfo    field.String // 联赛信息（原 lg）
	TeamInfo      field.String // 球队信息，第一个是主队，第二个是客队（原 ts）
	MatchClock    field.String // 比赛时钟信息，滚球走表信息（原 mc）
	MatchFormatID field.Int64  // 赛制的场次、局数、节数（原 fid）
	MatchFormat   field.Int64  // 赛制（原 fmt）
	LiveUrls      field.String // 动画直播地址集合
	ServingSide   field.Int64  // 主/客发球，1主队发球，2客队发球
	MatchPhase    field.String // 赛事阶段（如附加赛、季前赛等）
	SubMatchType  field.Int64  // 滚球赛事当前阶段标识（常规时间、加时赛、点球大战等）
	MatchType     field.Int64  // 赛事类型（1:冠军投注赛事, 2:正常赛事）
	Season        field.String // 冠军赛事赛季（如2019年）
	Scoreboard    field.String // 比分板（原 sb）
	MarketTags    field.String // 盘口组标签集合（如热门、角球等）
	SellingStatus field.Int64  // 销售状态（1 开售，2 暂停）
	OddsType      field.Int64  // 赔率类型（支持搜索）
	MarketType    field.Int64  // 玩法类型
	PlayStage     field.Int64  // 玩法阶段
	BetOption     field.String // 投注选项名称
	BetShortname  field.String // 投注选项简称
	SelectionType field.Int64  // 选项类型（主、客、大、小等）
	EuropeanOdds  field.Field  // 欧洲盘赔率
	BackupOdds    field.Field  // 备用赔率
	SettleResult  field.Int64  // 选项结算结果
	MarketStatus  field.Int64  // 玩法销售状态（0 暂停，1 开售，-1 未开售）
	IsParlay      field.Int64  // 是否支持串关（0 否，1 是）
	IsBestLine    field.Int64  // 是否为最优线（0 否，1 是）
	LineValue     field.String // 带线玩法的线值（如大小球2.5线）
	LeagueLevel   field.Int64  // 联赛等级（值越小，级别越高）
	RegionID      field.Int64  // 区域ID
	RegionName    field.String // 区域名称
	LeagueLogo    field.String // 联赛LOGO地址
	CreateTime    field.Int64  // 记录创建时间
	UpdateTime    field.Int64  // 记录更新时间

	fieldMap map[string]field.Expr
}

func (f fbSportsMatchRecommendation) Table(newTableName string) *fbSportsMatchRecommendation {
	f.fbSportsMatchRecommendationDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbSportsMatchRecommendation) As(alias string) *fbSportsMatchRecommendation {
	f.fbSportsMatchRecommendationDo.DO = *(f.fbSportsMatchRecommendationDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbSportsMatchRecommendation) updateTableName(table string) *fbSportsMatchRecommendation {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.MatchID = field.NewInt64(table, "match_id")
	f.MatchName = field.NewString(table, "match_name")
	f.LeagueID = field.NewInt64(table, "league_id")
	f.LeagueName = field.NewString(table, "league_name")
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
	f.MatchFormatID = field.NewInt64(table, "match_format_id")
	f.MatchFormat = field.NewInt64(table, "match_format")
	f.LiveUrls = field.NewString(table, "live_urls")
	f.ServingSide = field.NewInt64(table, "serving_side")
	f.MatchPhase = field.NewString(table, "match_phase")
	f.SubMatchType = field.NewInt64(table, "sub_match_type")
	f.MatchType = field.NewInt64(table, "match_type")
	f.Season = field.NewString(table, "season")
	f.Scoreboard = field.NewString(table, "scoreboard")
	f.MarketTags = field.NewString(table, "market_tags")
	f.SellingStatus = field.NewInt64(table, "selling_status")
	f.OddsType = field.NewInt64(table, "odds_type")
	f.MarketType = field.NewInt64(table, "market_type")
	f.PlayStage = field.NewInt64(table, "play_stage")
	f.BetOption = field.NewString(table, "bet_option")
	f.BetShortname = field.NewString(table, "bet_shortname")
	f.SelectionType = field.NewInt64(table, "selection_type")
	f.EuropeanOdds = field.NewField(table, "european_odds")
	f.BackupOdds = field.NewField(table, "backup_odds")
	f.SettleResult = field.NewInt64(table, "settle_result")
	f.MarketStatus = field.NewInt64(table, "market_status")
	f.IsParlay = field.NewInt64(table, "is_parlay")
	f.IsBestLine = field.NewInt64(table, "is_best_line")
	f.LineValue = field.NewString(table, "line_value")
	f.LeagueLevel = field.NewInt64(table, "league_level")
	f.RegionID = field.NewInt64(table, "region_id")
	f.RegionName = field.NewString(table, "region_name")
	f.LeagueLogo = field.NewString(table, "league_logo")
	f.CreateTime = field.NewInt64(table, "create_time")
	f.UpdateTime = field.NewInt64(table, "update_time")

	f.fillFieldMap()

	return f
}

func (f *fbSportsMatchRecommendation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbSportsMatchRecommendation) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 45)
	f.fieldMap["id"] = f.ID
	f.fieldMap["match_id"] = f.MatchID
	f.fieldMap["match_name"] = f.MatchName
	f.fieldMap["league_id"] = f.LeagueID
	f.fieldMap["league_name"] = f.LeagueName
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
	f.fieldMap["match_format_id"] = f.MatchFormatID
	f.fieldMap["match_format"] = f.MatchFormat
	f.fieldMap["live_urls"] = f.LiveUrls
	f.fieldMap["serving_side"] = f.ServingSide
	f.fieldMap["match_phase"] = f.MatchPhase
	f.fieldMap["sub_match_type"] = f.SubMatchType
	f.fieldMap["match_type"] = f.MatchType
	f.fieldMap["season"] = f.Season
	f.fieldMap["scoreboard"] = f.Scoreboard
	f.fieldMap["market_tags"] = f.MarketTags
	f.fieldMap["selling_status"] = f.SellingStatus
	f.fieldMap["odds_type"] = f.OddsType
	f.fieldMap["market_type"] = f.MarketType
	f.fieldMap["play_stage"] = f.PlayStage
	f.fieldMap["bet_option"] = f.BetOption
	f.fieldMap["bet_shortname"] = f.BetShortname
	f.fieldMap["selection_type"] = f.SelectionType
	f.fieldMap["european_odds"] = f.EuropeanOdds
	f.fieldMap["backup_odds"] = f.BackupOdds
	f.fieldMap["settle_result"] = f.SettleResult
	f.fieldMap["market_status"] = f.MarketStatus
	f.fieldMap["is_parlay"] = f.IsParlay
	f.fieldMap["is_best_line"] = f.IsBestLine
	f.fieldMap["line_value"] = f.LineValue
	f.fieldMap["league_level"] = f.LeagueLevel
	f.fieldMap["region_id"] = f.RegionID
	f.fieldMap["region_name"] = f.RegionName
	f.fieldMap["league_logo"] = f.LeagueLogo
	f.fieldMap["create_time"] = f.CreateTime
	f.fieldMap["update_time"] = f.UpdateTime
}

func (f fbSportsMatchRecommendation) clone(db *gorm.DB) fbSportsMatchRecommendation {
	f.fbSportsMatchRecommendationDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbSportsMatchRecommendation) replaceDB(db *gorm.DB) fbSportsMatchRecommendation {
	f.fbSportsMatchRecommendationDo.ReplaceDB(db)
	return f
}

type fbSportsMatchRecommendationDo struct{ gen.DO }

type IFbSportsMatchRecommendationDo interface {
	gen.SubQuery
	Debug() IFbSportsMatchRecommendationDo
	WithContext(ctx context.Context) IFbSportsMatchRecommendationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbSportsMatchRecommendationDo
	WriteDB() IFbSportsMatchRecommendationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbSportsMatchRecommendationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbSportsMatchRecommendationDo
	Not(conds ...gen.Condition) IFbSportsMatchRecommendationDo
	Or(conds ...gen.Condition) IFbSportsMatchRecommendationDo
	Select(conds ...field.Expr) IFbSportsMatchRecommendationDo
	Where(conds ...gen.Condition) IFbSportsMatchRecommendationDo
	Order(conds ...field.Expr) IFbSportsMatchRecommendationDo
	Distinct(cols ...field.Expr) IFbSportsMatchRecommendationDo
	Omit(cols ...field.Expr) IFbSportsMatchRecommendationDo
	Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchRecommendationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchRecommendationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchRecommendationDo
	Group(cols ...field.Expr) IFbSportsMatchRecommendationDo
	Having(conds ...gen.Condition) IFbSportsMatchRecommendationDo
	Limit(limit int) IFbSportsMatchRecommendationDo
	Offset(offset int) IFbSportsMatchRecommendationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchRecommendationDo
	Unscoped() IFbSportsMatchRecommendationDo
	Create(values ...*model.FbSportsMatchRecommendation) error
	CreateInBatches(values []*model.FbSportsMatchRecommendation, batchSize int) error
	Save(values ...*model.FbSportsMatchRecommendation) error
	First() (*model.FbSportsMatchRecommendation, error)
	Take() (*model.FbSportsMatchRecommendation, error)
	Last() (*model.FbSportsMatchRecommendation, error)
	Find() ([]*model.FbSportsMatchRecommendation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatchRecommendation, err error)
	FindInBatches(result *[]*model.FbSportsMatchRecommendation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbSportsMatchRecommendation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbSportsMatchRecommendationDo
	Assign(attrs ...field.AssignExpr) IFbSportsMatchRecommendationDo
	Joins(fields ...field.RelationField) IFbSportsMatchRecommendationDo
	Preload(fields ...field.RelationField) IFbSportsMatchRecommendationDo
	FirstOrInit() (*model.FbSportsMatchRecommendation, error)
	FirstOrCreate() (*model.FbSportsMatchRecommendation, error)
	FindByPage(offset int, limit int) (result []*model.FbSportsMatchRecommendation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbSportsMatchRecommendationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbSportsMatchRecommendationDo) Debug() IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Debug())
}

func (f fbSportsMatchRecommendationDo) WithContext(ctx context.Context) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbSportsMatchRecommendationDo) ReadDB() IFbSportsMatchRecommendationDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbSportsMatchRecommendationDo) WriteDB() IFbSportsMatchRecommendationDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbSportsMatchRecommendationDo) Session(config *gorm.Session) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbSportsMatchRecommendationDo) Clauses(conds ...clause.Expression) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbSportsMatchRecommendationDo) Returning(value interface{}, columns ...string) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbSportsMatchRecommendationDo) Not(conds ...gen.Condition) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbSportsMatchRecommendationDo) Or(conds ...gen.Condition) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbSportsMatchRecommendationDo) Select(conds ...field.Expr) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbSportsMatchRecommendationDo) Where(conds ...gen.Condition) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbSportsMatchRecommendationDo) Order(conds ...field.Expr) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbSportsMatchRecommendationDo) Distinct(cols ...field.Expr) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbSportsMatchRecommendationDo) Omit(cols ...field.Expr) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbSportsMatchRecommendationDo) Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbSportsMatchRecommendationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbSportsMatchRecommendationDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbSportsMatchRecommendationDo) Group(cols ...field.Expr) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbSportsMatchRecommendationDo) Having(conds ...gen.Condition) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbSportsMatchRecommendationDo) Limit(limit int) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbSportsMatchRecommendationDo) Offset(offset int) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbSportsMatchRecommendationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbSportsMatchRecommendationDo) Unscoped() IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbSportsMatchRecommendationDo) Create(values ...*model.FbSportsMatchRecommendation) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbSportsMatchRecommendationDo) CreateInBatches(values []*model.FbSportsMatchRecommendation, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbSportsMatchRecommendationDo) Save(values ...*model.FbSportsMatchRecommendation) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbSportsMatchRecommendationDo) First() (*model.FbSportsMatchRecommendation, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchRecommendation), nil
	}
}

func (f fbSportsMatchRecommendationDo) Take() (*model.FbSportsMatchRecommendation, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchRecommendation), nil
	}
}

func (f fbSportsMatchRecommendationDo) Last() (*model.FbSportsMatchRecommendation, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchRecommendation), nil
	}
}

func (f fbSportsMatchRecommendationDo) Find() ([]*model.FbSportsMatchRecommendation, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbSportsMatchRecommendation), err
}

func (f fbSportsMatchRecommendationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatchRecommendation, err error) {
	buf := make([]*model.FbSportsMatchRecommendation, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbSportsMatchRecommendationDo) FindInBatches(result *[]*model.FbSportsMatchRecommendation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbSportsMatchRecommendationDo) Attrs(attrs ...field.AssignExpr) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbSportsMatchRecommendationDo) Assign(attrs ...field.AssignExpr) IFbSportsMatchRecommendationDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbSportsMatchRecommendationDo) Joins(fields ...field.RelationField) IFbSportsMatchRecommendationDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbSportsMatchRecommendationDo) Preload(fields ...field.RelationField) IFbSportsMatchRecommendationDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbSportsMatchRecommendationDo) FirstOrInit() (*model.FbSportsMatchRecommendation, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchRecommendation), nil
	}
}

func (f fbSportsMatchRecommendationDo) FirstOrCreate() (*model.FbSportsMatchRecommendation, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchRecommendation), nil
	}
}

func (f fbSportsMatchRecommendationDo) FindByPage(offset int, limit int) (result []*model.FbSportsMatchRecommendation, count int64, err error) {
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

func (f fbSportsMatchRecommendationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbSportsMatchRecommendationDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbSportsMatchRecommendationDo) Delete(models ...*model.FbSportsMatchRecommendation) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbSportsMatchRecommendationDo) withDO(do gen.Dao) *fbSportsMatchRecommendationDo {
	f.DO = *do.(*gen.DO)
	return f
}
