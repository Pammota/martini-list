package models

type Item struct {
	ID        int    `json:"id" sql:"AUTO_INCREMENT" gorm:"primaryKey;index;unique;autoIncrement"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
