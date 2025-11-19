package model

import (
	"time"
)

//go:generate xgorm-dao-generator -model-dir=. -model-names=MailE -dao-dir=../dao/ -sub-pkg-enable=true
type MailE struct {
	ID       int       `gorm:"column:id"`        // 邮件ID
	Title    string    `gorm:"column:title"`     // 邮件标题
	Content  string    `gorm:"column:content"`   // 邮件内容
	Sender   int64     `gorm:"column:sender"`    // 邮件发送者
	Receiver int64     `gorm:"column:receiver"`  // 邮件接受者
	Status   int       `gorm:"column:status"`    // 邮件状态
	SendTime time.Time `gorm:"column:send_time"` // 发送时间
}
