package fixtures

import (
	"fmt"

	"github.com/bifidokk/go-api/internal/entity"
	"gorm.io/gorm"
)

type NoteMap []entity.Note

var NoteFixtures = NoteMap{
	{
		NoteTitle:       "Title 1",
		NoteDescription: "Description 1",
		UserID:          1,
	},
	{
		NoteTitle:       "Title 2",
		NoteDescription: "Description 2",
		UserID:          1,
	},
	{
		NoteTitle:       "Title 3",
		NoteDescription: "Description 3",
		UserID:          2,
	},
}

func CreateNoteFixtures(db *gorm.DB) {
	fmt.Println("Create note fixtures")

	for _, entity := range NoteFixtures {
		db.Create(&entity)
	}
}
