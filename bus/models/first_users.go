package models

type FirstUsers struct {
	Phone    string `json:"phone" gorm:"type:varchar(15);comment:phone"`
	Password string `json:"password" gorm:"type:varchar(20);comment:password"`
}

func (FirstUsers) TableName() string {
	return "first_users"
}
