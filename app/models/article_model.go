package models

import "time"

type Article struct {
	ID          uint      `db:"id" json:"id" gorm:"primary_key"`
	Title       string    `gorm:"not null" json:"price"`
    Category    []Category  `gorm:"Foreignkey:ArticleID;association_foreignkey:ID;" json:"category"`
    CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}