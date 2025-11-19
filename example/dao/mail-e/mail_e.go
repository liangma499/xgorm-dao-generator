package maile

import (
	"xgorm-dao-generator/example/dao/mail-e/internal"
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

type MailE struct {
	*internal.MailE
}

func NewMailE(db *gorm.DB) *MailE {
	return &MailE{MailE: internal.NewMailE(db)}
}
