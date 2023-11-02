package fixtures

import (
	"fmt"

	"github.com/bifidokk/go-api/internal/entity"
	"gorm.io/gorm"
)

type Tables map[string]interface{}

var Entities = Tables{
	entity.Note{}.TableName(): &entity.Note{},
	entity.User{}.TableName(): &entity.User{},
}

func (list Tables) ResetTestFixtures(db *gorm.DB) {
	list.Truncate(db)
	CreateTestFixtures(db)
}

func (list Tables) Truncate(db *gorm.DB) {
	for tableName := range list {
		if err := db.Exec(fmt.Sprintf("DELETE FROM %s", tableName)).Error; err == nil {
			fmt.Println("Remove data from", tableName)
		} else if err.Error() != "record not found" {
			fmt.Printf("Migrate: %s in %s\n", err, tableName)
		}
	}
}

func CreateTestFixtures(db *gorm.DB) {
	CreateUserFixtures(db)
}
