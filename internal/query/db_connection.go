package query

import "gorm.io/gorm"

var dbConnection Gorm

type Gorm interface {
	Db() *gorm.DB
}

func SetDbProvider(connection Gorm) {
	dbConnection = connection
}

func Db() *gorm.DB {
	if dbConnection == nil {
		return nil
	}

	return dbConnection.Db()
}
