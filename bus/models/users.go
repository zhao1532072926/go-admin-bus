package models

type Users struct {
	Phone    string `json:"phone" gorm:"type:varchar(15);comment:phone"`
	Password string `json:"password" gorm:"type:varchar(20);comment:password"`
}

func (Users) TableName() string {
	return "" // 表名将在运行时动态设置
}