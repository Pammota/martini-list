package models

type Item struct {
	ID        uint   `json:"id" gorm:"primaryKey;uniqueIndex;autoIncrement"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
