package query

import "github.com/bifidokk/go-api/internal/entity"

func Notes() (results entity.Notes, err error) {
	err = Db().Find(&results).Error

	return results, err
}
