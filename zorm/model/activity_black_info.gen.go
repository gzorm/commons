// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameActivityBlackInfo = "activity_black_info"

// ActivityBlackInfo 活动黑名单信息表
type ActivityBlackInfo struct {
	ID          int64  `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;comment:规则ID，自增主键" json:"id,string"`              // 规则ID，自增主键
	BlackType   int64  `gorm:"column:black_type;type:tinyint;not null;comment:黑名单用户类型 (0: 用户名, 1: IP, 2: 真实姓名, 3: 银行卡)" json:"blackType"` // 黑名单用户类型 (0: 用户名, 1: IP, 2: 真实姓名, 3: 银行卡)
	BlackValue  string `gorm:"column:black_value;type:varchar(255);not null;comment:黑名单用户的值（例如，用户名、IP地址、真实姓名、银行卡号）" json:"blackValue"`    // 黑名单用户的值（例如，用户名、IP地址、真实姓名、银行卡号）
	CreateAt    int64  `gorm:"column:create_at;type:int;not null;comment:创建时间" json:"createAt"`                                           // 创建时间
	UpdateAt    int64  `gorm:"column:update_at;type:int;not null;comment:修改时间" json:"updateAt"`                                           // 修改时间
	OpUser      string `gorm:"column:op_user;type:varchar(32);not null;default:system;comment:操作人" json:"opUser"`                         // 操作人
	BlackStatus int64  `gorm:"column:black_status;type:tinyint;not null;comment:状态：1-启用,2-禁用" json:"blackStatus"`                         // 状态：1-启用,2-禁用
}

// TableName ActivityBlackInfo's table name
func (*ActivityBlackInfo) TableName() string {
	return TableNameActivityBlackInfo
}