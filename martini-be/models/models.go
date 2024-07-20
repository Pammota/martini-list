package models

type Item struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Index     int    `json:"index" gorm:"index,unique"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
