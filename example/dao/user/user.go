package user

import (
	"sync"

	mysqlimp "github.com/liangma499/xgorm-dao-generator/example/bbb/cc"
	"github.com/liangma499/xgorm-dao-generator/example/dao/user/internal"
	modelpkg "github.com/liangma499/xgorm-dao-generator/example/model"
	"gorm.io/gorm"
)

type (
	Columns    = internal.Columns
	OrderBy    = internal.OrderBy
	FilterFunc = internal.FilterFunc
	UpdateFunc = internal.UpdateFunc
	ColumnFunc = internal.ColumnFunc
	OrderFunc  = internal.OrderFunc
)

type User struct {
	*internal.User
}

func NewUser(db *gorm.DB) *User {
	return &User{User: internal.NewUser(db)}
}

var (
	once     sync.Once
	instance *User
)

func Instance() *User {
	once.Do(func() {
		instance = NewUser(mysqlimp.Instance())
	})
	return instance
}
func (dao *User) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.User{})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
