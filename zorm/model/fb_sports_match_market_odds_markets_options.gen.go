// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameFbSportsMatchMarketOddsMarketsOptions = "fb_sports_match_market_odds_markets_options"

// FbSportsMatchMarketOddsMarketsOptions 玩法选项表
type FbSportsMatchMarketOddsMarketsOptions struct {
	ID        int64           `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键，自增 ID" json:"id,string"` // 主键，自增 ID
	MarketID  int64           `gorm:"column:market_id;type:bigint;not null;comment:外键，对应玩法表的 ID" json:"marketId"`            // 外键，对应玩法表的 ID
	MatchID   int64           `gorm:"column:match_id;type:bigint;comment:赛事ID" json:"matchId"`                               // 赛事ID
	NameFull  string          `gorm:"column:name_full;type:varchar(100);not null;comment:选项全称" json:"nameFull"`              // 选项全称
	NameShort string          `gorm:"column:name_short;type:varchar(50);not null;comment:选项简称" json:"nameShort"`             // 选项简称
	TeamID    int64           `gorm:"column:team_id;type:bigint;comment:球员或球队 ID" json:"teamId"`                             // 球员或球队 ID
	OpType    int64           `gorm:"column:op_type;type:int;not null;comment:选项类型，例如主、客、大、小" json:"opType"`                 // 选项类型，例如主、客、大、小
	OddsEuro  decimal.Decimal `gorm:"column:odds_euro;type:decimal(8,2);not null;comment:欧盘赔率" json:"oddsEuro"`              // 欧盘赔率
	OddsType  int64           `gorm:"column:odds_type;type:int;not null;comment:赔率类型" json:"oddsType"`                       // 赔率类型
	Outcome   int64           `gorm:"column:outcome;type:int;comment:选项结算结果" json:"outcome"`                                 // 选项结算结果
	Line      string          `gorm:"column:line;type:varchar(50);comment:带线玩法的线" json:"line"`                               // 带线玩法的线
	CreatedAt int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`                                       // 创建时间
	UpdatedAt int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`                                       // 更新时间
}

// TableName FbSportsMatchMarketOddsMarketsOptions's table name
func (*FbSportsMatchMarketOddsMarketsOptions) TableName() string {
	return TableNameFbSportsMatchMarketOddsMarketsOptions
}
