package model

import (
	"time"

	"gorm.io/datatypes"
)

// Job job model
type Job struct {
	ID          uint64         `gorm:"column:id;primaryKey;autoIncrement;not null"`
	Queue       string         `gorm:"column:queue;primaryKey"`
	Payload     datatypes.JSON `gorm:"column:payload"`
	Attempts    uint8          `gorm:"column:attempts"`
	ExecutedAt  *time.Time     `gorm:"column:executed_at"`
	AvaliableAt *time.Time     `gorm:"column:avaliable_at"`
	CreatedAt   *time.Time     `gorm:"column:created_at;autoCreateTime"`
}
