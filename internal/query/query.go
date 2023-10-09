package query

import (
	"github.com/bifidokk/go-api/internal/entity"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	return entity.Db()
}
