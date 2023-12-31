package models

type Photo struct {
	ID       string `gorm:"primaryKey;not null" valid:"required" json:"id"`
	Title    string `gorm:"type:text;not null" json:"title"`
	Caption  string `gorm:"type:text;not null" json:"caption"`
	PhotoURL string `gorm:"type:text;not null" json:"photourl"`
	UserID   string `gorm:"column:userid;references:ID;not null;" json:"userid"`
}
