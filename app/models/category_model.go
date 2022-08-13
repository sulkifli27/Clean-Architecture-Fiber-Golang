package models

import "time"

type Category struct {
	ID        uint      `db:"id" json:"id" gorm:"primary_key"`
    Name      string    `gorm:"not null" json:"price"`
	ArticleID uint      `json:"article_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}