package models

type Skill struct {
	Id   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"size:100;uniqueIndex;not null" json:"name"`
}
