package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Tags []string

type SourceSeed struct {
	ID          uint      `json:"id" gorm:"primarykey;->"`
	SourceID    uint      `gorm:"not null"`
	SourceURL   string    `gorm:"not null;unique"`
	Title       string    `gorm:"not null"`
	Cover       string    `gorm:"not null"`
	BigCover    string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	PublishTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	Type        int       `gorm:"not null;comment:'0电影 1国产剧 2美剧 3韩剧 4港剧 5日剧 6动画 7音乐 8精品'"`
	Tag         Tags      `gorm:"type:json"`
	Origin      string    `gorm:"not null"`
	Content     string    `gorm:"not null;type:text"`
	Links       string    `gorm:"not null;type:text"`
	UpdatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func (SourceSeed) TableName() string {
	return "seed"
}

func (t *Tags) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Tags) Value() (driver.Value, error) {
	return json.Marshal(t)
}
