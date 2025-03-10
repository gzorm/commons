// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFbSportsMatchStatisticsStageList = "fb_sports_match_statistics_stage_list"

// FbSportsMatchStatisticsStageList 赛事类型表，存储每种类型的赛事统计信息
type FbSportsMatchStatisticsStageList struct {
	ID           int64  `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键ID" json:"id,string"`                                  // 主键ID
	StatisticsID int64  `gorm:"column:statistics_id;type:bigint;not null;comment:关联主表的 ID（对应主表 fb_sports_match_statistics.id）" json:"statisticsId"` // 关联主表的 ID（对应主表 fb_sports_match_statistics.id）
	SlType       int64  `gorm:"column:sl_type;type:int;not null;comment:分类类型（对应 data.sl.ty，枚举 match_play_type）" json:"slType"`                      // 分类类型（对应 data.sl.ty，枚举 match_play_type）
	Description  string `gorm:"column:description;type:varchar(255);comment:分类描述（对应 data.sl.des）" json:"description"`                               // 分类描述（对应 data.sl.des）
	TotalCount   int64  `gorm:"column:total_count;type:int;not null;comment:赛事总数（对应 data.sl.tc）" json:"totalCount"`                                 // 赛事总数（对应 data.sl.tc）
	SubStageList string `gorm:"column:sub_stage_list;type:text;comment:每个类型下的子赛事统计信息（对应 data.sl.ssl，存储为 JSON）" json:"subStageList"`                 // 每个类型下的子赛事统计信息（对应 data.sl.ssl，存储为 JSON）
	CreatedAt    int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`                                                                    // 记录创建时间（时间戳）
	UpdatedAt    int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`                                                                    // 记录更新时间（时间戳）
}

// TableName FbSportsMatchStatisticsStageList's table name
func (*FbSportsMatchStatisticsStageList) TableName() string {
	return TableNameFbSportsMatchStatisticsStageList
}
