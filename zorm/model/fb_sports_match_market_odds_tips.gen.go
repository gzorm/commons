// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFbSportsMatchMarketOddsTips = "fb_sports_match_market_odds_tips"

// FbSportsMatchMarketOddsTips 赛事玩法推荐
type FbSportsMatchMarketOddsTips struct {
	ID          int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	MatchID     int64  `gorm:"column:match_id;type:bigint;not null;comment:赛事ID 原 mid" json:"matchId"`       // 赛事ID 原 mid
	Options     string `gorm:"column:options;type:text;comment:选项集合 原 ops" json:"options"`                   // 选项集合 原 ops
	TeamsInfo   string `gorm:"column:teams_info;type:text;comment:球队信息，第一个是主队，第二个是客队 原 ts" json:"teamsInfo"` // 球队信息，第一个是主队，第二个是客队 原 ts
	StartTime   int64  `gorm:"column:start_time;type:bigint;not null;comment:赛事开赛时间 原 bt" json:"startTime"`  // 赛事开赛时间 原 bt
	LeagueInfo  string `gorm:"column:league_info;type:text;comment:联赛信息 原 lg" json:"leagueInfo"`             // 联赛信息 原 lg
	MatchStatus int64  `gorm:"column:match_status;type:int;not null;comment:赛事进行状态 原 ms" json:"matchStatus"` // 赛事进行状态 原 ms
	CreatedAt   int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`                              // 创建时间戳
	UpdatedAt   int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`                              // 更新时间戳
}

// TableName FbSportsMatchMarketOddsTips's table name
func (*FbSportsMatchMarketOddsTips) TableName() string {
	return TableNameFbSportsMatchMarketOddsTips
}
