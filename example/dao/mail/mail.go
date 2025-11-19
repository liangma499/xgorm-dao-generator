package mail

import (
	"xgorm-dao-generator/example/dao/mail/internal"
	"gorm.io/gorm"
)

type (
	Columns = internal.Columns
	OrderBy = internal.OrderBy
	FilterFunc = internal.FilterFunc
	UpdateFunc = internal.UpdateFunc
	ColumnFunc = internal.ColumnFunc
	OrderFunc = internal.OrderFunc
)

type Mail struct {
	*internal.Mail
}

func NewMail(db *gorm.DB) *Mail {
	return &Mail{Mail: internal.NewMail(db)}
}
