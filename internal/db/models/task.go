package models

import "time"

import "github.com/nrssi/kanbee/internal/services/kanban"

type TaskState string

const (
	TODO        TaskState = "TODO"
	IN_PROGRESS TaskState = "IN_PROGRESS"
	DONE        TaskState = "DONE"
	BLOCKED     TaskState = "BLOCKED"
)

type Task struct {
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	Deadline    time.Time
	Title       string        `gorm:"not null"`
	Description string        `gorm:"type:text"`
	Updates     []Update      `gorm:"foreignKey:TaskID"`
	ID          uint          `gorm:"primaryKey"`
	ProjectID   uint          `gorm:"not null"`
	State       kbp.TaskState `gorm:"type:varchar(20);not null"`
}
