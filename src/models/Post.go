package models

import (
	"time"
)

type Post struct {
	Id             uint64    `gorm:"autoIncrement;primaryKey" json:"id,omitempty"`
	Title          string    `gorm:"type:varchar(50);not null" json:"title,omitempty"`
	Content        string    `gorm:"type:varchar(500);not null" json:"content,omitempty"`
	AuthorId       uint64    `gorm:"primaryKey" json:"author_id,omitempty"`
	User           User      `gorm:"foreignKey:AuthorId;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	AuthorUsername string    `gorm:"-" json:"author_username,omitempty"`
	Likes          uint64    `gorm:"int;default:0" json:"likes"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
}
