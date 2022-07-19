package po

import "time"

type Example struct {
	ID        string     `gorm:"id"`
	Name      string     `gorm:"name"`
	CreatedAt *time.Time `gorm:"<-:create;column:created_at"` // 建立時間
	UpdatedAt *time.Time `gorm:"column:updated_at"`           // 修改時間
}

func (Example) TableName() string {
	return "example"
}

type ExampleCondByGet struct {
	ID string
}
