package models

type Item struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Index     int    `json:"index" sql:"AUTO_INCREMENT" gorm:"index;unique;autoIncrement"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
