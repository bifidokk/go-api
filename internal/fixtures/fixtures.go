package fixtures

import (
	"fmt"

	"github.com/bifidokk/go-api/internal/entity"
	"gorm.io/gorm"
)

type fixtures struct {
	database *gorm.DB
	tables   Tables
}

type Fixtures interface {
	ResetTestFixtures()
}

type Tables map[string]interface{}

var Entities = Tables{
	entity.Note{}.TableName(): &entity.Note{},
	entity.User{}.TableName(): &entity.User{},
}

func NewFixtures(db *gorm.DB) Fixtures {
	return &fixtures{
		database: db,
		tables:   Entities,
	}
}

func (fixtures *fixtures) ResetTestFixtures() {
	fixtures.Truncate()
	fixtures.CreateTestFixtures()
}

func (fixtures *fixtures) Truncate() {
	for tableName := range fixtures.tables {
		if err := fixtures.database.Exec(fmt.Sprintf("TRUNCATE %s RESTART IDENTITY CASCADE", tableName)).Error; err == nil {
			fmt.Println("Remove data from", tableName)
		} else if err.Error() != "record not found" {
			fmt.Printf("Migrate: %s in %s\n", err, tableName)
		}
	}
}

func (fixtures *fixtures) CreateTestFixtures() {
	CreateUserFixtures(fixtures.database)
	CreateNoteFixtures(fixtures.database)
}
