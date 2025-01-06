// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinDictItem = "win_dict_item"

// WinDictItem 字典项表
type WinDictItem struct {
	ID        int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	Code      string `gorm:"column:code;type:varchar(32);not null;comment:字典码" json:"code"`                          // 字典码
	Title     string `gorm:"column:title;type:varchar(128);not null;comment:字典名称" json:"title"`                      // 字典名称
	TitleEn   string `gorm:"column:title_en;type:varchar(225);comment:英文" json:"titleEn"`                            // 英文
	TitleAr   string `gorm:"column:title_ar;type:varchar(225);comment:阿拉伯文" json:"titleAr"`                          // 阿拉伯文
	Remark    string `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark"`                               // 备注
	Sort      int64  `gorm:"column:sort;type:int;not null;comment:排序:从高到低" json:"sort"`                              // 排序:从高到低
	ReferID   int64  `gorm:"column:refer_id;type:int;not null;comment:字典表ID" json:"referId"`                         // 字典表ID
	Status    int64  `gorm:"column:status;type:tinyint;not null;default:1;comment:状态:1-启用 0-禁用" json:"status"`       // 状态:1-启用 0-禁用
	IsShow    int64  `gorm:"column:is_show;type:tinyint;not null;default:2;comment:类型:0-全部 1-前端 2-后台" json:"isShow"` // 类型:0-全部 1-前端 2-后台
	CreatedAt int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinDictItem's table name
func (*WinDictItem) TableName() string {
	return TableNameWinDictItem
}