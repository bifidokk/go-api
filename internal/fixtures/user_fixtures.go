package fixtures

import (
	"fmt"
	"log"

	"github.com/bifidokk/go-api/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserMap []entity.User

var UserFixtures = UserMap{
	{
		Email:    "user@test.com",
		Password: "123456!",
	},
	{
		Email:    "user2@test.com",
		Password: "123456!",
	},
}

func CreateUserFixtures(db *gorm.DB) {
	fmt.Println("Create user fixtures")

	for _, entity := range UserFixtures {
		password, err := bcrypt.GenerateFromPassword(
			[]byte(entity.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			log.Println("Error occurred during user fixture password generation ", err)
			break
		}

		entity.Password = string(password)

		db.Create(&entity)
	}
}
