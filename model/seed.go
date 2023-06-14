package model

import (
	"database/sql/driver"
	"encoding/json"
	"pow/global/orm"
	"time"
)

type Tags []string

type SourceSeed struct {
	ID          uint      `json:"id" gorm:"primarykey;->"`
	SourceURL   string    `gorm:"not null;unique"`
	Title       string    `gorm:"not null"`
	Cover       string    `gorm:"not null"`
	BigCover    string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	PublishTime time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	Type        string    `gorm:"not null;"`
	Tag         Tags      `gorm:"type:json"`
	Origin      string    `gorm:"not null"`
	Content     string    `gorm:"not null;type:text"`
	RawContent  string    `gorm:"not null;type:text"`
	Links       string    `gorm:"not null;type:text"`
	UpdatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func (SourceSeed) TableName() string {
	return "source_seed"
}

func (t *Tags) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Tags) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (s *SourceSeed) Exists(url string) bool {
	result := orm.Eloquent.Where("source_url", url).Limit(1).Find(&s)
	if result.Error == nil && result.RowsAffected == 1 {
		return true
	}
	return false
}
