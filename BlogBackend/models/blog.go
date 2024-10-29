package models

type Blog struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Image  string `json:"image"`
	UserID string `json:"userid"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
}
