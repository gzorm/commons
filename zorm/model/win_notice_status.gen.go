// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinNoticeStatus = "win_notice_status"

// WinNoticeStatus 用户消息状态表
type WinNoticeStatus struct {
	ID       int64 `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:ID" json:"id,string"`                  // ID
	UID      int64 `gorm:"column:uid;type:int;not null;comment:用户ID" json:"uid"`                                          // 用户ID
	NoticeID int64 `gorm:"column:notice_id;type:bigint;not null;comment:消息ID" json:"noticeId"`                            // 消息ID
	IsRead   int64 `gorm:"column:is_read;type:int;not null;comment:0-未读1-已读" json:"isRead"`                               // 0-未读1-已读
	IsDelete int64 `gorm:"column:is_delete;type:int;not null;comment:0-未删除1-删除" json:"isDelete"`                          // 0-未删除1-删除
	CreateAt int64 `gorm:"column:create_at;type:int;not null;comment:创建时间" json:"createAt"`                               // 创建时间
	UpdateAt int64 `gorm:"column:update_at;type:int;not null;comment:创基时间" json:"updateAt"`                               // 创基时间
	Category int64 `gorm:"column:category;type:tinyint;not null;default:4;comment:类型:1-系统公告2-站内信 3-系统消息" json:"category"` // 类型:1-系统公告2-站内信 3-系统消息
}

// TableName WinNoticeStatus's table name
func (*WinNoticeStatus) TableName() string {
	return TableNameWinNoticeStatus
}